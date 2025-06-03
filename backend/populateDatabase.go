package main

import (
	"fmt"
	"os"

	uuid "github.com/google/uuid"
)

type User struct {
	Id uuid.UUID
	Username string
	Email string
	Hash_password string
	Image_src string
	max_hp float32
	current_hp float32
	base_attack_damage float32
	xp float32
	gold float32
	created_at string
	updated_at string
}

func (conf apiConfig) CreateUserTable() {
	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			username VARCHAR(32) NOT NULL,
			email VARCHAR(96) NOT NULL,
			hash_password TEXT NOT NULL,
			image_src TEXT,
			max_hp FLOAT default 10.0,
			current_hp FLOAT default 10.0,
			base_attack FLOAT default 1.0,
			xp FLOAT default 0.0,
			gold FLOAT default 0.0,
			created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
			updated_at TIMESTAMPTZ DEFAULT now() NOT NULL
		)
	`

	_,err := conf.pool.Exec(conf.ctx, sql)
	queryFail(err, "User")
}

func queryFail(err error, tableName string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Table '%s' created successfully\n", tableName)
}