package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:changeit@/CursoGo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("UPDATE Usuarios SET nome = ? WHERE id = ?")

	// update
	stmt.Exec("Uchinton", 1)
	stmt.Exec("Sheronistone", 2)

	// delete
	stmt2, _ := db.Prepare("DELETE FROM Usuarios WHERE id = ?")
	stmt2.Exec(3)
}
