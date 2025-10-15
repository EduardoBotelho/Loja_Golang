package db

import(

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConectaDb() *sql.DB {
	conexao := "root:12345678@tcp(localhost:3306)/loja"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
