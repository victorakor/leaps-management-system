package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"leaps/internal/models"
	"leaps/internal/services"
)

type FinanceController struct {
	financeService *services.FinanceService
}

func NewFinanceController(financeService *services.FinanceService) *FinanceController {
	return &FinanceController{financeService: financeService}
}

func (fc *FinanceController) CreateFeeStructure(c *gin.Context) {
	var fee models.Fee
	if err := c.ShouldBindJSON(&fee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := fc.financeService.CreateFeeStructure(c.Request.Context(), &fee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"fee": created})
}

func (fc *FinanceController) RecordPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	receipt, err := fc.financeService.RecordPayment(c.Request.Context(), &payment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"receipt": receipt})
}

func (fc *FinanceController) GetStudentBalance(c *gin.Context) {
	studentID := c.Param("student_id")
	balance, err := fc.financeService.GetStudentBalance(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (fc *FinanceController) GetPaymentHistory(c *gin.Context) {
	studentID := c.Param("student_id")
	payments, err := fc.financeService.GetPaymentHistory(c.Request.Context(), studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

func (fc *FinanceController) GetDebtors(c *gin.Context) {
	schoolID := c.Query("school_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	debtors, err := fc.financeService.GetDebtors(c.Request.Context(), schoolID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"debtors": debtors})
}

func (fc *FinanceController) GetFinancialSummary(c *gin.Context) {
	schoolID := c.Query("school_id")
	sessionID := c.Query("session_id")
	summary, err := fc.financeService.GetFinancialSummary(c.Request.Context(), schoolID, sessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, summary)
}
