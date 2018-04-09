package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/PGonLib/PGo-Auth/pkg/security"
)

func main() {
	basepath, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("There was an error when generating the basepath for the keys...", err.Error())
		os.Exit(1)
	}
	if err = security.CreateKeyFiles(basepath, 32, 32); err != nil {
		fmt.Println("Error generating keys...", err.Error())
		os.Exit(1)
	}
	fmt.Println("Successfully generated keys!!!")
}
