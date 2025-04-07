package handlers

import (
	//"log"
	"fmt"
	"net/http"

	//"strings"

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

		// Select("name", "category").
		// From("user_restaurants").
		// AndWhere(dbx.NewExp("user_id = {:userId}", dbx.Params{ "userId": userId })).

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	// rList := []models.Restaurant{
	// 	{Name: "Holy Chuck", Category: "burgers"},
	// 	{Name: "Burger Priest", Category: "burgers"},
	// }
	ui.MainTempl(records).Render(e.Request.Context(), e.Response)
	return nil
}
