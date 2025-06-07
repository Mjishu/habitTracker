package Interactions

import (
	"database/sql"
	"fmt"
)

type PlayerInteractions string 

const (
	Create PlayerInteractions = "Create"
	Attack PlayerInteractions = "Attack"
	Defend PlayerInteractions = "Defend"
	Flee   PlayerInteractions = "Flee"
)

func EnemyPlayerInteraction(interactionType PlayerInteractions, db *sql.DB) {
	switch interactionType {
	case Create:
		fmt.Println("enemy interaction is create")
	case Attack:
		fmt.Println("enemy interaction is attack")
	case Defend:
		fmt.Println("enemy interaction is Defend")
	case Flee:
		fmt.Println("enemy interaction is Flee")
	}
}