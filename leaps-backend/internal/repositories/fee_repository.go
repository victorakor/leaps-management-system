package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type FeeRepository struct {
	db *sql.DB
}

func NewFeeRepository(db *sql.DB) *FeeRepository {
	return &FeeRepository{db: db}
}

func (fr *FeeRepository) Create(fee *models.Fee) error {
	query := `
		INSERT INTO fees (id, class_id, session_id, amount, created_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err := fr.db.Exec(query, fee.ID, fee.ClassID, fee.SessionID, fee.Amount)
	return err
}

func (fr *FeeRepository) GetByID(id string) (*models.Fee, error) {
	query := `
		SELECT id, class_id, session_id, amount, created_at
		FROM fees
		WHERE id = $1
	`

	var fee models.Fee
	err := fr.db.QueryRow(query, id).Scan(
		&fee.ID, &fee.ClassID, &fee.SessionID, &fee.Amount, &fee.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("fee not found")
		}
		return nil, err
	}

	return &fee, nil
}

func (fr *FeeRepository) GetByClass(classID string, sessionID string) (*models.Fee, error) {
	query := `
		SELECT id, class_id, session_id, amount, created_at
		FROM fees
		WHERE class_id = $1 AND session_id = $2
	`

	var fee models.Fee
	err := fr.db.QueryRow(query, classID, sessionID).Scan(
		&fee.ID, &fee.ClassID, &fee.SessionID, &fee.Amount, &fee.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("fee not found")
		}
		return nil, err
	}

	return &fee, nil
}

func (fr *FeeRepository) GetTotalByStudent(studentID string, sessionID string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(f.amount), 0)
		FROM fees f
		JOIN students s ON f.class_id = s.current_class_id
		WHERE s.id = $1 AND f.session_id = $2
	`

	var total float64
	err := fr.db.QueryRow(query, studentID, sessionID).Scan(&total)
	return total, err
}

func (fr *FeeRepository) GetTotalBySession(sessionID string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(f.amount * (SELECT COUNT(*) FROM students WHERE current_class_id = f.class_id)), 0)
		FROM fees f
		WHERE f.session_id = $1
	`

	var total float64
	err := fr.db.QueryRow(query, sessionID).Scan(&total)
	return total, err
}

func (fr *FeeRepository) Update(fee *models.Fee) error {
	query := `
		UPDATE fees
		SET amount = $1
		WHERE id = $2
	`

	result, err := fr.db.Exec(query, fee.Amount, fee.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("fee not found")
	}

	return nil
}
