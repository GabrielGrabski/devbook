package rotas

import (
	"api-go/src/controller"
	"net/http"
)

var userRoutes = []Rota{
	{
		Uri:          "/usuarios",
		Metodo:       http.MethodPost,
		Funcao:       controller.SaveUsuario,
		Autenticacao: false,
	},
	{
		Uri:          "/usuarios",
		Metodo:       http.MethodGet,
		Funcao:       controller.GetAllUsuarios,
		Autenticacao: false,
	},
	{
		Uri:          "/usuarios/{id}",
		Metodo:       http.MethodGet,
		Funcao:       controller.GetUsuarioById,
		Autenticacao: false,
	},
	{
		Uri:          "/usuarios/{id}",
		Metodo:       http.MethodPut,
		Funcao:       controller.UpdateUsuario,
		Autenticacao: false,
	},
	{
		Uri:          "/usuarios/{id}",
		Metodo:       http.MethodDelete,
		Funcao:       controller.DeleteUsuario,
		Autenticacao: false,
	},
}
