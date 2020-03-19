package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	/* ------------- SERVING FILES BEGINNING ---------------------- */
	/* Handles all the files in the main.go directory, (everything under project) */
	http.Handle("/", http.FileServer(http.Dir(".")))
	/* Defining a port for Google Cloud environment, (defaults to port 8080 for testing) */
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s ", port)
	}
	/* logs a fatal error if no port is identified to serve these files */
	log.Printf("listenting on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	/* ------------- SERVING FILES ENDING ---------------------- */
}
