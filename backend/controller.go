package main

import (
	"fmt"
	"net/http"
)

//** --------------------------
//**		General
//** --------------------------



func (conf apiConfig) CreateController(w http.ResponseWriter, r *http.Request, tableName string, values []string) {


	// Example SQL insert statement
	sql := `
		INSERT INTO ? (?)
		VALUES (?, ?, ?)
	`

	placeholders := ""
	for i := range values {
		if i > 0 {
			placeholders += ", "
		}
		placeholders += "?"
	}

	// Example: execute the SQL statement (assuming you have a db object)
	_, err := conf.db.Exec(sql, tableName,placeholders, username, email, password)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User created successfully")
		respondWithJSON(w, 200, "successful")

}

func (conf apiConfig) DeleteController(w http.ResponseWriter, r *http.Request, table string) {
	userId := r.PathValue("userID")
	fmt.Println("Deleting a user " + userId)

	sql := `
		DELETE FROM ? WHERE id = ?
	`

	_,err := conf.db.Exec(sql,table, userId)
	if err != nil {
		http.Error(w, "Failed to delete row", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "row deleted successfully")
	respondWithJSON(w, 200, "successful")
}

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
	respondWithJSON(w, 200, userInfo)
}


//** --------------------------
//**		Quests
//** --------------------------

func (conf apiConfig) AllQuests(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT * FROM quests;
	`

	rows, err := conf.db.Query(sql)
	if err != nil {
		http.Error(w, "Quest table not found", http.StatusNotFound)
		return
	}
	defer rows.Close()

	var quests []Quest
	for rows.Next() {
		var quest Quest
		err := rows.Scan(&quest)
		if err != nil {
			http.Error(w, "unable to scan quests", http.StatusNotFound)
			return
		}
		quests = append(quests, quest)
	}

	respondWithJSON(w, 200, quests)
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

	respondWithJSON(w, 200, questInfo)
}

//** --------------------------
//**		Rewards
//** --------------------------

func (conf apiConfig) GetAllRewards(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT * FROM items;
	`

	rows, err := conf.db.Query(sql)
	if err != nil {
		http.Error(w, "Reward table not found", http.StatusBadRequest)
	}
	defer rows.Close()

	var rewards []Item

	for rows.Next() {
		var reward Item
		err := rows.Scan(&reward)
		if err != nil {
			http.Error(w, "unable to scan quests", http.StatusNotFound)
			return
		}
		rewards = append(rewards, reward)
	}
	
	respondWithJSON(w, 200, rewards)
}

func (conf apiConfig) GetAllQuestRewards(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT 
			items.*, 
			quest_rewards.*
		FROM 
			quest_rewards
		JOIN 
			items ON quest_rewards.item_id = items.id
		WHERE 
			quest_rewards.quest_id = ?;
	`
	questId := r.PathValue("questId")

	rows, err := conf.db.Query(sql, questId)
	if err != nil {
		http.Error(w, "unable to find quest rewards", http.StatusNotFound)
		return
	}
	defer rows.Close()

	var rewards []struct {
		Item        Item
		QuestReward QuestReward
	}
	for rows.Next() {
		var item Item
		var questReward QuestReward
		err := rows.Scan(
			&item.Id, &item.Name, &item.Description, // adjust fields as per your Item struct
			&questReward.Id, &questReward.QuestId, &questReward.ItemId, &questReward.Amount, // adjust fields as per your QuestReward struct
		)
		if err != nil {
			http.Error(w, "unable to scan quest rewards", http.StatusBadRequest)
			return
		}
		rewards = append(rewards, struct {
			Item        Item
			QuestReward QuestReward
		}{item, questReward})
	}

	respondWithJSON(w, 200, rewards)
}

func (conf apiConfig)  GetRewardById(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT * FROM items WHERE id = ?
	`

	rewardId := r.PathValue("rewardId")
	row := conf.db.QueryRow(sql, rewardId)

	var reward Item
	err := row.Scan(&reward)
	if err != nil {
		http.Error(w, "could not item ", http.StatusBadRequest)
		return
	}

	respondWithJSON(w, 200, reward)
}

//** --------------------------
//**		Enemies
//** --------------------------

func (conf apiConfig) GetAllEnemies(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT * FROM enemies`

	rows, err := conf.db.Query(sql)
	if err != nil {
		respondWithError(w, 400, "failed to query all enemies", err)
		return
	}
	defer rows.Close()

	var enemies []Enemy
	for rows.Next() {
		var enemy Enemy
		err := rows.Scan(&enemy)
		if err != nil {
			respondWithError(w, 400, "could not scan enemy", err)
			return
		}
		enemies = append(enemies, enemy)
	}

	respondWithJSON(w, 200, enemies)
}

func (conf apiConfig) GetEnemyById(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT * FROM enemies WHERE id = ?`

	enemyId := r.PathValue("enemyId")
	row := conf.db.QueryRow(sql, enemyId)

	var enemy Enemy
	err := row.Scan(&enemy)
	if err != nil {
		respondWithError(w, 400, "could not scan enemy", err)
		return
	}
	respondWithJSON(w, 200, enemy)
}

//** --------------------------
//**		Enemy Rewards
//** --------------------------

func (conf apiConfig) GetSpecificEnemyRewards(w http.ResponseWriter, r *http.Request) {
	sql := `
		SELECT items.*, enemy_rewards.*
		FROM enemy_rewards
		JOIN items ON enemy_rewards.item_id = items.id
		WHERE enemy_rewards.enemy_id = ?
	`
	enemyId := r.PathValue("enemyId")
	rows, err := conf.db.Query(sql, enemyId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not find based on the query given the id", err)
		return
	}
	defer rows.Close()

	var rewards []struct {
		Item        Item
		EnemyReward EnemyReward
	}
	for rows.Next() {
		var item Item
		var enemyReward EnemyReward
		err := rows.Scan(
			&item.Id, &item.Name, &item.Description, // adjust fields as per your Item struct
			&enemyReward.Id, &enemyReward.EnemyId, &enemyReward.ItemId, &enemyReward.Amount, // adjust fields as per your EnemyReward struct
		)
		if err != nil {
			http.Error(w, "unable to scan enemy rewards", http.StatusBadRequest)
			return
		}
		rewards = append(rewards, struct {
			Item        Item
			EnemyReward EnemyReward
		}{item, enemyReward})
	}

	respondWithJSON(w, 200, rewards)
}

//** --------------------------
//**		Items
//** --------------------------

func (conf apiConfig) GetAllItems(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT * FROM items`

	rows, err := conf.db.Query(sql)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not find items", err)
		return
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		err := rows.Scan(&item)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "could not scan item", err)
			return
		}
		items = append(items, item)
	}
	respondWithJSON(w, 200, items)
}

func (conf apiConfig) GetItemById(w http.ResponseWriter, r *http.Request) {
	sql := `SELECT * FROM items WHERE id = ?`

	itemId := r.PathValue("itemId")
	row := conf.db.QueryRow(sql, itemId)
	
	var item Item
	err := row.Scan(&item)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "could not scan item", err)
		return
	}
	respondWithJSON(w, 200, item)
}