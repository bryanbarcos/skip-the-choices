package handlers

import (
	//"log"
	"fmt"
	"strings"

	"github.com/pocketbase/pocketbase/core"
	//"github.com/pocketbase/pocketbase"
	//"github.com/bryanbarcos/skip-the-choices/ui"
)

type SearchRestaurantRequest struct {
	Query string `json:"query"`
}

type RestaurantRecord struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	PrimaryType string `json:"primary_type"`
}

func SearchHandler(e *core.RequestEvent) error {
	app := e.App
	query := e.Request.URL.Query().Get("query")
	fmt.Println("Query: {}", query)

	dbResults := []string{"Restaurant1", "Restaurant2"}

	record, err := app.FindRecordById("restaurants", "f4laixq715ejo1v")
	//dbx.NewExp("LOWER(name) = {:name}", dbx.Params{"name": "test"}))

	if err != nil {
		fmt.Println("error: ", err)
		return err
	}

	fmt.Println("Name = ", record.Get("name"))

	// fmt.Println("Results: {}", )

	// dbResults, err := app.FindAllRecords("restaurants",
	// 	dbx.NewExp("LOWER(name) = {:name}", dbx.Params{"name": "test"}))

	// return func(w http.ResponseWriter, r *http.Request) {
	// 	var req SearchRestaurantRequest
	// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 		http.Error(w, "Invalid request", http.StatusBadRequest)
	// 		return
	// 	}
	// }
	//if err == nil && len(localResults) > 0 {
	// Found restaurant in database, return results
	var htmlResults []string
	for _, result := range dbResults {
		htmlResults = append(htmlResults, fmt.Sprintf(
			`<div class="p-2 border rounded mt-2 fade-in">%s</div>`, result,
		))
	}
	e.Response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(e.Response, strings.Join(htmlResults, ""))
	//}
	return nil
}
