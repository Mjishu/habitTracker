package main

import (
	"fmt"
	"net/http"
)

func (cf apiConfig) UserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	fmt.Println("this is from inside of userbyid" + userId)
}