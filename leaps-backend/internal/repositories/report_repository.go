package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (rr *ReportRepository) Create(report *models.ReportCard) error {
	query := `
		INSERT INTO report_cards (id, student_id, session_id, term_id, file_url, result_id, is_locked, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := rr.db.Exec(query, report.ID, report.StudentID, report.SessionID, report.TermID, report.FileURL, report.ResultID, report.IsLocked)
	return err
}

func (rr *ReportRepository) GetByID(id string) (*models.ReportCard, error) {
	query := `
		SELECT id, student_id, session_id, term_id, file_url, result_id, is_locked, created_at
		FROM report_cards
		WHERE id = $1
	`

	var report models.ReportCard
	err := rr.db.QueryRow(query, id).Scan(
		&report.ID, &report.StudentID, &report.SessionID, &report.TermID, &report.FileURL, &report.ResultID, &report.IsLocked, &report.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("report card not found")
		}
		return nil, err
	}

	return &report, nil
}

func (rr *ReportRepository) GetByStudent(studentID string, sessionID string, termID string) (*models.ReportCard, error) {
	query := `
		SELECT id, student_id, session_id, term_id, file_url, result_id, is_locked, created_at
		FROM report_cards
		WHERE student_id = $1 AND session_id = $2 AND term_id = $3
	`

	var report models.ReportCard
	err := rr.db.QueryRow(query, studentID, sessionID, termID).Scan(
		&report.ID, &report.StudentID, &report.SessionID, &report.TermID, &report.FileURL, &report.ResultID, &report.IsLocked, &report.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("report card not found")
		}
		return nil, err
	}

	return &report, nil
}

func (rr *ReportRepository) Update(report *models.ReportCard) error {
	query := `
		UPDATE report_cards
		SET file_url = $1, is_locked = $2
		WHERE id = $3
	`

	result, err := rr.db.Exec(query, report.FileURL, report.IsLocked, report.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("report card not found")
	}

	return nil
}
