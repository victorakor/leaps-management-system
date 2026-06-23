package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type QuizRepository struct {
	db *sql.DB
}

func NewQuizRepository(db *sql.DB) *QuizRepository {
	return &QuizRepository{db: db}
}

func (qr *QuizRepository) Create(quiz *models.Quiz) error {
	query := `
		INSERT INTO quizzes (id, title, class_id, created_by, duration_minutes, start_time, end_time, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
	`

	_, err := qr.db.Exec(query, quiz.ID, quiz.Title, quiz.ClassID, quiz.CreatedBy, quiz.DurationMinutes, quiz.StartTime, quiz.EndTime)
	return err
}

func (qr *QuizRepository) GetByID(id string) (*models.Quiz, error) {
	query := `
		SELECT id, title, class_id, created_by, duration_minutes, start_time, end_time
		FROM quizzes
		WHERE id = $1
	`

	var quiz models.Quiz
	err := qr.db.QueryRow(query, id).Scan(
		&quiz.ID, &quiz.Title, &quiz.ClassID, &quiz.CreatedBy, &quiz.DurationMinutes, &quiz.StartTime, &quiz.EndTime,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("quiz not found")
		}
		return nil, err
	}

	return &quiz, nil
}

func (qr *QuizRepository) GetByClass(classID string) ([]models.Quiz, error) {
	query := `
		SELECT id, title, class_id, created_by, duration_minutes, start_time, end_time
		FROM quizzes
		WHERE class_id = $1
		ORDER BY start_time DESC
	`

	rows, err := qr.db.Query(query, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []models.Quiz
	for rows.Next() {
		var quiz models.Quiz
		err := rows.Scan(
			&quiz.ID, &quiz.Title, &quiz.ClassID, &quiz.CreatedBy, &quiz.DurationMinutes, &quiz.StartTime, &quiz.EndTime,
		)
		if err != nil {
			return nil, err
		}
		quizzes = append(quizzes, quiz)
	}

	return quizzes, rows.Err()
}

func (qr *QuizRepository) Update(quiz *models.Quiz) error {
	query := `
		UPDATE quizzes
		SET title = $1, duration_minutes = $2, start_time = $3, end_time = $4
		WHERE id = $5
	`

	result, err := qr.db.Exec(query, quiz.Title, quiz.DurationMinutes, quiz.StartTime, quiz.EndTime, quiz.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("quiz not found")
	}

	return nil
}
