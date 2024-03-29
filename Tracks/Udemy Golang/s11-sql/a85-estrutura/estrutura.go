// para instalar o pacote mysql executar:
//    go get -u github.com/go-sql-driver/mysql
//
package main

import (
	"database/sql"

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
	db, err := sql.Open("mysql", "root:root")
	if err != nil {
		panic(err)
	}
	// deixa pendente o close que é executado somente na saída do programa
	// garante que a conexão com o banco será fechado.
	defer db.Close()

	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
	)`)
}
