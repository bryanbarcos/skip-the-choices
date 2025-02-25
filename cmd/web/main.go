package main

import (
	"log"
	"net/http"
)

func main() {
	printBanner()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

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
