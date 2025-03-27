package main

import (
	"golangProject.com/private"
	"log"
	"net/http"
)

func buildHTTP() {
	// Register the handler
	http.HandleFunc("/events", private.EventsHandler)

	// Start the server
	port := ":8080"
	log.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
