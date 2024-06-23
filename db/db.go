package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func ConexaoDB() *sql.DB {
	conexao := fmt.Sprintf("%s:%s@tcp(localhost:3306)/db_loja_golang", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"))
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
