package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/cursogo")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, _ := db.Prepare("update usuarios sett nome = ? where id = ?")

	stmt.Exec("Teste ativo", 1)
	stmt.Exec("Valeska", 2)

}
