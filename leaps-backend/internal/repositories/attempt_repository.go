package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type AttemptRepository struct {
	db *sql.DB
}

func NewAttemptRepository(db *sql.DB) *AttemptRepository {
	return &AttemptRepository{db: db}
}

func (ar *AttemptRepository) Create(attempt *models.QuizAttempt) error {
	query := `
		INSERT INTO quiz_attempts (id, quiz_id, student_id, score, started_at, submitted_at, status)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := ar.db.Exec(query, attempt.ID, attempt.QuizID, attempt.StudentID, attempt.Score, attempt.StartedAt, attempt.SubmittedAt, attempt.Status)
	return err
}

func (ar *AttemptRepository) GetByID(id string) (*models.QuizAttempt, error) {
	query := `
		SELECT id, quiz_id, student_id, score, started_at, submitted_at, status
		FROM quiz_attempts
		WHERE id = $1
	`

	var attempt models.QuizAttempt
	err := ar.db.QueryRow(query, id).Scan(
		&attempt.ID, &attempt.QuizID, &attempt.StudentID, &attempt.Score, &attempt.StartedAt, &attempt.SubmittedAt, &attempt.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("attempt not found")
		}
		return nil, err
	}

	return &attempt, nil
}

func (ar *AttemptRepository) GetByQuizID(quizID string) ([]models.QuizAttempt, error) {
	query := `
		SELECT id, quiz_id, student_id, score, started_at, submitted_at, status
		FROM quiz_attempts
		WHERE quiz_id = $1
		ORDER BY started_at DESC
	`

	rows, err := ar.db.Query(query, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var attempts []models.QuizAttempt
	for rows.Next() {
		var attempt models.QuizAttempt
		err := rows.Scan(
			&attempt.ID, &attempt.QuizID, &attempt.StudentID, &attempt.Score, &attempt.StartedAt, &attempt.SubmittedAt, &attempt.Status,
		)
		if err != nil {
			return nil, err
		}
		attempts = append(attempts, attempt)
	}

	return attempts, rows.Err()
}

func (ar *AttemptRepository) GetByQuizAndStudent(quizID string, studentID string) (*models.QuizAttempt, error) {
	query := `
		SELECT id, quiz_id, student_id, score, started_at, submitted_at, status
		FROM quiz_attempts
		WHERE quiz_id = $1 AND student_id = $2
		ORDER BY started_at DESC
		LIMIT 1
	`

	var attempt models.QuizAttempt
	err := ar.db.QueryRow(query, quizID, studentID).Scan(
		&attempt.ID, &attempt.QuizID, &attempt.StudentID, &attempt.Score, &attempt.StartedAt, &attempt.SubmittedAt, &attempt.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("attempt not found")
		}
		return nil, err
	}

	return &attempt, nil
}

func (ar *AttemptRepository) Update(attempt *models.QuizAttempt) error {
	query := `
		UPDATE quiz_attempts
		SET score = $1, submitted_at = $2, status = $3
		WHERE id = $4
	`

	result, err := ar.db.Exec(query, attempt.Score, attempt.SubmittedAt, attempt.Status, attempt.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("attempt not found")
	}

	return nil
}
