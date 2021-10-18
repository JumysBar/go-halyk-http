package main

import (
	// "html/template"
	"net/http"
	"text/template"
)

func XssHandler(w http.ResponseWriter, r *http.Request) {
	var args map[string]string
	if r.Method == http.MethodPost {
		args = make(map[string]string)
		args["Name"] = r.FormValue("name")
	}

	tmpl, _ := template.ParseFiles("templates/xss.html")
	tmpl.Execute(w, args)
}

func main() {
	http.HandleFunc("/", XssHandler)

	http.ListenAndServe(":8080", nil)
}
