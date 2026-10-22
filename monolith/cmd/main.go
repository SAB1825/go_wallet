package main

import (
	"log"

	"github.com/SAB1825/go_wallet/internal/config"
	"github.com/SAB1825/go_wallet/internal/database"
	userHandler "github.com/SAB1825/go_wallet/internal/user/handler"
	userRepository "github.com/SAB1825/go_wallet/internal/user/repository"
	userService "github.com/SAB1825/go_wallet/internal/user/service"
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

	// REPOSITORIES
	uRepo := userRepository.NewMySQLUserRepository(db)

	// SERVICES
	uSvc := userService.NewUserService(uRepo)

	// HANDLER
	uHandler := userHandler.NewUserHandler(uSvc)

	// GIN SETUP
	r := gin.Default()

	r.POST("/api/v1/users", uHandler.Reqister)
	r.GET("/api/v1/users/:id", uHandler.GetProfile)
	r.PUT("/api/v1/users/:id", uHandler.UpdateProfile)

	// STARTING APPLICATION
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed to start: %v\n", err)
	}
}
