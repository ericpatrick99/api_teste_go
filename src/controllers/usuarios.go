package controllers

import (
	"api_teste_go/src/banco"
	"api_teste_go/src/modelos"
	"api_teste_go/src/repositorios"
	"encoding/json"
	"io"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro != nil {
		return
	}
	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		return
	}

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	repositorio.Criar(usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BUscando todos os usuarios"))
}

func BuscaUsuarioPeloId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuario"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualizando um usuario"))
}

func DeletaUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleta um usuario"))
}
