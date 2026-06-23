package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"leaps/internal/models"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (sr *StudentRepository) Create(student *models.Student) error {
	query := `
		INSERT INTO students (id, user_id, admission_number, current_class_id, admission_date, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`

	_, err := sr.db.Exec(query, student.ID, student.UserID, student.AdmissionNumber, student.CurrentClassID, student.AdmissionDate, student.Status)
	return err
}

func (sr *StudentRepository) GetByID(id string) (*models.Student, error) {
	query := `
		SELECT id, user_id, admission_number, current_class_id, admission_date, status
		FROM students
		WHERE id = $1
	`

	var student models.Student
	err := sr.db.QueryRow(query, id).Scan(
		&student.ID, &student.UserID, &student.AdmissionNumber, &student.CurrentClassID, &student.AdmissionDate, &student.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	return &student, nil
}

func (sr *StudentRepository) GetByAdmissionNumber(admissionNumber string) (*models.Student, error) {
	query := `
		SELECT id, user_id, admission_number, current_class_id, admission_date, status
		FROM students
		WHERE admission_number = $1
	`

	var student models.Student
	err := sr.db.QueryRow(query, admissionNumber).Scan(
		&student.ID, &student.UserID, &student.AdmissionNumber, &student.CurrentClassID, &student.AdmissionDate, &student.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("student not found")
		}
		return nil, err
	}

	return &student, nil
}

func (sr *StudentRepository) Update(student *models.Student) error {
	query := `
		UPDATE students
		SET current_class_id = $1, status = $2
		WHERE id = $3
	`

	result, err := sr.db.Exec(query, student.CurrentClassID, student.Status, student.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("student not found")
	}

	return nil
}

func (sr *StudentRepository) List(classID string, limit int, offset int) ([]models.Student, error) {
	query := `
		SELECT id, user_id, admission_number, current_class_id, admission_date, status
		FROM students
		WHERE current_class_id = $1
		ORDER BY admission_date DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := sr.db.Query(query, classID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(
			&student.ID, &student.UserID, &student.AdmissionNumber, &student.CurrentClassID, &student.AdmissionDate, &student.Status,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, rows.Err()
}

func (sr *StudentRepository) GetByClass(classID string) ([]models.Student, error) {
	query := `
		SELECT id, user_id, admission_number, current_class_id, admission_date, status
		FROM students
		WHERE current_class_id = $1 AND status = 'active'
		ORDER BY admission_number ASC
	`

	rows, err := sr.db.Query(query, classID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		err := rows.Scan(
			&student.ID, &student.UserID, &student.AdmissionNumber, &student.CurrentClassID, &student.AdmissionDate, &student.Status,
		)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, rows.Err()
}

func (sr *StudentRepository) LogClassTransfer(studentID string, oldClassID string, newClassID string) error {
	query := `
		INSERT INTO student_class_history (id, student_id, class_id, session_id, term_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	_, err := sr.db.Exec(query, uuid.New().String(), studentID, oldClassID, nil, nil, time.Now(), time.Now())
	return err
}
