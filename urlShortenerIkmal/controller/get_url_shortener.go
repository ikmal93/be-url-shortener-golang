package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"goland/saviory-go/config/utility"
	"goland/urlShortenerIkmal/config"
	"net/http"
)

func GetUrlShortener(router *mux.Router) *mux.Router {
	router.HandleFunc("/{slug}", GetUrlShortenerFunc).
		Methods("GET")
	return router
}

func GetUrlShortenerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slug := mux.Vars(r)["slug"]

	var db = config.ConnectDB()
	defer db.Close()

	var rows, err = db.Query("select * from redirect where slug = ?",
		slug)
	if err != nil {
		json.NewEncoder(w).Encode(utility.ErrorResponse{
			Code:    200,
			Error:   true,
			Message: err.Error(),
		})
		return
	}
	defer rows.Close()

	var redirect config.Redirect
	for rows.Next() {
		var each = config.Redirect{}
		var err = rows.Scan(
			&each.ID,
			&each.Slug,
			&each.URL,
		)
		if err != nil {
			json.NewEncoder(w).Encode(utility.ErrorResponse{
				Code:    200,
				Error:   true,
				Message: err.Error(),
			})
			return
		}
		redirect = each
	}
	if err = rows.Err(); err != nil {
		json.NewEncoder(w).Encode(utility.ErrorResponse{
			Code:    200,
			Error:   true,
			Message: err.Error(),
		})
		return
	} else {
		if redirect.ID == 0 {
			json.NewEncoder(w).Encode(utility.ErrorResponse{
				Code:    200,
				Error:   true,
				Message: "Resource Not Found",
			})
			return
		} else {
			http.Redirect(w, r, redirect.URL, http.StatusSeeOther)
		}
	}
}
