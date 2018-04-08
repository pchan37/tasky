package views

import (
	"net/http"

	"github.com/pchan37/tasky/app/lib/templateManager"
)

func RegisterPublicViews() {
	http.HandleFunc("/", IndexPage)
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderTemplate(w, "index.tmpl", nil)
}
