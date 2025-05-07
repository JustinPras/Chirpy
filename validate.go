package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"slices"
)

func (cfg *apiConfig) handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	
	censoredChirp := censorProfanity(params.Body)

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: censoredChirp,
	})
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