package main

import (
	"fmt"
	"net/http"
)

//** --------------------------
//**		User
//** --------------------------

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

//** --------------------------
//**		Quests
//** --------------------------

func (conf apiConfig) AllQuests(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT * FROM quests;
	`

	row := conf.db.QueryRow(sql)

	var quests []Quest
	err := row.Scan(&quests)
	if err != nil {
		http.Error(w, "Quest table not found", http.StatusNotFound)
		return
	}
}

func (conf apiConfig) QuestById(w http.ResponseWriter, r *http.Request) {
	questId := r.PathValue("QuestID")
	sql := `
		SELECT * FROM quests WHERE id = ?
	`

	row := conf.db.QueryRow(sql, questId)

	var questInfo Quest 

	err := row.Scan(&questInfo)
	if err != nil {
		http.Error(w, "Quest not found", http.StatusNotFound)
		return
	}
}

//** --------------------------
//**		Rewards
//** --------------------------

func (config apiConfig) GetAllRewards(w http.ResponseWriter, r *http.Request) {

}

func (config apiConfig) GetAllQuestRewards(w http.ResponseWriter, r *http.Request) {

}

func (config apiConfig)  GetRewardById(w http.ResponseWriter, r *http.Request) {

}