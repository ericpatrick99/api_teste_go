package banco

import (
	"api_teste_go/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConexao)
	if erro != nil {
		return nil, erro
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, erro
	}
	return db, erro
}
