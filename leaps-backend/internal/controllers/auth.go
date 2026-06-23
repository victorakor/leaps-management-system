package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"leaps/internal/auth"
	"leaps/pkg/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	db *sql.DB
}

func NewAuthController(db *sql.DB) *AuthController {
	return &AuthController{db: db}
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type authUser struct {
	ID           int    `json:"id"`
	SchoolID     int    `json:"school_id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"-"`
	Role         string `json:"role"`
	IsActive     bool   `json:"is_active"`
}

func (ac *AuthController) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	var user authUser
	var firstName, lastName string
	var status string
	query := `
		SELECT id, school_id, first_name, last_name, email, phone, password_hash, role, status
		FROM users
		WHERE email = $1
		  AND deleted_at IS NULL
	`
	err := ac.db.QueryRow(query, req.Email).Scan(
		&user.ID, &user.SchoolID, &firstName, &lastName,
		&user.Email, &user.Phone, &user.PasswordHash, &user.Role, &status,
	)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", "user not found")
		return
	}

	user.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	user.IsActive = status == "active"

	if !user.IsActive {
		response.Error(c, http.StatusUnauthorized, "Account inactive", "user account is not active")
		return
	}

	if !auth.VerifyPassword(user.PasswordHash, req.Password) {
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", "password mismatch")
		return
	}

	token, err := auth.GenerateTokenRaw(fmt.Sprintf("%d", user.ID), user.Email, user.Role)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Token generation failed", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  user,
	})
}

func (ac *AuthController) Register(c *gin.Context) {
	var req struct {
		FirstName string `json:"first_name" binding:"required"`
		LastName  string `json:"last_name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Phone     string `json:"phone"`
		Password  string `json:"password" binding:"required"`
		SchoolID  int    `json:"school_id" binding:"required"`
		Role      string `json:"role"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	if req.Role == "" {
		req.Role = "staff"
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Password hashing failed", err.Error())
		return
	}

	query := `
		INSERT INTO users (first_name, last_name, email, phone, password_hash, role, school_id, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, 'active', NOW(), NOW())
		RETURNING id
	`
	var userID int
	err = ac.db.QueryRow(query,
		req.FirstName, req.LastName, req.Email, req.Phone,
		hashedPassword, req.Role, req.SchoolID,
	).Scan(&userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", gin.H{
		"id":    userID,
		"email": req.Email,
	})
}

func (ac *AuthController) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User not found", "")
		return
	}

	var user authUser
	var firstName, lastName, status string
	query := `
		SELECT id, school_id, first_name, last_name, email, phone, role, status
		FROM users
		WHERE id = $1
		  AND deleted_at IS NULL
	`
	err := ac.db.QueryRow(query, userID).Scan(
		&user.ID, &user.SchoolID, &firstName, &lastName,
		&user.Email, &user.Phone, &user.Role, &status,
	)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	user.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	user.IsActive = status == "active"

	response.Success(c, http.StatusOK, "User retrieved", user)
}
