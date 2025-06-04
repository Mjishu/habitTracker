package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tursodatabase/go-libsql"
)

func (conf apiConfig) CreateConnection() {
	dbName := "habitTracker"
	primaryUrl := GetItemFromEnv("TURSO_DATABASE_URL")
	authToken := GetItemFromEnv("TURSO_DATABASE_TOKEN")
	ctx := context.Background()
	conf.context = ctx

	dir,err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		fmt.Println("error making temp directory: ", err)
		os.Exit(1)
	}
	defer os.RemoveAll(dir)
	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, libsql.WithAuthToken(authToken),)
	if err != nil {
		fmt.Println("error creating DB connection: ", err)
		os.Exit(1)
	}
	conf.connector = connector
	defer connector.Close()

	db := sql.OpenDB(connector)
	conf.db = db
	defer db.Close() //! might have to put these defers in main .go and do defer conf.db.Close()?
	
}

func CreateTables(conf apiConfig) {
	conf.CreateUserTable()
	conf.CreateQuestTable()
	conf.CreateQuestRewardsTable()
	conf.CreateQuestLogsTable()
	conf.CreateItemsTable()
	conf.CreateItemEffectsTable()
	conf.CreateUserItemsTable()
	conf.CreateEnemiesTable()
	conf.CreateEnemyRewardsTable()
	conf.CreateUserEnemyEncountersTable()
	conf.CreateHabitsTable()
	conf.CreateHabitLogsTable()
	conf.CreateCombatLogsTable()
	conf.CreateAchievementsTable()
	conf.CreateUserAchievementsTable()
}


func (conf apiConfig) CreateUserTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			username TEXT NOT NULL,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			image_src TEXT,
			max_hp REAL DEFAULT 10.0,
			current_hp REAL DEFAULT 10.0,
			base_attack_damage REAL DEFAULT 1.0,
			xp REAL DEFAULT 0.0,
			gold REAL DEFAULT 0.0,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "users")
}

func (conf apiConfig) CreateQuestTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS quests (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "quests")
}

func (conf apiConfig) CreateQuestRewardsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS quest_rewards (
			id TEXT PRIMARY KEY,
			quest_id TEXT NOT NULL,
			reward TEXT NOT NULL,
			FOREIGN KEY (quest_id) REFERENCES quests(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "quest_rewards")
}

func (conf apiConfig) CreateQuestLogsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS quest_logs (
			id TEXT PRIMARY KEY,
			quest_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			created_at TEXT NOT NULL,
			completed_at TEXT,
			FOREIGN KEY (quest_id) REFERENCES quests(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "quest_logs")
}

func (conf apiConfig) CreateItemsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS items (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			type TEXT,
			value REAL DEFAULT 0.0,
			description TEXT,
			image_src TEXT,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "items")
}

func (conf apiConfig) CreateItemEffectsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS item_effects (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			item_id INTEGER NOT NULL,
			effect_type TEXT CHECK(type IN ('increase_damage','increase_health', 'increase_defense','increase_speed, 'critical', 'lifesteal', 'xp_boost', 'immune')) NOT NULL
			effect_value TEXT NOT NULL,
			FOREIGN KEY (item_id) REFERENCES items(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "item_effects")
}

func (conf apiConfig) CreateUserItemsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS user_items (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			item_id INTEGER NOT NULL,
			quantity REAL NOT NULL,
			acquired_at TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (item_id) REFERENCES items(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "user_items")
}

func (conf apiConfig) CreateEnemiesTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS enemies (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			max_hp REAL NOT NULL,
			attack REAL NOT NULL,
			defense REAL NOT NULL,
			xp_reward REAL NOT NULL,
			gold_reward REAL NOT NULL,
			description TEXT NOT NULL,
			type TEXT CHECK(type IN ('normal', 'elite', 'boss')) NOT NULL
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "enemies")
}

func (conf apiConfig) CreateEnemyRewardsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS enemy_rewards (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			enemy_id TEXT NOT NULL,
			item_id INTEGER NOT NULL,
			FOREIGN KEY (enemy_id) REFERENCES enemies(id),
			FOREIGN KEY (item_id) REFERENCES items(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "enemy_rewards")
}

func (conf apiConfig) CreateUserEnemyEncountersTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS user_enemy_encounters (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			enemy_id TEXT NOT NULL,
			current_hp REAL NOT NULL,
			encounter_started_at TEXT NOT NULL,
			encounter_status TEXT CHECK(encounter_status IN ('active', 'fled', 'defeated')) NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (enemy_id) REFERENCES enemies(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "user_enemy_encounters")
}

func (conf apiConfig) CreateHabitsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS habits (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			name TEXT NOT NULL,
			description TEXT,
			habit_type TEXT,
			difficulty TEXT,
			created_at TEXT NOT NULL,
			updated_at TEXT NOT NULL,
			effects_stat TEXT,
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "habits")
}

func (conf apiConfig) CreateHabitLogsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS habit_logs (
			id TEXT PRIMARY KEY,
			habit_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			performed_at TEXT NOT NULL,
			effect TEXT,
			notes TEXT,
			mood TEXT,
			FOREIGN KEY (habit_id) REFERENCES habits(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "habit_logs")
}

func (conf apiConfig) CreateCombatLogsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS combat_logs (
			id TEXT PRIMARY KEY,
			enemy_id TEXT NOT NULL,
			user_id TEXT NOT NULL,
			action TEXT CHECK(action IN ('attack', 'flee', 'use_item')) NOT NULL,
			damage_dealt REAL,
			damage_taken REAL,
			hp_after_action REAL,
			created_at TEXT NOT NULL,
			FOREIGN KEY (enemy_id) REFERENCES enemies(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "combat_logs")
}

func (conf apiConfig) CreateAchievementsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS achievements (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT,
			reward_type TEXT,
			reward_value REAL
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "achievements")
}

func (conf apiConfig) CreateUserAchievementsTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS user_achievements (
			id TEXT PRIMARY KEY,
			user_id TEXT NOT NULL,
			achievement_id TEXT NOT NULL,
			achieved_at TEXT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id),
			FOREIGN KEY (achievement_id) REFERENCES achievements(id)
		);
	`
	_, err := conf.db.Exec(sql)
	queryFail(err, "user_achievements")
}


func queryFail(err error, tableName string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Table '%s' created successfully\n", tableName)
}