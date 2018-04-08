package views

import (
	"net/http"
	"tasky/src/lib/templateManager"
)

func RegisterPublicViews() {
	http.HandleFunc("/", IndexPage)
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderTemplate(w, "index.tmpl", nil)
}
