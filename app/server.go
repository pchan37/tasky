package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PGonLib/PGo-Auth/pkg/security"
	"github.com/pchan37/tasky/app/lib/dbManager"
	"github.com/pchan37/tasky/app/lib/taskDatabase"
	"github.com/pchan37/tasky/app/lib/templateManager"
	"github.com/pchan37/tasky/app/views"
)

const SOCK = "/tmp/tasky.sock"

type config struct {
	LayoutPath  string
	IncludePath string
}

func loadConfig(filename string) {
	file, _ := os.Open(filename)
	decoder := json.NewDecoder(file)
	config := config{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("Error unpacking config:", err)
	}
	log.Println("Layout path:", config.LayoutPath)
	log.Println("Include path:", config.IncludePath)
	templateManager.SetTemplateConfig(config.LayoutPath, config.IncludePath)
}

func main() {
	loadConfig("config/config.json")
	templateManager.LoadTemplates()

	manager := taskDatabase.InitializeDatabase()
	defer dbManager.Close(manager)

	authKey, err := ioutil.ReadFile("keys/authKey")
	if err != nil {
		panic(err)
	}
	encryptKey, err := ioutil.ReadFile("keys/encryptKey")
	if err != nil {
		panic(err)
	}

	keys := security.Keys{
		AuthKey:    authKey,
		EncryptKey: encryptKey,
	}
	authManager := security.InitAuthManager("127.0.0.1:27017", "tasky", keys)
	defer authManager.Close()

	os.Remove(SOCK)
	unixListener, err := net.Listen("unix", SOCK)
	if err != nil {
		log.Fatal("Listen (UNIX socket): ", err)
	}
	os.Chmod(SOCK, 0666)
	log.Println("Ready!")
	defer unixListener.Close()

	views.RegisterStaticViews()
	views.RegisterPrivateViews()
	views.RegisterTaskViews()
	views.RegisterSecurityViews()

	http.Serve(unixListener, http.DefaultServeMux)
}
