package controllers

import (
	"database/sql"
	"log"
	"net/http"

	"leaps/internal/auth"
	"leaps/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	ID           string `json:"id"`
	SchoolID     string `json:"school_id"`
	FullName     string `json:"full_name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	PasswordHash string `json:"-"`
	RoleID       string `json:"role_id"`
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
	var phone sql.NullString
	var schoolID sql.NullString
	query := `
		SELECT u.id, u.school_id, u.full_name, u.email, u.phone, u.password_hash, u.role_id, u.is_active, r.name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.email = $1
	`
	err := ac.db.QueryRow(query, req.Email).Scan(
		&user.ID, &schoolID, &user.FullName, &user.Email, &phone,
		&user.PasswordHash, &user.RoleID, &user.IsActive, &user.Role,
	)
	if err != nil {
		log.Printf("LOGIN ERROR for email=%s: %v", req.Email, err)
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", err.Error())
		return
	}

	if phone.Valid {
		user.Phone = phone.String
	}
	if schoolID.Valid {
		user.SchoolID = schoolID.String
	}

	if !user.IsActive {
		response.Error(c, http.StatusUnauthorized, "Account inactive", "user account is not active")
		return
	}

	if !auth.VerifyPassword(user.PasswordHash, req.Password) {
		log.Printf("PASSWORD MISMATCH for email=%s", req.Email)
		response.Error(c, http.StatusUnauthorized, "Invalid credentials", "password mismatch")
		return
	}

	token, err := auth.GenerateTokenRaw(user.ID, user.Email, user.Role)
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
		FullName string `json:"full_name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Phone    string `json:"phone"`
		Password string `json:"password" binding:"required"`
		SchoolID string `json:"school_id"`
		RoleID   string `json:"role_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request", err.Error())
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Password hashing failed", err.Error())
		return
	}

	userID := uuid.New().String()
	query := `
		INSERT INTO users (id, school_id, full_name, email, phone, password_hash, role_id, is_active, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, true, NOW())
	`
	_, err = ac.db.Exec(query, userID, req.SchoolID, req.FullName, req.Email, req.Phone, hashedPassword, req.RoleID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Registration failed", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User registered successfully", gin.H{"id": userID, "email": req.Email})
}

func (ac *AuthController) GetCurrentUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		response.Error(c, http.StatusUnauthorized, "User not found", "")
		return
	}

	var user authUser
	var phone sql.NullString
	var schoolID sql.NullString
	query := `
		SELECT u.id, u.school_id, u.full_name, u.email, u.phone, u.role_id, u.is_active, r.name
		FROM users u
		LEFT JOIN roles r ON u.role_id = r.id
		WHERE u.id = $1
	`
	err := ac.db.QueryRow(query, userID).Scan(
		&user.ID, &schoolID, &user.FullName, &user.Email, &phone,
		&user.RoleID, &user.IsActive, &user.Role,
	)
	if err != nil {
		log.Printf("GET CURRENT USER ERROR for id=%v: %v", userID, err)
		response.Error(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	if phone.Valid {
		user.Phone = phone.String
	}
	if schoolID.Valid {
		user.SchoolID = schoolID.String
	}

	response.Success(c, http.StatusOK, "User retrieved", user)
}
