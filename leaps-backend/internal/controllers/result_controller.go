package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"leaps/internal/models"
	"leaps/internal/services"
)

type ResultController struct {
	resultService *services.ResultService
}

func NewResultController(resultService *services.ResultService) *ResultController {
	return &ResultController{resultService: resultService}
}

func (rc *ResultController) SubmitScores(c *gin.Context) {
	var req struct {
		StudentID string  `json:"student_id" binding:"required"`
		SubjectID string  `json:"subject_id" binding:"required"`
		TermID    string  `json:"term_id" binding:"required"`
		CA1       float64 `json:"ca1"`
		CA2       float64 `json:"ca2"`
		Exam      float64 `json:"exam"`
		Remark    string  `json:"remark"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	score := &models.ScoreEntry{
		StudentID: req.StudentID, SubjectID: req.SubjectID, TermID: req.TermID,
		CA1: req.CA1, CA2: req.CA2, Exam: req.Exam, Remark: req.Remark,
	}
	result, err := rc.resultService.SubmitScores(c.Request.Context(), score)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Scores submitted successfully", "result": result})
}

func (rc *ResultController) GetStudentResults(c *gin.Context) {
	studentID := c.Param("student_id")
	results, err := rc.resultService.GetStudentResults(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"results": results})
}

func (rc *ResultController) PublishResults(c *gin.Context) {
	var req struct {
		ClassID   string `json:"class_id"`
		SessionID string `json:"session_id"`
		TermID    string `json:"term_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := rc.resultService.PublishResults(c.Request.Context(), req.ClassID, req.SessionID, req.TermID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Results published successfully"})
}

func (rc *ResultController) LockResults(c *gin.Context) {
	var req struct {
		ClassID   string `json:"class_id"`
		SessionID string `json:"session_id"`
		TermID    string `json:"term_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := rc.resultService.LockResults(c.Request.Context(), req.ClassID, req.SessionID, req.TermID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Results locked successfully"})
}

func (rc *ResultController) GenerateReportCard(c *gin.Context) {
	studentID := c.Param("student_id")
	sessionID := c.Query("session_id")
	termID := c.Query("term_id")
	reportCard, err := rc.resultService.GenerateReportCard(c.Request.Context(), studentID, sessionID, termID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, reportCard)
}
