package main

import (
	"encoding/json"
	"net/http"

	"github.com/JustinPras/Chirpy/internal/auth"
	"github.com/JustinPras/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email 		string 	`json:"email"`
		Password 	string 	`json:"password`
	}

	type response struct {
		User
	}


	jwtToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(jwtToken, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate JWT", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPwd, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't hash password", err)
		return
	}

	newUserParams := database.UpdateUserParams{
		ID:				userID,
		Email:			params.Email,
		HashedPassword:	hashedPwd,
	}

	user, err := cfg.db.UpdateUser(r.Context(), newUserParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user", err)
		return
	}
	
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:	user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email: user.Email,
		},
	})
}