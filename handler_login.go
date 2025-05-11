package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/JustinPras/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password 	string 	`json:"password`
		Email 		string 	`json:"email"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email", err)
		return
	}

	err = auth.CheckPasswordHash(user.HashedPassword, params.Password)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect password", err)
		return
	}

	type returnVals struct {
		Id 			uuid.UUID 	`json:"id"`
		CreatedAt 	time.Time 	`json:"created_at"`
		UpdatedAt 	time.Time 	`json:"updated_at"`
		Email 		string 		`json:"email"`
	}


	respondWithJSON(w, http.StatusOK, returnVals{
		Id:			user.ID,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		Email: 		user.Email,
	})
}