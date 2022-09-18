package db

import (
	"database/sql"
<<<<<<< HEAD
=======

	"github.com/lib/pq"
>>>>>>> cb12aa2672af6231fa7efab9f769730f62dd2c81
)

func ConectaBd() *sql.DB {
	conexao := "user=postgres dbname=AEVP password=651414 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
