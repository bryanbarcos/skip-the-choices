package handlers

import (
	//"log"
	"net/http"

	"github.com/bryanbarcos/skip-the-choices/ui"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ui.MainTempl("Bryan Barcos Hello World").Render(r.Context(), w)
}
