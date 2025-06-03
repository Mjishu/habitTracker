package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type apiConfig struct {
	jwt_secret string;
}

func main() {
	fmt.Println("Hello")

	mux := http.NewServeMux()
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("could not load .env, %s", err)
	}

	jwt_secret := "ya momma"

	config := apiConfig{
		jwt_secret: jwt_secret,
	}
	

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