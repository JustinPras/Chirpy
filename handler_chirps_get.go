package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/JustinPras/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	
	var err error

	authorIDString := r.URL.Query().Get("author_id")
	authorID, err := uuid.Parse(authorIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Author ID", err)
		return
	}

	var dbChirps []database.Chirp
	

	if authorID != uuid.Nil {
		dbChirps, err = cfg.db.GetChirpsForUserOrderByCreatedAtAsc(r.Context(), authorID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve Chirps", err)
			return
		}

	} else {
		dbChirps, err = cfg.db.GetChirpsOrderByCreatedAtAsc(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve Chirps", err)
			return
		}
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

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp ID", err)
		return
	}

	dbChirp, err := cfg.db.GetChirpByID(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp ID does not exist", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:			dbChirp.ID,
		CreatedAt:	dbChirp.CreatedAt,
		UpdatedAt:	dbChirp.UpdatedAt,
		Body:		dbChirp.Body,
		UserID:		dbChirp.UserID,
	})
}