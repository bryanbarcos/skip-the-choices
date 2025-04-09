package main

import (
	"log"
	"os"

	"github.com/bryanbarcos/skip-the-choices/internal/handlers"
	"github.com/joho/godotenv"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	db_directory := os.Getenv("DATA_DIR")
	printBanner()
	// get database directory and load pocketbase
	app := pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: db_directory,
	})

	// userId := "hm77ox88p0j5t4v"

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/", handlers.Home)
		se.Router.GET("/api/search", handlers.SearchHandler)
		se.Router.POST("/api/restaurant/{id}", handlers.AddRestaurant)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
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
