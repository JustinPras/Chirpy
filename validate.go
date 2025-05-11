package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"slices"
	"fmt"
	"time"

	"github.com/JustinPras/Chirpy/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body 	string 		`json:"body"`
		UserId 	uuid.UUID	`json:"user_id"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	chirpBody, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp", err)
		return
	}

	chirpParams := database.CreateChirpParams {
		Body: 	chirpBody,
		UserID: params.UserId,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create new chirp", err)
		return
	}


	type returnVals struct {
		ID 			uuid.UUID 	`json:"id"`
		CreatedAt 	time.Time 	`json:"created_at"`
		UpdatedAt 	time.Time 	`json:"updated_at"`
		Body 		string 		`json:"body"`
		UserID		uuid.UUID 	`json:"user_id"`
	}

	respondWithJSON(w, http.StatusCreated, returnVals{
		ID:			chirp.ID,
		CreatedAt: 	chirp.CreatedAt,
		UpdatedAt: 	chirp.UpdatedAt,
		Body: 		chirp.Body,
		UserID:		chirp.UserID,
	})

}

func validateChirp(chirp string) (string, error) {
	const maxChirpLength = 140
	if len(chirp) > maxChirpLength {
		return "", fmt.Errorf("Chirp is too long")	
	}

	censoredChirp := censorProfanity(chirp)
	return censoredChirp, nil
}

func censorProfanity(chirp string) string {
	profanity := []string{"kerfuffle", "sharbert", "fornax"}
	const censor = "****"

	words := strings.Split(chirp, " ")
	for i, word := range(words) {
		if slices.Contains(profanity, strings.ToLower(word)) {
			words[i] = censor
		}
	}

	censoredChirp := strings.Join(words, " ")
	return censoredChirp
}