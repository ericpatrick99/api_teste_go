package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	StringConexao = ""
	Porta         = 0
)

func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar o configurar")
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	StringConexao = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_IP"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NOME"),
	)
}
