package handlers

import (
	//"log"
	"fmt"
	"strings"

	"github.com/bryanbarcos/skip-the-choices/internal/models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

func SearchHandler(e *core.RequestEvent) error {
	app := e.App
	query := strings.ToLower(e.Request.URL.Query().Get("query"))
	//page := e.Request.URL.Query().Get("page")
	fmt.Println("Query: ", query)

	// offset = (page - 1) * itemsPerPage + 1
	//offset := (page - 1) * 5 + 1

	expr2 := dbx.Like("name", query)
	// find records using pagination, using
	// records, err := app.FindRecordsByFilter(
	// 	"restaurants",     // collection
	// 	"name ~ {:query}", // filter
	// 	"-name",           // sort by name descending
	// 	5,                 // limit
	// 	0,                 // offset
	// )
	records := []models.Restaurant{}
	err := app.DB().
		Select("name", "category").
		From("restaurants").
		AndWhere(expr2).
		Limit(5).
		Offset(0).
		All(&records)

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	var htmlResults []string
	for _, result := range records {
		fmt.Println("Result value: ", result.Name)
		htmlResults = append(htmlResults, fmt.Sprintf(
			`<div class="p-2 border rounded mt-2 fade-in">%s</div>`, result.Name,
		))
	}
	e.Response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(e.Response, strings.Join(htmlResults, ""))
	return nil
}
