package services

import (
	"context"
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type ResultService struct {
	db *sql.DB
}

func NewResultService(db *sql.DB) *ResultService {
	return &ResultService{db: db}
}

func (rs *ResultService) SubmitScores(ctx context.Context, score *models.ScoreEntry) (*models.ScoreEntry, error) {
	if score.StudentID == "" {
		return nil, errors.New("student_id is required")
	}
	score.Total = score.CA1 + score.CA2 + score.Exam
	score.Grade = computeGrade(score.Total)
	query := `INSERT INTO score_entries (student_id, subject_id, term_id, ca1, ca2, exam, total, grade, remark)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id`
	err := rs.db.QueryRowContext(ctx, query,
		score.StudentID, score.SubjectID, score.TermID,
		score.CA1, score.CA2, score.Exam, score.Total, score.Grade, score.Remark,
	).Scan(&score.ID)
	return score, err
}

func computeGrade(total float64) string {
	switch {
	case total >= 70:
		return "A"
	case total >= 60:
		return "B"
	case total >= 50:
		return "C"
	case total >= 40:
		return "D"
	default:
		return "F"
	}
}

func (rs *ResultService) GetStudentResults(ctx context.Context, studentID string) ([]*models.ScoreEntry, error) {
	rows, err := rs.db.QueryContext(ctx,
		`SELECT id, student_id, subject_id, term_id, ca1, ca2, exam, total, grade, remark
		FROM score_entries WHERE student_id = $1`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var results []*models.ScoreEntry
	for rows.Next() {
		var s models.ScoreEntry
		rows.Scan(&s.ID, &s.StudentID, &s.SubjectID, &s.TermID, &s.CA1, &s.CA2, &s.Exam, &s.Total, &s.Grade, &s.Remark)
		results = append(results, &s)
	}
	return results, rows.Err()
}

func (rs *ResultService) PublishResults(ctx context.Context, classID, sessionID, termID string) error {
	_, err := rs.db.ExecContext(ctx,
		`UPDATE results SET is_published=true WHERE term_id=$1`, termID)
	return err
}

func (rs *ResultService) LockResults(ctx context.Context, classID, sessionID, termID string) error {
	_, err := rs.db.ExecContext(ctx,
		`UPDATE results SET is_locked=true WHERE term_id=$1`, termID)
	return err
}

func (rs *ResultService) GenerateReportCard(ctx context.Context, studentID, sessionID, termID string) (map[string]interface{}, error) {
	results, err := rs.GetStudentResults(ctx, studentID)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"student_id": studentID,
		"term_id":    termID,
		"session_id": sessionID,
		"results":    results,
	}, nil
}
