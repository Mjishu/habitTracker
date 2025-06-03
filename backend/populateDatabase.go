package main

import uuid "github.com/google/uuid"

type User struct {
	Id uuid.UUID
	Username string
	Email string
	HashPassword string
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

}