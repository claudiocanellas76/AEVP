package main

import (
	"aevp/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	todosAtletas := models.GetAtletas()
	temp.ExecuteTemplate(w, "Index", todosAtletas)
}
