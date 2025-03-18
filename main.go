package main

import (
	"log"
	"net/http"

	"github.com/bryanbarcos/skip-the-choices/handlers"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	printBanner()
	// get database directory and load pocketbase database
	//app := pocketbase.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Home)
	//mux.HandleFunc("/api/search", handlers.SearchHandler(app))

	//log.Println(app.HasTable("restaurants"))

	log.Println("Starting on server :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}

func printBanner() {
	log.Println(`
=========================================================================================
    ___ __   _  _  ___    _____  _    _  ___     __  _    _   ___   _    __  ___   ___
  / ___|| | / /| ||   \  |__ __|| |  | ||  _|   / _|| |  | | /   \ | |  / _||  _| / __|
 | (__  | |/_/ | || |)|    | |  | |__| || |_   / /  | |__| |/  _  \| | / /  | |_ | (_
  \__ \ |   _  | || __/    | |  |  __  ||  _| | |   |  __  || (_) || || |   |  _| \_ \
  ___) || |\ \ | || |      | |  | |  | || |_   \ \_ | |  | |\     /| | \ \_ | |_  __) |
 |____/ |_| \_\|_||_|      |_|  |_|  |_||___|   \__||_|  |_| \___/ |_|  \__||___||___/
/////////////////////////////////////////////////////////////////////////////////////////
=========================================================================================
    `)
}
