package main

import (
	"fmt"
	"net/http"
)

func DynamicHandler(w http.ResponseWriter, r *http.Request) {
	firstString := r.FormValue("first")
	secondString := r.FormValue("second")

	resp := fmt.Sprintf(`
<html>
<head></head>
<body>
	<h1>Динамическая страничка</h1>
	<p>	"%s" + "%s" = "%s" </p>
</body>
</html>
	`, firstString, secondString, firstString+secondString)

	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", DynamicHandler)

	http.ListenAndServe(":8080", nil)
}
