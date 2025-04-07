package handlers

import (
	//"log"
	"fmt"
	"net/http"

	"github.com/bryanbarcos/skip-the-choices/internal/models"
	ui "github.com/bryanbarcos/skip-the-choices/ui/pages"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func Home(e *core.RequestEvent) error {
	if e.Request.URL.RequestURI() != "/" {
		http.NotFound(e.Response, e.Request)
		return nil
	}
	userId := "hm77ox88p0j5t4v"
	app := e.App
	records := []models.Restaurant{}

	// Get list of restaurants associated with userId from user_restaurants
	// table, then get restaurant details from restaurants table with rest_id
	err := app.DB().
		NewQuery(`
			SELECT *
			FROM (
				SELECT rest_id FROM user_restaurants ur
				WHERE ur.user_id = {:userId}
			) as rests
			LEFT JOIN restaurants ur on ur.id = rests.rest_id
			LIMIT 10
		`).
		Bind(dbx.Params{
			"userId": userId,
		}).
		All(&records)

	if err != nil {
		fmt.Println("error loading restaurant list: ", err)
		return err
	}

	// load the home page templ and pass restaurant list query
	ui.MainTempl(records).Render(e.Request.Context(), e.Response)
	return nil
}
