package repositories

import (
	"database/sql"
	"errors"
	"time"

	"leaps/internal/models"
)

type AppointmentRepository struct {
	db *sql.DB
}

func NewAppointmentRepository(db *sql.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (ar *AppointmentRepository) Create(apt *models.AdmissionAppointment) error {
	query := `
		INSERT INTO admission_appointments (id, application_id, date, time, status)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := ar.db.Exec(query, apt.ID, apt.ApplicationID, apt.Date, apt.Time, apt.Status)
	return err
}

func (ar *AppointmentRepository) GetByApplicationID(appID string) (*models.AdmissionAppointment, error) {
	query := `
		SELECT id, application_id, date, time, status
		FROM admission_appointments
		WHERE application_id = $1
	`

	var apt models.AdmissionAppointment
	err := ar.db.QueryRow(query, appID).Scan(
		&apt.ID, &apt.ApplicationID, &apt.Date, &apt.Time, &apt.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("appointment not found")
		}
		return nil, err
	}

	return &apt, nil
}

func (ar *AppointmentRepository) IsSlotAvailable(date time.Time, timeSlot string) (bool, error) {
	query := `
		SELECT COUNT(*) FROM admission_appointments
		WHERE date = $1 AND time = $2 AND status = 'scheduled'
	`

	var count int
	err := ar.db.QueryRow(query, date, timeSlot).Scan(&count)
	if err != nil {
		return false, err
	}

	// Max 5 slots per time
	return count < 5, nil
}

func (ar *AppointmentRepository) Update(apt *models.AdmissionAppointment) error {
	query := `
		UPDATE admission_appointments
		SET status = $1
		WHERE id = $2
	`

	result, err := ar.db.Exec(query, apt.Status, apt.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("appointment not found")
	}

	return nil
}
