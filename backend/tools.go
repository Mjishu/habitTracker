package main

import (
	"fmt"
	"log"
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