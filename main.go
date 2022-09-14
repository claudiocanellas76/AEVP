package main

import (
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	db := db.conectaBd()
	getAtleta, err := db.Query("select * from atleta")
	if err != nil {
		panic(err.Error())
	}

	a := Atleta{}
	as := []Atleta{}

	for getAtleta.Next() {

		var atletaId int
		var nome string

		err = getAtleta.Scan(&atletaId, &nome)
		if err != nil {
			panic(err.Error())
		}
		a.Id = atletaId
		a.Nome = nome

		as = append(as, a)

	}

	temp.ExecuteTemplate(w, "Index", as)
	defer db.Close()
}
