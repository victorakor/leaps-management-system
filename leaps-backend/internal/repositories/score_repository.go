package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type ScoreRepository struct {
	db *sql.DB
}

func NewScoreRepository(db *sql.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (sr *ScoreRepository) Create(score *models.ScoreEntry) error {
	query := `
		INSERT INTO score_entries (id, student_id, subject_id, teacher_id, session_id, term_id, assignment_1, assignment_2, test_1, test_2, test_3, exam, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW())
	`

	_, err := sr.db.Exec(query, score.ID, score.StudentID, score.SubjectID, score.TeacherID, score.SessionID, score.TermID, score.Assignment1, score.Assignment2, score.Test1, score.Test2, score.Test3, score.Exam, score.Status)
	return err
}

func (sr *ScoreRepository) GetByStudent(studentID string, sessionID string, termID string) ([]models.ScoreEntry, error) {
	query := `
		SELECT id, student_id, subject_id, teacher_id, session_id, term_id, assignment_1, assignment_2, test_1, test_2, test_3, exam, status
		FROM score_entries
		WHERE student_id = $1 AND session_id = $2 AND term_id = $3
		ORDER BY subject_id ASC
	`

	rows, err := sr.db.Query(query, studentID, sessionID, termID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []models.ScoreEntry
	for rows.Next() {
		var score models.ScoreEntry
		err := rows.Scan(
			&score.ID, &score.StudentID, &score.SubjectID, &score.TeacherID, &score.SessionID, &score.TermID,
			&score.Assignment1, &score.Assignment2, &score.Test1, &score.Test2, &score.Test3, &score.Exam, &score.Status,
		)
		if err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	return scores, rows.Err()
}

func (sr *ScoreRepository) Update(score *models.ScoreEntry) error {
	query := `
		UPDATE score_entries
		SET assignment_1 = $1, assignment_2 = $2, test_1 = $3, test_2 = $4, test_3 = $5, exam = $6, status = $7, updated_at = NOW()
		WHERE id = $8
	`

	result, err := sr.db.Exec(query, score.Assignment1, score.Assignment2, score.Test1, score.Test2, score.Test3, score.Exam, score.Status, score.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("score entry not found")
	}

	return nil
}

func (sr *ScoreRepository) GetByID(id string) (*models.ScoreEntry, error) {
	query := `
		SELECT id, student_id, subject_id, teacher_id, session_id, term_id, assignment_1, assignment_2, test_1, test_2, test_3, exam, status
		FROM score_entries
		WHERE id = $1
	`

	var score models.ScoreEntry
	err := sr.db.QueryRow(query, id).Scan(
		&score.ID, &score.StudentID, &score.SubjectID, &score.TeacherID, &score.SessionID, &score.TermID,
		&score.Assignment1, &score.Assignment2, &score.Test1, &score.Test2, &score.Test3, &score.Exam, &score.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("score entry not found")
		}
		return nil, err
	}

	return &score, nil
}
