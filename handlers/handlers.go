package handlers

import (
	//"log"
	"fmt"
	"net/http"
	"strings"

	"github.com/bryanbarcos/skip-the-choices/ui"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	ui.MainTempl().Render(r.Context(), w)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	print("Query: {}", query)
	results := []string{"Restaurant A", "Restaurant B", "Restaurant C"}

	var htmlResults []string
	for _, result := range results {
		htmlResults = append(htmlResults, fmt.Sprintf(
			`<div class="p-2 border rounded mt-2 fade-in">%s</div>`, result,
		))
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, strings.Join(htmlResults, ""))
}
