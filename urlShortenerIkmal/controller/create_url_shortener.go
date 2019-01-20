package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goland/urlShortenerIkmal/config"
	"net/http"
)

func CreateUrlShortener(router *mux.Router) *mux.Router {
	router.HandleFunc("/create", CreateUrlShortenerFunc).
		Methods("POST")
	return router
}

func CreateUrlShortenerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var db = config.ConnectDB()
	defer db.Close()

	var url = r.FormValue("url")
	slug := config.RandomCharacter(5)

	var err error
	_, err = db.Query("insert into redirect (slug, url) values (?,?)",
		slug, url)

	if err != nil {
		json.NewEncoder(w).Encode(config.Response{
			Code:    200,
			Error:   true,
			Message: err.Error(),
		})
		return
	} else {
		json.NewEncoder(w).Encode(config.Response{
			Code:    201,
			Error:   false,
			Message: "Url Shortener Created : " + "http://localhost:12345/" + slug,
		})
		return
	}
}
