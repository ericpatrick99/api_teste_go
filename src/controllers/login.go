package controllers

import (
	"api_teste_go/src/banco"
	"api_teste_go/src/modelos"
	"encoding/json"
	"io"
	"net/http"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	corpoRequisicao, erro := io.ReadAll(request.Body)
	if erro != nil {
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		return
	}
	defer db.Close()

	//repositorio := repositorio.NovoRepositorioDeUsuario(db)
}
