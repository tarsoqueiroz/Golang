package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:changeit@/")
	if err != nil {
		fmt.Println("")
		panic(err)
	}
	defer db.Close()

	exec(db, "CREATE DATABASE IF NOT EXISTS CursoGo")
	exec(db, "USE CursoGo")
	exec(db, "DROP TABLE IF EXISTS Usuarios")
	exec(db, `CREATE TABLE Usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
		)`)
}
