package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	Uri          string
	Metodo       string
	Funcao       func(http.ResponseWriter, *http.Request)
	Autenticacao bool
}

func ConfigRouter(r *mux.Router) {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Funcao).Methods(route.Metodo)
	}
}
