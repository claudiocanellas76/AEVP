package main

import (	
	"html/template"
	"net/http"
	"https://github.com/AEVP/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	todosAtletas := models.getAtletas()
	temp.ExecuteTemplate(w, "Index", todosAtletas)
}
