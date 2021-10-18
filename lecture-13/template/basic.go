package main

import (
	"html/template"
	"net/http"
)

func DynamicTemplateHandler(w http.ResponseWriter, r *http.Request) {
	firstString := r.FormValue("first")
	secondString := r.FormValue("second")
	data := []string{firstString, secondString, firstString + secondString}
	tmpl, _ := template.ParseFiles("templates/basic.html")
	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", DynamicTemplateHandler)

	http.ListenAndServe(":8080", nil)
}
