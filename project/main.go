package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

func init() {
	template1 = template.Must(template.ParseGlob("gotempls/*"))
}

/* EXECUTING TEMPLATES BEGINNING */
// Handle Errors
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//index.gohtml
func indexExecute(w http.ResponseWriter, _ *http.Request) {
	err1 := template1.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err1)
}

/* EXECUTING TEMPLATES END */

/* TEMPLATE DEFINITION ENDING */

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

	/* ----------------- TEMPLATE EXECUTION BEGINNING -------------- */
	http.HandleFunc("/", indexExecute)
}
