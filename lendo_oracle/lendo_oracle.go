package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {
	db, err := sql.Open("godror", `user="hr" password="hr" connectString="localhost:1521/XEPDB1"`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var campoTabela string
	err = db.QueryRow("SELECT department_name FROM departments WHERE department_id = 20").Scan(&campoTabela)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(campoTabela)
}
