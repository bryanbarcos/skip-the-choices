package handlers

import (
	//"log"
	//"fmt"
	"net/http"
	//"strings"

	"github.com/bryanbarcos/skip-the-choices/ui"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ui.MainTempl().Render(r.Context(), w)
}
