package models

import "github.com/AEVP/db"

type Atleta struct {
	Id   int
	Nome string
}

func getAtletas() []Atleta {

	db := db.ConectaBd()
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
