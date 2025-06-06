package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func GetItemFromEnv(key string) string {
	err := godotenv.Load("/.env")
	if err != nil {
		err = godotenv.Load("../.env")
		if err != nil {
			err = godotenv.Load(".env")
			if err != nil {
				fmt.Printf("could not load .env")
			}
		}
	}

	dbURL := os.Getenv(key)
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set in .env")
	}
	return dbURL
}

//* ------------------
//* 		JSON
//* ------------------


func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

//* ------------------
//*	POST DATA
//* ------------------


func checkBody(w http.ResponseWriter, r *http.Request, user interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to read body", http.StatusInternalServerError)
		return nil
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, user)
	if err != nil {
		http.Error(w, "unable to unmarshal json", http.StatusInternalServerError)
		return err
	}
	return nil
}
