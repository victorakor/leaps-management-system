package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"leaps/internal/models"
	"leaps/internal/services"
)

type QuizController struct {
	quizService *services.QuizService
}

func NewQuizController(quizService *services.QuizService) *QuizController {
	return &QuizController{quizService: quizService}
}

func (qc *QuizController) CreateQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := qc.quizService.CreateQuiz(c.Request.Context(), &quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Quiz created", "quiz": created})
}

func (qc *QuizController) AddQuestion(c *gin.Context) {
	quizID := c.Param("quiz_id")
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	question.QuizID = quizID
	created, err := qc.quizService.AddQuestion(c.Request.Context(), &question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Question added", "question": created})
}

func (qc *QuizController) StartQuiz(c *gin.Context) {
	quizID := c.Param("quiz_id")
	var req struct{ StudentID string `json:"student_id" binding:"required"` }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attempt, err := qc.quizService.StartQuiz(c.Request.Context(), quizID, req.StudentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"attempt": attempt})
}

func (qc *QuizController) SubmitQuiz(c *gin.Context) {
	attemptID := c.Param("attempt_id")
	var req struct{ Answers map[string]string `json:"answers"` }
	c.ShouldBindJSON(&req)
	result, err := qc.quizService.SubmitQuiz(c.Request.Context(), attemptID, req.Answers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": result})
}

func (qc *QuizController) GetQuizResults(c *gin.Context) {
	studentID := c.Param("student_id")
	results, err := qc.quizService.GetQuizResults(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})
}

func (qc *QuizController) GetQuiz(c *gin.Context) {
	id := c.Param("id")
	quiz, err := qc.quizService.GetQuiz(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quiz)
}
