package main

import (
	"html/template"

	"fmt"
	"math/rand"
	"net/http"
	// "text/template"
)

type Card struct {
	Number int
	Expire string
	CVC    int
}

type PaymentInfo struct {
	Name string
	*Card
	lastName          string
	SomeFieldFunction func(int, string) string
}

var info *PaymentInfo = &PaymentInfo{
	Name: "Vladimir",
	Card: &Card{
		Number: 8729123274859234,
		Expire: "10/24",
		CVC:    324,
	},
	lastName: "Kim",
	SomeFieldFunction: func(i int, s string) string {
		return fmt.Sprintf("%d %s", i, s)
	},
}

func (i *PaymentInfo) SomeMethod() string {
	return i.lastName
}

func (i *PaymentInfo) SomeErrorMethod() (string, error) {
	if rand.Int()%2 == 0 {
		return "Hello world!", nil
	}
	return "", fmt.Errorf("Error while executing")
}

func StructHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panic error: %v\n", err)
		}
	}()
	tmpl, err := template.New("struct.html").Funcs(template.FuncMap{
		"myInc": func(i int) int {
			return i + 1
		},
	}).ParseFiles("templates/struct.html")
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		return
	}

	err = tmpl.Execute(w, info)
	if err != nil {
		fmt.Printf("Error while executing template: %v\n", err)
	}
}

func main() {
	http.HandleFunc("/", StructHandler)

	http.ListenAndServe(":8080", nil)
}
