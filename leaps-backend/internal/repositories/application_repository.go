package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type ApplicationRepository struct {
	db *sql.DB
}

func NewApplicationRepository(db *sql.DB) *ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (ar *ApplicationRepository) Create(app *models.Application) error {
	query := `
		INSERT INTO applications (id, application_number, full_name, desired_class, status, created_at)
		VALUES ($1, $2, $3, $4, $5, NOW())
	`

	_, err := ar.db.Exec(query, app.ID, app.ApplicationNumber, app.FullName, app.DesiredClass, app.Status)
	return err
}

func (ar *ApplicationRepository) GetByID(id string) (*models.Application, error) {
	query := `
		SELECT id, application_number, full_name, desired_class, status, created_at
		FROM applications
		WHERE id = $1
	`

	var app models.Application
	err := ar.db.QueryRow(query, id).Scan(
		&app.ID, &app.ApplicationNumber, &app.FullName, &app.DesiredClass, &app.Status, &app.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("application not found")
		}
		return nil, err
	}

	return &app, nil
}

func (ar *ApplicationRepository) Update(app *models.Application) error {
	query := `
		UPDATE applications
		SET status = $1
		WHERE id = $2
	`

	result, err := ar.db.Exec(query, app.Status, app.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("application not found")
	}

	return nil
}

func (ar *ApplicationRepository) List(status string, limit int, offset int) ([]models.Application, error) {
	var query string
	var args []interface{}

	if status == "" {
		query = `
			SELECT id, application_number, full_name, desired_class, status, created_at
			FROM applications
			ORDER BY created_at DESC
			LIMIT $1 OFFSET $2
		`
		args = []interface{}{limit, offset}
	} else {
		query = `
			SELECT id, application_number, full_name, desired_class, status, created_at
			FROM applications
			WHERE status = $1
			ORDER BY created_at DESC
			LIMIT $2 OFFSET $3
		`
		args = []interface{}{status, limit, offset}
	}

	rows, err := ar.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apps []models.Application
	for rows.Next() {
		var app models.Application
		err := rows.Scan(
			&app.ID, &app.ApplicationNumber, &app.FullName, &app.DesiredClass, &app.Status, &app.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		apps = append(apps, app)
	}

	return apps, rows.Err()
}
