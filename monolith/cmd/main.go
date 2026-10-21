package main

import (
	"log"

	"github.com/SAB1825/go_wallet/internal/config"
	"github.com/SAB1825/go_wallet/internal/database"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load the env config : %v", err)
	}

	db, err := database.SetupDB(cfg.DBDSN)

	if err != nil {
		log.Fatalf("Failed to setup the db: %v", err)
	}
	defer db.Close()

	log.Println("application started")
}
