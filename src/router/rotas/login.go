package rotas

import (
	"api_teste_go/src/controllers"
	"net/http"
)

var rotasLogin = Rotas{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
