package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Conctar() (*sql.DB, error) {
	stringConexao := "golang:golang@/golangestudos?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro

	}
	if db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
