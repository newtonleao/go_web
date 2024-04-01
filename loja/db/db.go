package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectaComBancoDeDados() *sql.DB {
	db, err := sql.Open("mysql", "root:teste@tcp(localhost:3306)/alura_loja")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
