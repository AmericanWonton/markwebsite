package main

import (
	"fmt"
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
func bookdjExecute(w http.ResponseWriter, req *http.Request) {
	/*Execute template, handle error */
	err1 := template1.ExecuteTemplate(w, "bookdjmark.gohtml", nil)
	HandleError(w, err1)
	var fName string = string(req.FormValue("firstname"))
	var lName string = string(req.FormValue("lastname"))
	var email string = string(req.FormValue("email"))

	/* Handle Forms submitted to me for emails */
	if (len(fName) == 0) && (len(lName) == 0) && (len(email) == 0) {
		/* Bad form fillout, instruct User to try again. */
		fmt.Println("Hey, it's empty for everything")
	} else {
		/* Things are filled out, need to check if they're okay for an email to me with proper
		formatting */
		goodForm := okayForm(fName, lName, email)

		if goodForm == true {
			/* Send the form, everything is good */
		} else {
			/* User should have been warned, per the 'okayForm' function */
		}
	}
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

/* Check to see if Form they submitted is kosher */
func okayForm(fName string, lName string, email string) bool {
	goodForm := true
	fmt.Printf("fName is: %v\n", fName)
	fmt.Printf("lName is: %v\n", lName)
	fmt.Printf("email is: %v\n", email)

	return goodForm
}
