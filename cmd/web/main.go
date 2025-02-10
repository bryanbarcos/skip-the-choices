package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting on server :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
