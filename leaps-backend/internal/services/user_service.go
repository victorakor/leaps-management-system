package services

import (
	"context"
	"errors"

	"leaps/internal/models"
)

// UserRepository defines the interface the UserService depends on
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeactivateUser(ctx context.Context, id int) error
	ListUsers(ctx context.Context, schoolID, page, limit int) ([]*models.User, error)
	UserExists(ctx context.Context, email string) (bool, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.Email == "" {
		return nil, errors.New("email is required")
	}
	if user.FirstName == "" {
		return nil, errors.New("first name is required")
	}
	user.Status = "active"
	return us.repo.CreateUser(ctx, user)
}

func (us *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("valid user id is required")
	}
	return us.repo.GetUserByID(ctx, id)
}

func (us *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return us.repo.GetUserByEmail(ctx, email)
}

func (us *UserService) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user.ID == 0 {
		return nil, errors.New("user id is required")
	}
	return us.repo.UpdateUser(ctx, user)
}

func (us *UserService) DeactivateUser(ctx context.Context, id int) error {
	if id <= 0 {
		return errors.New("valid user id is required")
	}
	return us.repo.DeactivateUser(ctx, id)
}

func (us *UserService) ListUsers(ctx context.Context, schoolID, page, limit int) ([]*models.User, error) {
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	return us.repo.ListUsers(ctx, schoolID, page, limit)
}

func (us *UserService) UserExists(ctx context.Context, email string) (bool, error) {
	if email == "" {
		return false, errors.New("email is required")
	}
	return us.repo.UserExists(ctx, email)
}
