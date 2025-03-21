package handlers

import (
	//"log"
	//"fmt"
	"net/http"
	//"strings"

	"github.com/bryanbarcos/skip-the-choices/ui"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	if e.Request.URL.RequestURI() != "/" {
		http.NotFound(e.Response, e.Request)
		return nil
	}
	ui.MainTempl().Render(e.Request.Context(), e.Response)
	return nil
}
