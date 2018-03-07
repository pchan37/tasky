package templateManager

import (
	"html/template"
	"log"
	"path/filepath"
)

func getLayoutFiles() (layoutFiles []string, err error) {
	layoutFiles, err = filepath.Glob(config.layoutPath + "*.tmpl");
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getIncludeFiles() (includeFiles []string, err error) {
	includeFiles, err = filepath.Glob(config.includePath + "*.tmpl");
	if err != nil {
		log.Fatal(err)
	}
	return
}

func getMainTemplate() (mainTemplate *template.Template, err error) {
	mainTemplate, err = template.New("main").Parse(mainTmpl);
	if err != nil {
		log.Fatal(err)
	}
	return
}
