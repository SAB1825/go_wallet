package routes

import (
	"database/sql"

	"github.com/SAB1825/go_wallet/internal/user/handler"
	"github.com/SAB1825/go_wallet/internal/user/repository"
	"github.com/SAB1825/go_wallet/internal/user/service"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, db *sql.DB) {
	repo := repository.NewMySQLUserRepository(db)
	svc := service.NewUserService(repo)
	h := handler.NewUserHandler(svc)

	g := r.Group("/api/v1/users")
	g.POST("", h.Register)
	g.GET(":id", h.GetProfile)
	g.PUT(":id", h.UpdateProfile)
}
