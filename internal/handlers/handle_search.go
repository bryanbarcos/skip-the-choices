package handlers

import (
	//"log"
	"fmt"
	"net/http"
	"strings"

	"github.com/bryanbarcos/skip-the-choices/internal/models"
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

// Search database restaurants table for restaurants with name
// like input search -- user will be able to add results to their
// restaurant list
func SearchHandler(e *core.RequestEvent) error {
	app := e.App
	query := strings.ToLower(e.Request.URL.Query().Get("query"))
	//page := e.Request.URL.Query().Get("page")
	fmt.Println("Query: ", query)

	// offset = (page - 1) * itemsPerPage + 1
	//offset := (page - 1) * 5 + 1

	expr2 := dbx.Like("name", query)

	records := []models.Restaurant{}
	err := app.DB().
		Select("id", "name", "category").
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
		htmlResults = append(htmlResults, fmt.Sprintf(
			`<div class="p-4 border rounded-lg mt-2 bg-white shadow-sm flex items-center justify-between">
				<div>
					<p class="text-lg font-semibold text-gray-800">%s</p>
					<p class="text-sm text-gray-500">%s</p>
				</div>
				<button 
					hx-post="/api/restaurant/%s" 
					hx-swap="none"
					class="bg-blue-500 hover:bg-blue-600 text-white text-sm px-4 py-2 rounded-lg transition-all"
				>
					Add
				</button>
			</div>`,
			result.Name, result.Category, result.Id,
		))
	}
	e.Response.Header().Set("Content-Type", "text/html")
	e.Response.WriteHeader(http.StatusOK)
	fmt.Fprint(e.Response, strings.Join(htmlResults, ""))
	return nil
}
