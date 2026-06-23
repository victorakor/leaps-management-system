package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"leaps/internal/models"
	"leaps/internal/services"
)

// StudentController handles student-related HTTP requests
type StudentController struct {
	studentService *services.StudentService
}

// NewStudentController creates a new student controller
func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{
		studentService: studentService,
	}
}

// CreateStudent creates a new student
// POST /api/students
func (sc *StudentController) CreateStudent(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Phone     string `json:"phone"`
		ClassID   int    `json:"class_id" binding:"required"`
		SchoolID  int    `json:"school_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := &models.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		ClassID:   strconv.Itoa(req.ClassID),
		SchoolID:  req.SchoolID,
		Status:    "active",
	}

	createdStudent, err := sc.studentService.CreateStudent(c.Request.Context(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created successfully",
		"student": createdStudent,
	})
}

// GetStudent retrieves a student by ID
// GET /api/students/:id
func (sc *StudentController) GetStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	student, err := sc.studentService.GetStudentByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// GetStudentByAdmissionNumber retrieves a student by admission number
// GET /api/students/admission/:admission_number
func (sc *StudentController) GetStudentByAdmissionNumber(c *gin.Context) {
	admissionNumber := c.Param("admission_number")

	student, err := sc.studentService.GetStudentByAdmissionNumber(c.Request.Context(), admissionNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	c.JSON(http.StatusOK, student)
}

// UpdateStudent updates a student
// PUT /api/students/:id
func (sc *StudentController) UpdateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := &models.Student{
		ID:        strconv.Itoa(id),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
	}

	updatedStudent, err := sc.studentService.UpdateStudent(c.Request.Context(), student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated successfully",
		"student": updatedStudent,
	})
}

// TransferStudent transfers a student to a new class
// POST /api/students/:id/transfer
func (sc *StudentController) TransferStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	var req struct {
		NewClassID int    `json:"new_class_id" binding:"required"`
		Reason     string `json:"reason"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = sc.studentService.TransferStudent(c.Request.Context(), id, req.NewClassID, req.Reason)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student transferred successfully"})
}

// DeactivateStudent deactivates a student
// DELETE /api/students/:id
func (sc *StudentController) DeactivateStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	err = sc.studentService.DeactivateStudent(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student deactivated successfully"})
}
