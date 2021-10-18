package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/index3.html")
		return
	}

	r.ParseMultipartForm(30 * 1024)
	file, handler, err := r.FormFile("f")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	f, err := os.OpenFile("./uploaded/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		http.ServeFile(w, r, "static/index3.html")
		return
	}

	m, err := r.MultipartReader()
	if err != nil {
		fmt.Println(err)
	}

	for {
		part, err := m.NextPart()
		if err == io.EOF {
			break
		}

		if part.FileName() == "" {
			continue
		}

		dst, err := os.Create("./uploaded/" + part.FileName())
		if err != nil {
			fmt.Println(err)
		}

		io.Copy(dst, part)
	}
}

func main() {
	http.HandleFunc("/", FirstHandler)
	// http.HandleFunc("/", SecondHandler)

	http.ListenAndServe(":8080", nil)
}
