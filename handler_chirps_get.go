package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	
	dbChirps, err := cfg.db.GetChirpsOrderByCreatedAtAsc(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve Chirps", err)
		return
	}

	chirps := []Chirp{}

	for _, chirp := range(dbChirps) {
		chirps = append(chirps, Chirp{
			ID: 		chirp.ID,
			CreatedAt: 	chirp.CreatedAt,
			UpdatedAt:	chirp.UpdatedAt,
			Body:		chirp.Body,
			UserID:		chirp.UserID,
		})
	}

	respondWithJSON(w, 200, chirps)
}