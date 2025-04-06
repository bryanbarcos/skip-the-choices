package handlers

import (
	//"log"
	"fmt"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func SearchHandler(e *core.RequestEvent) error {
	app := e.App
	query := strings.ToLower(e.Request.URL.Query().Get("query"))
	fmt.Println("Query: ", query)

	expr2 := dbx.Like("name", query)
	records, err := app.FindAllRecords("restaurants", expr2)

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	var htmlResults []string
	for _, result := range records {
		fmt.Println("Result value: ", result.Get("name"))
		htmlResults = append(htmlResults, fmt.Sprintf(
			`<div class="p-2 border rounded mt-2 fade-in">%s</div>`, result.Get("name"),
		))
	}
	e.Response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(e.Response, strings.Join(htmlResults, ""))
	return nil
}
