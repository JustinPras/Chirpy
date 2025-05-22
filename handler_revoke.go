package main

import (
	"net/http"

	"github.com/JustinPras/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	type response struct {
	}

	refreshTokenString, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find refresh token", err)
		return
	}

	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't revoke refresh token", err)
		return
	}


	respondWithJSON(w, http.StatusNoContent, response{})
}