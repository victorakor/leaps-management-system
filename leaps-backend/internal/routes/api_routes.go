package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"leaps/internal/controllers"
	"leaps/internal/middleware"
	"leaps/internal/repositories"
	"leaps/internal/services"
)

// SetupRoutes configures all API routes
func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo)
	studentService := services.NewStudentService(db)
	resultService := services.NewResultService(db)
	admissionService := services.NewAdmissionService(db)
	quizService := services.NewQuizService(db)
	financeService := services.NewFinanceService(db)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	studentController := controllers.NewStudentController(studentService)
	resultController := controllers.NewResultController(resultService)
	admissionController := controllers.NewAdmissionController(admissionService)
	quizController := controllers.NewQuizController(quizService)
	financeController := controllers.NewFinanceController(financeService)

	// Public routes
	public := router.Group("/api")
	{
		// User routes
		public.POST("/users", userController.CreateUser)
		public.GET("/users/:id", userController.GetUser)
		public.GET("/users/email/:email", userController.GetUserByEmail)
		public.GET("/users", userController.ListUsers)
		public.PUT("/users/:id", userController.UpdateUser)
		public.DELETE("/users/:id", userController.DeactivateUser)
		public.GET("/users/exists/:email", userController.UserExists)

		// Student routes
		public.POST("/students", studentController.CreateStudent)
		public.GET("/students/:id", studentController.GetStudent)
		public.GET("/students/admission/:admission_number", studentController.GetStudentByAdmissionNumber)
		public.PUT("/students/:id", studentController.UpdateStudent)
		public.POST("/students/:id/transfer", studentController.TransferStudent)
		public.DELETE("/students/:id", studentController.DeactivateStudent)

		// Result routes
		public.POST("/results/scores", resultController.SubmitScores)
		public.GET("/results/student/:student_id", resultController.GetStudentResults)
		public.POST("/results/publish", resultController.PublishResults)
		public.POST("/results/lock", resultController.LockResults)
		public.GET("/results/report-card/:student_id", resultController.GenerateReportCard)

		// Admission routes
		public.POST("/admissions/apply", admissionController.SubmitApplication)
		public.POST("/admissions/schedule-interview", admissionController.ScheduleInterview)
		public.POST("/admissions/:id/approve", admissionController.ApproveApplication)
		public.POST("/admissions/:id/reject", admissionController.RejectApplication)
		public.GET("/admissions", admissionController.GetApplications)
		public.GET("/admissions/:id/letter", admissionController.GenerateAdmissionLetter)

		// Quiz routes
		public.POST("/quizzes", quizController.CreateQuiz)
		public.POST("/quizzes/:quiz_id/questions", quizController.AddQuestion)
		public.POST("/quizzes/:quiz_id/start", quizController.StartQuiz)
		public.POST("/quizzes/attempts/:attempt_id/submit", quizController.SubmitQuiz)
		public.GET("/quizzes/results/:student_id", quizController.GetQuizResults)
		public.GET("/quizzes/:id", quizController.GetQuiz)

		// Finance routes
		public.POST("/finance/fees", financeController.CreateFeeStructure)
		public.POST("/finance/payments", financeController.RecordPayment)
		public.GET("/finance/balance/:student_id", financeController.GetStudentBalance)
		public.GET("/finance/payments/:student_id", financeController.GetPaymentHistory)
		public.GET("/finance/debtors", financeController.GetDebtors)
		public.GET("/finance/summary", financeController.GetFinancialSummary)
	}

	// Protected routes (require authentication)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		_ = protected // placeholder for future protected routes
	}

	// Admin routes (require admin role)
	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.RoleMiddleware("admin"))
	{
		_ = admin // placeholder for future admin routes
	}
}
