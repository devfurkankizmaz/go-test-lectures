package main

import (
	"log"
	"net/http"
)

type application struct {
}

func main() {
	// Set up an App config
	app := application{}

	// Get application routes
	mux := app.routes()

	log.Println("Starting the server on :8080")
	// Start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
