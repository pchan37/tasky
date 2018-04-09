package views

import (
	"net/http"

	"github.com/PGonLib/PGo-Auth/pkg/security"
	"github.com/pchan37/tasky/app/lib/templateManager"
)

func RegisterPrivateViews() {
	http.HandleFunc("/", security.AuthenticationHandler(IndexPage))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	templateManager.RenderTemplate(w, "index.tmpl", nil)
}
