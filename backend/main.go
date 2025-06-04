package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tursodatabase/go-libsql"
)

type apiConfig struct {
	jwt_secret string;
	context context.Context;
	connector *libsql.Connector
	db *sql.DB
}

func main() {
	fmt.Println("Hello")

	mux := http.NewServeMux()
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("could not load .env, %s", err)
	}

	jwt_secret := "ya momma" //update this store jwt_secret in .env

	config := apiConfig{ // might need to add empty context and empty pool here
		jwt_secret: jwt_secret,
	}

	// config.CreateConnection()
	config.PopulateTables()
	

	mux.HandleFunc("GET /users/{userID}", func(w http.ResponseWriter, r *http.Request) {
		config.UserById(w, r)
		fmt.Println("called user for specific id")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("listening on port" + port)
	err = http.ListenAndServe(":" + port, mux)
	if err != nil {
		log.Fatal(err)
	}
}