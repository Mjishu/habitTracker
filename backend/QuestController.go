package main

import "net/http"

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