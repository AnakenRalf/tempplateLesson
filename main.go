package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	Data string
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/answer", answerHandler)

	fmt.Println("Starting server on: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	http.Redirect(w, r, "/answer?data="+data, http.StatusSeeOther)
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("data")

	tmpl, err := template.ParseFiles("templates/answer.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pageData := PageData{Data: data}
	tmpl.Execute(w, pageData)
}
