package routers

import (
	"github.com/gorilla/mux"
	"goland/urlShortenerIkmal/controller"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = controller.CreateUrlShortener(router)
	router = controller.GetUrlShortener(router)
	return router

}
