package main

import (
	"fmt"
	"net/http"
)

func (conf apiConfig) UserById(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	fmt.Println("this is from inside of userbyid" + userId)

	sql := `
		SELECT id, username, email
		FROM users
		WHERE id = ?
	`
	row := conf.db.QueryRow(sql, userId)

	var userInfo User
	err := row.Scan(&userInfo)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "User: ID=%d, Username=%s, Email=%s", userInfo.Id, userInfo.Username, userInfo.Email)
}

func (conf apiConfig) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Parse form values
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Example SQL insert statement
	sql := `
		INSERT INTO users (username, email, password)
		VALUES (?, ?, ?)
	`

	// Example: execute the SQL statement (assuming you have a db object)
	_, err := conf.db.Exec(sql, username, email, password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created successfully")
}

func (conf apiConfig) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("userID")
	fmt.Println("Deleting a user " + userId)

	sql := `
		DELETE FROM users WHERE id = ?
	`

	_,err := conf.db.Exec(sql, userId)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User deleted successfully")
}