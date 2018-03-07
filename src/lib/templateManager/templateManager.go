package templateManager

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"tasky/pkg/bpool"
)

type templateConfig struct {
	layoutPath string
	includePath string
}

var templates map[string]*template.Template
var bufpool *bpool.BufferPool
var config *templateConfig

var mainTmpl = `{{ define "main" }} {{ template "base" . }} {{ end }}`

func SetTemplateConfig(layoutPath string, includePath string) {
	config = &templateConfig{layoutPath, includePath}
}

func LoadTemplates() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	layoutFiles, _ := getLayoutFiles()
	includeFiles, _ := getIncludeFiles()
	mainTemplate, _ := getMainTemplate()

	for _, file := range includeFiles {
		fileName := filepath.Base(file)
		files := append(layoutFiles, file)
		var err error
		templates[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal(err)
		}
		templates[fileName] = template.Must(templates[fileName].ParseFiles(files...))
	}
	log.Println("Successfully loaded templates!")

	bufpool = bpool.NewBufferPool(64)
	log.Println("Successfully allocated buffers!")
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	template, ok := templates[name]
	if !ok {
		http.Error(w, fmt.Sprintf("The template %s does not exist.", name),
			http.StatusInternalServerError)
	}

	buf := bufpool.Get()
	defer bufpool.Put(buf)

	err := template.Execute(buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
