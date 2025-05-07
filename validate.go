package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (cfg *apiConfig) handlerValidate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type chirpParams struct {
		Body string `json:"body"`
	}

	type badReq struct {
		Error string `json:"error"`
	}

	type goodReq struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	chirp := chirpParams{}
	err := decoder.Decode(&chirp)
	if err != nil {
		respBody := badReq{
			Error: "Something went wrong",
		}
		data, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}
	
	if len(chirp.Body) > 140 {
		respBody := badReq{
			Error: "Chirp is too long",
		}
		data, err := json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(data)
		return
	}

	goodR := goodReq{
		Valid: true,
	}

	dat, err := json.Marshal(goodR)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
	w.WriteHeader(http.StatusOK)
	w.Write(dat)
}