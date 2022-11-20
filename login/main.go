package main

import (
	"html/template"
	"net/http"
)

var tp *template.Template

func main() {
	tp, _ = template.ParseFiles("index.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)

}
func indexHandler(w http.ResponseWriter, r *http.Request) {
	tp.Execute(w, nil)
}
