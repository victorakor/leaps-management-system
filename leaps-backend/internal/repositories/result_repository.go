package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type ResultRepository struct {
	db *sql.DB
}

func NewResultRepository(db *sql.DB) *ResultRepository {
	return &ResultRepository{db: db}
}

func (rr *ResultRepository) Create(result *models.ComputedResult) error {
	query := `
		INSERT INTO computed_results (id, student_id, subject_id, total_score, grade, subject_position, class_position, session_id, term_id, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, NOW())
	`

	_, err := rr.db.Exec(query, result.ID, result.StudentID, result.SubjectID, result.TotalScore, result.Grade, result.SubjectPosition, result.ClassPosition, result.SessionID, result.TermID)
	return err
}

func (rr *ResultRepository) GetByStudent(studentID string, sessionID string, termID string) ([]models.ComputedResult, error) {
	query := `
		SELECT id, student_id, subject_id, total_score, grade, subject_position, class_position, session_id, term_id
		FROM computed_results
		WHERE student_id = $1 AND session_id = $2 AND term_id = $3
		ORDER BY subject_id ASC
	`

	rows, err := rr.db.Query(query, studentID, sessionID, termID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.ComputedResult
	for rows.Next() {
		var result models.ComputedResult
		err := rows.Scan(
			&result.ID, &result.StudentID, &result.SubjectID, &result.TotalScore, &result.Grade,
			&result.SubjectPosition, &result.ClassPosition, &result.SessionID, &result.TermID,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, rows.Err()
}

func (rr *ResultRepository) UpdateStatus(sessionID string, termID string, status string) error {
	query := `
		UPDATE computed_results
		SET grade = CASE WHEN $3 = 'LOCKED' THEN grade ELSE grade END
		WHERE session_id = $1 AND term_id = $2
	`

	_, err := rr.db.Exec(query, sessionID, termID, status)
	return err
}

func (rr *ResultRepository) GetByID(id string) (*models.ComputedResult, error) {
	query := `
		SELECT id, student_id, subject_id, total_score, grade, subject_position, class_position, session_id, term_id
		FROM computed_results
		WHERE id = $1
	`

	var result models.ComputedResult
	err := rr.db.QueryRow(query, id).Scan(
		&result.ID, &result.StudentID, &result.SubjectID, &result.TotalScore, &result.Grade,
		&result.SubjectPosition, &result.ClassPosition, &result.SessionID, &result.TermID,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("result not found")
		}
		return nil, err
	}

	return &result, nil
}
