package main

import (
	"context"
	"testing"

	"leaps/internal/models"
	"leaps/internal/services"
)

// MockUserRepository is a mock implementation
type MockUserRepository struct {
	users map[int]*models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]*models.User),
	}
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.ID = len(m.users) + 1
	m.users[user.ID] = user
	return user, nil
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	if user, exists := m.users[id]; exists {
		return user, nil
	}
	return nil, nil
}

func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if _, exists := m.users[user.ID]; exists {
		m.users[user.ID] = user
		return user, nil
	}
	return nil, nil
}

func (m *MockUserRepository) DeactivateUser(ctx context.Context, id int) error {
	if user, exists := m.users[id]; exists {
		user.Status = "inactive"
		return nil
	}
	return nil
}

func (m *MockUserRepository) ListUsers(ctx context.Context, schoolID, page, limit int) ([]*models.User, error) {
	var users []*models.User
	for _, user := range m.users {
		if user.SchoolID == schoolID {
			users = append(users, user)
		}
	}
	return users, nil
}

func (m *MockUserRepository) UserExists(ctx context.Context, email string) (bool, error) {
	for _, user := range m.users {
		if user.Email == email {
			return true, nil
		}
	}
	return false, nil
}

// Test Cases
func TestCreateUser(t *testing.T) {
	mockRepo := NewMockUserRepository()
	service := services.NewUserService(mockRepo)

	user := &models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "08012345678",
		Role:      "teacher",
		SchoolID:  1,
		Status:    "active",
	}

	createdUser, err := service.CreateUser(context.Background(), user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	if createdUser.ID == 0 {
		t.Error("User ID should not be 0")
	}

	if createdUser.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, createdUser.Email)
	}
}

func TestGetUserByID(t *testing.T) {
	mockRepo := NewMockUserRepository()
	service := services.NewUserService(mockRepo)

	user := &models.User{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane@example.com",
		Phone:     "08087654321",
		Role:      "admin",
		SchoolID:  1,
		Status:    "active",
	}

	createdUser, _ := service.CreateUser(context.Background(), user)

	retrievedUser, err := service.GetUserByID(context.Background(), createdUser.ID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}

	if retrievedUser.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, retrievedUser.Email)
	}
}

func TestUserExists(t *testing.T) {
	mockRepo := NewMockUserRepository()
	service := services.NewUserService(mockRepo)

	user := &models.User{
		FirstName: "Alice",
		LastName:  "Brown",
		Email:     "alice@example.com",
		Phone:     "08022222222",
		Role:      "parent",
		SchoolID:  1,
		Status:    "active",
	}

	service.CreateUser(context.Background(), user)

	exists, err := service.UserExists(context.Background(), user.Email)
	if err != nil {
		t.Fatalf("UserExists failed: %v", err)
	}

	if !exists {
		t.Error("User should exist")
	}
}
