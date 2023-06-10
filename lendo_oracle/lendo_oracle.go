package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"
)

func main() {

	db, err := sql.Open("godror", "hr/hr@(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=localhost)(PORT=1521))(CONNECT_DATA=(SERVICE_NAME = XEPDB1)))")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var campoTabela string
	err = db.QueryRow("SELECT department_name FROM departments where department_id=20").Scan(&campoTabela)
	if err != nil {
		db.Close()
		log.Fatal(err)

	}
	fmt.Println(campoTabela)

}
