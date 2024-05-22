package router

import (
	"api_teste_go/src/router/rotas"
	"github.com/gorilla/mux"
)

// Gerar rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
