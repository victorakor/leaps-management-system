package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"leaps/internal/models"
)

type AdmissionService struct {
	db *sql.DB
}

func NewAdmissionService(db *sql.DB) *AdmissionService {
	return &AdmissionService{db: db}
}

func (as *AdmissionService) SubmitApplication(ctx context.Context, app *models.Application) (*models.Application, error) {
	if app.Email == "" {
		return nil, errors.New("email is required")
	}
	app.Status = "pending"
	query := `INSERT INTO applications (first_name, last_name, email, phone, class_id, status, applied_at)
		VALUES ($1,$2,$3,$4,$5,$6,NOW()) RETURNING id, applied_at`
	err := as.db.QueryRowContext(ctx, query,
		app.FirstName, app.LastName, app.Email, app.Phone, app.ClassID, app.Status,
	).Scan(&app.ID, &app.AppliedAt)
	return app, err
}

func (as *AdmissionService) ScheduleInterview(ctx context.Context, appt *models.Appointment) (*models.Appointment, error) {
	if appt.ApplicationID == "" {
		return nil, errors.New("application_id is required")
	}
	query := `INSERT INTO appointments (application_id, scheduled_at, notes)
		VALUES ($1,$2,$3) RETURNING id`
	err := as.db.QueryRowContext(ctx, query, appt.ApplicationID, appt.ScheduledAt, appt.Notes).Scan(&appt.ID)
	return appt, err
}

func (as *AdmissionService) ApproveApplication(ctx context.Context, id string) error {
	result, err := as.db.ExecContext(ctx, `UPDATE applications SET status='approved' WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("application not found")
	}
	return nil
}

func (as *AdmissionService) RejectApplication(ctx context.Context, id, reason string) error {
	result, err := as.db.ExecContext(ctx, `UPDATE applications SET status='rejected' WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("application not found")
	}
	return nil
}

func (as *AdmissionService) GetApplications(ctx context.Context, status string, page, limit int) ([]*models.Application, error) {
	offset := (page - 1) * limit
	query := `SELECT id, first_name, last_name, email, phone, class_id, status, applied_at
		FROM applications WHERE ($1='' OR status=$1) ORDER BY applied_at DESC LIMIT $2 OFFSET $3`
	rows, err := as.db.QueryContext(ctx, query, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var apps []*models.Application
	for rows.Next() {
		var a models.Application
		rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Email, &a.Phone, &a.ClassID, &a.Status, &a.AppliedAt)
		apps = append(apps, &a)
	}
	return apps, rows.Err()
}

func (as *AdmissionService) GenerateAdmissionLetter(ctx context.Context, id string) (string, error) {
	var app models.Application
	err := as.db.QueryRowContext(ctx, `SELECT id, first_name, last_name, status FROM applications WHERE id=$1`, id).
		Scan(&app.ID, &app.FirstName, &app.LastName, &app.Status)
	if err == sql.ErrNoRows {
		return "", errors.New("application not found")
	}
	if err != nil {
		return "", err
	}
	letter := fmt.Sprintf("Dear %s %s,\n\nYour application status is: %s.\n\nRegards,\nLEAPS Admissions", app.FirstName, app.LastName, app.Status)
	return letter, nil
}
