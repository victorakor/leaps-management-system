package services

import (
	"context"
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type QuizService struct {
	db *sql.DB
}

func NewQuizService(db *sql.DB) *QuizService {
	return &QuizService{db: db}
}

func (qs *QuizService) CreateQuiz(ctx context.Context, quiz *models.Quiz) (*models.Quiz, error) {
	if quiz.Title == "" {
		return nil, errors.New("title is required")
	}
	query := `INSERT INTO quizzes (title, subject_id, class_id, duration, is_active, created_at)
		VALUES ($1,$2,$3,$4,true,NOW()) RETURNING id, created_at`
	err := qs.db.QueryRowContext(ctx, query, quiz.Title, quiz.SubjectID, quiz.ClassID, quiz.Duration).
		Scan(&quiz.ID, &quiz.CreatedAt)
	return quiz, err
}

func (qs *QuizService) AddQuestion(ctx context.Context, question *models.Question) (*models.Question, error) {
	if question.QuizID == "" {
		return nil, errors.New("quiz_id is required")
	}
	query := `INSERT INTO questions (quiz_id, text, answer) VALUES ($1,$2,$3) RETURNING id`
	err := qs.db.QueryRowContext(ctx, query, question.QuizID, question.Text, question.Answer).Scan(&question.ID)
	return question, err
}

func (qs *QuizService) StartQuiz(ctx context.Context, quizID, studentID string) (*models.Attempt, error) {
	attempt := &models.Attempt{QuizID: quizID, StudentID: studentID}
	query := `INSERT INTO attempts (quiz_id, student_id, started_at) VALUES ($1,$2,NOW()) RETURNING id, started_at`
	err := qs.db.QueryRowContext(ctx, query, quizID, studentID).Scan(&attempt.ID, &attempt.StartedAt)
	return attempt, err
}

func (qs *QuizService) SubmitQuiz(ctx context.Context, attemptID string, answers map[string]string) (*models.Attempt, error) {
	var attempt models.Attempt
	err := qs.db.QueryRowContext(ctx, `SELECT id, quiz_id, student_id FROM attempts WHERE id=$1`, attemptID).
		Scan(&attempt.ID, &attempt.QuizID, &attempt.StudentID)
	if err == sql.ErrNoRows {
		return nil, errors.New("attempt not found")
	}
	attempt.Score = float64(len(answers)) * 10 // simplified scoring
	qs.db.ExecContext(ctx, `UPDATE attempts SET score=$1, ended_at=NOW() WHERE id=$2`, attempt.Score, attemptID)
	return &attempt, nil
}

func (qs *QuizService) GetQuizResults(ctx context.Context, studentID string) ([]*models.Attempt, error) {
	rows, err := qs.db.QueryContext(ctx, `SELECT id, quiz_id, student_id, score, started_at, ended_at FROM attempts WHERE student_id=$1`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*models.Attempt
	for rows.Next() {
		var a models.Attempt
		rows.Scan(&a.ID, &a.QuizID, &a.StudentID, &a.Score, &a.StartedAt, &a.EndedAt)
		results = append(results, &a)
	}
	return results, rows.Err()
}

func (qs *QuizService) GetQuiz(ctx context.Context, id string) (*models.Quiz, error) {
	var q models.Quiz
	err := qs.db.QueryRowContext(ctx, `SELECT id, title, subject_id, class_id, duration, is_active, created_at FROM quizzes WHERE id=$1`, id).
		Scan(&q.ID, &q.Title, &q.SubjectID, &q.ClassID, &q.Duration, &q.IsActive, &q.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("quiz not found")
	}
	return &q, err
}
