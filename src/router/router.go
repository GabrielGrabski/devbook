package router

import (
	"api-go/src/router/rotas"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	rotas.ConfigRouter(r)
	return r
}
