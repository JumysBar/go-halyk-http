package main

import (
	// "html/template"
	"fmt"
	"net/http"
	"text/template"
)

var slice []string

func SliceHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		slice = append(slice, r.FormValue("elem"))
		fmt.Printf("Slice len: %d\n", len(slice))
	}

	tmpl, _ := template.ParseFiles("templates/advanced.html")
	tmpl.Execute(w, slice)
}

func main() {
	http.HandleFunc("/", SliceHandler)

	http.ListenAndServe(":8080", nil)
}
