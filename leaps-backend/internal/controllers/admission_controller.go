package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"leaps/internal/models"
	"leaps/internal/services"
)

type AdmissionController struct {
	admissionService *services.AdmissionService
}

func NewAdmissionController(admissionService *services.AdmissionService) *AdmissionController {
	return &AdmissionController{admissionService: admissionService}
}

func (ac *AdmissionController) SubmitApplication(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Phone     string `json:"phone"`
		ClassID   string `json:"class_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	app := &models.Application{
		FirstName: req.FirstName, LastName: req.LastName,
		Email: req.Email, Phone: req.Phone, ClassID: req.ClassID,
	}
	created, err := ac.admissionService.SubmitApplication(c.Request.Context(), app)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Application submitted", "application": created})
}

func (ac *AdmissionController) ScheduleInterview(c *gin.Context) {
	var req struct {
		ApplicationID string    `json:"application_id" binding:"required"`
		ScheduledAt   time.Time `json:"scheduled_at" binding:"required"`
		Notes         string    `json:"notes"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	appt := &models.Appointment{ApplicationID: req.ApplicationID, ScheduledAt: req.ScheduledAt, Notes: req.Notes}
	created, err := ac.admissionService.ScheduleInterview(c.Request.Context(), appt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Interview scheduled", "appointment": created})
}

func (ac *AdmissionController) ApproveApplication(c *gin.Context) {
	id := c.Param("id")
	if err := ac.admissionService.ApproveApplication(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application approved"})
}

func (ac *AdmissionController) RejectApplication(c *gin.Context) {
	id := c.Param("id")
	var req struct{ Reason string `json:"reason"` }
	c.ShouldBindJSON(&req)
	if err := ac.admissionService.RejectApplication(c.Request.Context(), id, req.Reason); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Application rejected"})
}

func (ac *AdmissionController) GetApplications(c *gin.Context) {
	status := c.Query("status")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	apps, err := ac.admissionService.GetApplications(c.Request.Context(), status, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"applications": apps})
}

func (ac *AdmissionController) GenerateAdmissionLetter(c *gin.Context) {
	id := c.Param("id")
	letter, err := ac.admissionService.GenerateAdmissionLetter(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"letter": letter})
}
