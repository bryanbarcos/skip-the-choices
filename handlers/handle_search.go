package handlers

import (
	//"log"
	"fmt"
	"net/http"
	"strings"

	"github.com/pocketbase/pocketbase"
	//"github.com/bryanbarcos/skip-the-choices/ui"
)

type SearchRestaurantRequest struct {
	Query string `json:"query"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request, app *pocketbase.PocketBase) {
	query := r.URL.Query().Get("query")
	print("Query: {}", query)

	dbResults := []string{"Restaurant1", "Restaurant2"}

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
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, strings.Join(htmlResults, ""))
	//}

}
