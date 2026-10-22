package main

import (
	"log"

	"github.com/SAB1825/go_wallet/internal/config"
	"github.com/SAB1825/go_wallet/internal/database"
	userRoutes "github.com/SAB1825/go_wallet/internal/user/routes"
	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	userRoutes.RegisterUserRoutes(r, db)

	// STARTING APPLICATION
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v\n", err)
	}
}
