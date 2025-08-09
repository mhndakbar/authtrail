package main

import (
	"encoding/json"
	"net/http"

	"github.com/mhndakbar/authtrails/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type userParams struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		decoder := json.NewDecoder(r.Body)
		params := userParams{}
		err := decoder.Decode(&params)
		if err != nil {
			http.Error(w, "Error decoding parameters", http.StatusBadRequest)
			return
		}

		user, err := apiCfg.DB.GetUserByNameAndPassword(r.Context(), database.GetUserByNameAndPasswordParams{
			Name:     params.Name,
			Password: params.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		handler(w, r, user)
	}
}
