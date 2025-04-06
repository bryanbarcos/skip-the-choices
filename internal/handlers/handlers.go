package handlers

import (
	//"log"
	//"fmt"
	"net/http"
	//"strings"

	"github.com/bryanbarcos/skip-the-choices/internal/models"
	ui "github.com/bryanbarcos/skip-the-choices/ui/pages"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	if e.Request.URL.RequestURI() != "/" {
		http.NotFound(e.Response, e.Request)
		return nil
	}
	rList := []models.Restaurant{
		{Name: "Holy Chuck", Category: "burgers"},
		{Name: "Burger Priest", Category: "burgers"},
	}
	ui.MainTempl(rList).Render(e.Request.Context(), e.Response)
	return nil
}
