package main

import (
	"html/template"
	"log"
	"net/http"
)

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

func init() {
	template1 = template.Must(template.ParseGlob("templates/*"))
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
	/*Execute template, handle error */
	err1 := template1.ExecuteTemplate(w, "index.html", nil)
	HandleError(w, err1)
}

//bookdjmark.gohtml
func bookdjExecute(w http.ResponseWriter, _ *http.Request) {
	/*Execute template, handle error */
	err1 := template1.ExecuteTemplate(w, "bookdjmark.gohtml", nil)
	HandleError(w, err1)
}

/* EXECUTING TEMPLATES END */

/* TEMPLATE DEFINITION ENDING */

func main() {
	/* ----------------- TEMPLATE EXECUTION BEGINNING -------------- */
	http.HandleFunc("/", indexExecute)
	http.HandleFunc("/bookdjmark.gohtml", bookdjExecute)
	/* ---------------- TEMPLATE EXECUTION ENDING ---------------- */
	/* Handle all files in the static path */
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//http.Handle("/static/jpg", http.StripPrefix("/static/jpg", http.FileServer(http.Dir("static/jpg"))))
	/* Listen and serve */
	http.ListenAndServe(":8080", nil)
}
