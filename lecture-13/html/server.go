package main

import (
	"embed"
	"io/ioutil"
	"net/http"
)

//go:embed static
var res embed.FS

func main() {
	http.HandleFunc("/first",
		func(w http.ResponseWriter, r *http.Request) {
			htmlBody := `
			<!DOCTYPE html>
			<html>
			  <head>
				  <title>ШОК!!!Вы не поверите...</title>
			  </head>
			  <body>
			  	<h1>ШОК!!! КОНТЕНТ!!!</h1>
				<h2>Лектор Владимир до сих пор скрывал от вас <a href="https://rb.gy/ttftv7">ЭТО</a>...</h2>
			  </body>
			</html>
			`
			w.Write([]byte(htmlBody))
		})
	http.HandleFunc("/second",
		func(w http.ResponseWriter, r *http.Request) {
			body, _ := ioutil.ReadFile("static/index.html")
			w.Write(body)
		})
	http.Handle("/", http.FileServer(http.Dir("static")))

	// http.Handle("/", http.FileServer(http.FS(res)))
	http.ListenAndServe(":8080", nil)
}
