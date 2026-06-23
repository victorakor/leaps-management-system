package routes

import (
	"database/sql"

	"leaps/internal/controllers"
	"leaps/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	authCtrl := controllers.NewAuthController(db)
	auth := router.Group("/api/auth")
	{
		auth.POST("/login", authCtrl.Login)
		auth.POST("/register", authCtrl.Register)
		auth.GET("/me", middleware.AuthMiddleware(), authCtrl.GetCurrentUser)
	}
}
