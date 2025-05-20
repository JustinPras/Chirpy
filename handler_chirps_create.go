package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"slices"
	"fmt"
	"time"

	"github.com/JustinPras/Chirpy/internal/database"
	"github.com/JustinPras/Chirpy/internal/auth"

	"github.com/google/uuid"
)

type Chirp struct {
	ID 			uuid.UUID 	`json:"id"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	Body 		string 		`json:"body"`
	UserID		uuid.UUID 	`json:"user_id"`
}

func (cfg *apiConfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body 	string 		`json:"body"`
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

	chirpBody, err := validateChirp(params.Body)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Chirp", err)
		return
	}

	chirpParams := database.CreateChirpParams {
		Body: 	chirpBody,
		UserID:	userID,
	}

	chirp, err := cfg.db.CreateChirp(r.Context(), chirpParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create new chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, Chirp{
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