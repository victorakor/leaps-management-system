package repositories

import (
	"context"
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: db}
}

func (ur *UserRepositoryDB) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
		INSERT INTO users (first_name, last_name, email, phone, role, school_id, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
		RETURNING id, created_at
	`
	err := ur.db.QueryRowContext(ctx, query,
		user.FirstName, user.LastName, user.Email, user.Phone,
		user.Role, user.SchoolID, user.Status,
	).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepositoryDB) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, role, school_id, status, created_at
		FROM users WHERE id = $1 AND status != 'deleted'
	`
	var user models.User
	err := ur.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone,
		&user.Role, &user.SchoolID, &user.Status, &user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryDB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, phone, role, school_id, status, created_at
		FROM users WHERE email = $1 AND status != 'deleted'
	`
	var user models.User
	err := ur.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone,
		&user.Role, &user.SchoolID, &user.Status, &user.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepositoryDB) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := `
		UPDATE users
		SET first_name = $1, last_name = $2, phone = $3, role = $4, updated_at = NOW()
		WHERE id = $5
		RETURNING id, first_name, last_name, email, phone, role, school_id, status, created_at
	`
	err := ur.db.QueryRowContext(ctx, query,
		user.FirstName, user.LastName, user.Phone, user.Role, user.ID,
	).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Phone,
		&user.Role, &user.SchoolID, &user.Status, &user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepositoryDB) DeactivateUser(ctx context.Context, id int) error {
	query := `UPDATE users SET status = 'inactive', updated_at = NOW() WHERE id = $1`
	result, err := ur.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (ur *UserRepositoryDB) ListUsers(ctx context.Context, schoolID, page, limit int) ([]*models.User, error) {
	offset := (page - 1) * limit
	query := `
		SELECT id, first_name, last_name, email, phone, role, school_id, status, created_at
		FROM users WHERE school_id = $1 AND status != 'deleted'
		ORDER BY created_at DESC LIMIT $2 OFFSET $3
	`
	rows, err := ur.db.QueryContext(ctx, query, schoolID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Phone,
			&u.Role, &u.SchoolID, &u.Status, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, rows.Err()
}

func (ur *UserRepositoryDB) UserExists(ctx context.Context, email string) (bool, error) {
	var count int
	err := ur.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM users WHERE email = $1`, email).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
