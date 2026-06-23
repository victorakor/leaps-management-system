package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (qr *QuestionRepository) Create(question *models.QuizQuestion) error {
	query := `
		INSERT INTO quiz_questions (id, quiz_id, question, option_a, option_b, option_c, option_d, correct_answer)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := qr.db.Exec(query, question.ID, question.QuizID, question.Question, question.OptionA, question.OptionB, question.OptionC, question.OptionD, question.CorrectAnswer)
	return err
}

func (qr *QuestionRepository) GetByQuizID(quizID string) ([]models.QuizQuestion, error) {
	query := `
		SELECT id, quiz_id, question, option_a, option_b, option_c, option_d, correct_answer
		FROM quiz_questions
		WHERE quiz_id = $1
		ORDER BY id ASC
	`

	rows, err := qr.db.Query(query, quizID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.QuizQuestion
	for rows.Next() {
		var q models.QuizQuestion
		err := rows.Scan(
			&q.ID, &q.QuizID, &q.Question, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectAnswer,
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}

	return questions, rows.Err()
}

func (qr *QuestionRepository) GetByID(id string) (*models.QuizQuestion, error) {
	query := `
		SELECT id, quiz_id, question, option_a, option_b, option_c, option_d, correct_answer
		FROM quiz_questions
		WHERE id = $1
	`

	var q models.QuizQuestion
	err := qr.db.QueryRow(query, id).Scan(
		&q.ID, &q.QuizID, &q.Question, &q.OptionA, &q.OptionB, &q.OptionC, &q.OptionD, &q.CorrectAnswer,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("question not found")
		}
		return nil, err
	}

	return &q, nil
}

func (qr *QuestionRepository) Update(question *models.QuizQuestion) error {
	query := `
		UPDATE quiz_questions
		SET question = $1, option_a = $2, option_b = $3, option_c = $4, option_d = $5, correct_answer = $6
		WHERE id = $7
	`

	result, err := qr.db.Exec(query, question.Question, question.OptionA, question.OptionB, question.OptionC, question.OptionD, question.CorrectAnswer, question.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("question not found")
	}

	return nil
}
