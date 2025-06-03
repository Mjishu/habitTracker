package main

import (
	"fmt"
	"net/http"
)

func (cf apiConfig) UserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	fmt.Println("this is from inside of userbyid" + userId)
}

func (cf apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating a new user....")
}

func (cf apiConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	fmt.Println("Deleting a user " + userId)
}