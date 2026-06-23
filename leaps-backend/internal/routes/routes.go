package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"leaps/internal/controllers"
	"leaps/internal/middleware"
)

func RegisterRoutes(router *gin.Engine, db *sql.DB) {
	// Auth routes
	authCtrl := controllers.NewAuthController(db)

	auth := router.Group("/api/auth")
	{
		auth.POST("/login", authCtrl.Login)
		auth.POST("/register", authCtrl.Register)
		auth.GET("/me", middleware.AuthMiddleware(), authCtrl.GetCurrentUser)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Students
		protected.GET("/students", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Students endpoint"})
		})
		protected.POST("/students", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Student created"})
		})

		// Admissions
		protected.POST("/admissions/apply", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Application submitted"})
		})
		protected.GET("/admissions", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Admissions list"})
		})

		// Results
		protected.POST("/scores/submit", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Scores submitted"})
		})
		protected.GET("/results/student/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Student results"})
		})

		// Quizzes
		protected.POST("/quizzes", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Quiz created"})
		})
		protected.POST("/quizzes/:id/start", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Quiz started"})
		})

		// Finance
		protected.POST("/fees", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Fee created"})
		})
		protected.POST("/payments", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Payment recorded"})
		})

		// Reports
		protected.GET("/reports/student/:id", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Student report"})
		})
		protected.POST("/reports/generate", func(c *gin.Context) {
			c.JSON(201, gin.H{"message": "Report generated"})
		})
	}

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
}
