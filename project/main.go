package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

/* Struct/type definition */
type formdata struct {
	Firstname   string
	Lastname    string
	Email       string
	EmailSubmit int
}

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

/* FUNCMAP DEFINITION */
var funcMap = template.FuncMap{
	"upperCase": strings.ToUpper, //upperCase is a key we can call inside of the template html file
}

func init() {
	//template1 = template.Must(template.ParseGlob("templates/*"))
	template1 = template.Must(template.New("").Funcs(funcMap).ParseGlob("templates/*html"))
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

}

func formSubmit(data formdata) {
	//Execute Template with our data
	fmt.Println("Printing data to our webpage.")
	err := template1.ExecuteTemplate(os.Stdout, "bookdjmark.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

/* EXECUTING TEMPLATES END */

/* TEMPLATE DEFINITION ENDING */

func main() {
	/* ----------------- TEMPLATE EXECUTION BEGINNING -------------- */
	http.HandleFunc("/", indexExecute)
	http.HandleFunc("/bookdjmark.gohtml", bookdjExecute)
	http.HandleFunc("/okayFormNames", okayFormNames)
	http.HandleFunc("/okayFormEmail", okayFormEmail)
	http.HandleFunc("/okayFormBigMessage", okayFormBigMessage)
	http.HandleFunc("/okayFormNumbers", okayFormBigMessage)
	http.Handle("/favicon.ico", http.NotFoundHandler()) //For missing FavIcon
	/* ---------------- TEMPLATE EXECUTION ENDING ---------------- */
	/* Handle all files in the static path */
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//http.Handle("/static/jpg", http.StripPrefix("/static/jpg", http.FileServer(http.Dir("static/jpg"))))
	/* Listen and serve */
	http.ListenAndServe(":8080", nil)
}

/* Check to see if Form they submitted is kosher */
func okayFormNames(w http.ResponseWriter, req *http.Request) {
	//Here we will spell out ALL the no no inputs

	/* debugging purposes */
	//fmt.Printf("Here is the Head:\n %v\n And here is the body:\n%v\n", req.Header, req.Body)
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)

	//fmt.Println("USERNAME: ", sbs) DEBUG

	fmt.Fprint(w, sbs)
}

func okayFormEmail(w http.ResponseWriter, req *http.Request) {
	/* debugging purposes */
	//fmt.Printf("Here is the Head:\n %v\n And here is the body:\n%v\n", req.Header, req.Body) DEBUG
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	//fmt.Printf("Email value is: %v\n", sbs) DEBUG
	//Check to see if there's anything bad about our email string
	if len(sbs) <= 0 {
		sbs = "Please enter a valid email address" //Alert User that NOTHING is in here
	} else if !strings.Contains(sbs, "@") {
		sbs = "Please enter a valid email address" //Alert User that they are missing an @ symbol
	} else if !strings.Contains(sbs, ".") {
		sbs = "Please enter a valid email address" //Alert User that they are missing a period symbol
	} else if strings.Contains(sbs, " ") {
		sbs = "Please enter a valid email address" //Alert User that they have a space
	} else if strings.Contains(sbs, "\n") {
		sbs = "Please enter a valid email address" //Alert User that they have a newline character
	} else {
		//fmt.Println("USERNAME: ", sbs) //Debugging Purposes
	}
	fmt.Fprint(w, sbs)
}

func okayFormBigMessage(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	//fmt.Printf("Email value is: %v\n", sbs) DEBUG
	//Check to see if there's anything bad about our email string
	if len(sbs) <= 0 {
		sbs = "Please leave me a message and subject!" //Alert User that NOTHING is in here
	}
	fmt.Fprint(w, sbs)
}

func okayFormNumbers(w http.ResponseWriter, req *http.Request) {
	bs, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}

	sbs := string(bs)
	_, err = strconv.Atoi(sbs)
	if err != nil {
		sbs = "Please enter a phone number so I can get back to you!" //Alert User that NOTHING is in here
	} else if len(sbs) <= 0 {
		sbs = "Please enter a phone number so I can get back to you!" //Alert User that NOTHING is in here
	} else {

	}

	fmt.Fprint(w, sbs)
}
