package main

import uuid "github.com/google/uuid"

type User struct {
	Id                uuid.UUID
	Username          string
	Email             string
	Hash_password     string
	Image_src         string
	Max_hp            float32
	Current_hp        float32
	Base_attack_damage float32
	Xp                float32
	Gold              float32
	Created_at        string
	Updated_at        string
}

type Quest struct {
	Id          string  // change to uuid.UUID? 
	Name        string
	Description string
}

type QuestReward struct {
	Id      uuid.UUID
	QuestId uuid.UUID
	Reward  string
}

type QuestLog struct {
	Id         uuid.UUID
	QuestId    uuid.UUID
	UserId     uuid.UUID
	CreatedAt  string
	CompletedAt *string
}

type Item struct {
	Id          int64
	Name        string
	Type        string
	Value       float32
	Description string
	Image_src   string
	Created_at  string
	Updated_at  string
}

type ItemEffect struct {
	Id          int64
	ItemId      int64
	Effect_type string
	Effect_value string
}

type UserItem struct {
	Id         uuid.UUID
	UserId     uuid.UUID
	ItemId     int64
	Quantity   float32
	AcquiredAt string
}

type Enemy struct {
	Id         uuid.UUID
	Name       string
	Max_hp     float32
	Attack     float32
	Defense    float32
	Xp_reward  float32
	Gold_reward float32
	Description string
	Type       string
}

type EnemyReward struct {
	Id      int64
	EnemyId uuid.UUID
	ItemId  int64
}

type UserEnemyEncounter struct {
	Id                 uuid.UUID
	UserId             uuid.UUID
	EnemyId            uuid.UUID
	Current_hp         float32
	EncounterStartedAt string
	EncounterStatus    string
}

type Habit struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	Name        string
	Description string
	HabitType   string
	Difficulty  string
	Created_at  string
	Updated_at  string
	EffectsStat string
}

type HabitLog struct {
	Id         uuid.UUID
	HabitId    uuid.UUID
	UserId     uuid.UUID
	PerformedAt string
	Effect     string
	Notes      string
	Mood       string
}

type CombatLog struct {
	Id            uuid.UUID
	EnemyId       uuid.UUID
	UserId        uuid.UUID
	Action        string
	DamageDealt   *float32
	DamageTaken   *float32
	HpAfterAction *float32
	CreatedAt     string
}

type Achievement struct {
	Id          uuid.UUID
	Name        string
	Description string
	RewardType  string
	RewardValue *float32
}

type UserAchievement struct {
	Id            uuid.UUID
	UserId        uuid.UUID
	AchievementId uuid.UUID
	AchievedAt    string
}
