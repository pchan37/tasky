package views

import (
	"net/http"
)

func RegisterStaticViews() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
}
