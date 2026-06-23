package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"leaps/internal/models"
)

type StudentService struct {
	db *sql.DB
}

func NewStudentService(db *sql.DB) *StudentService {
	return &StudentService{db: db}
}

func (ss *StudentService) CreateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	if student.Email == "" {
		return nil, errors.New("email is required")
	}
	admissionNumber := fmt.Sprintf("LEA%d%d", student.SchoolID, time.Now().Unix())
	student.AdmissionNumber = admissionNumber
	student.Status = "active"

	query := `
		INSERT INTO students (first_name, last_name, email, phone, class_id, school_id, admission_number, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
		RETURNING id, created_at
	`
	err := ss.db.QueryRowContext(ctx, query,
		student.FirstName, student.LastName, student.Email, student.Phone,
		student.ClassID, student.SchoolID, student.AdmissionNumber, student.Status,
	).Scan(&student.ID, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (ss *StudentService) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	query := `SELECT id, first_name, last_name, email, phone, class_id, school_id, admission_number, status, created_at
		FROM students WHERE id = $1 AND status != 'deleted'`
	var s models.Student
	err := ss.db.QueryRowContext(ctx, query, id).Scan(
		&s.ID, &s.FirstName, &s.LastName, &s.Email, &s.Phone,
		&s.ClassID, &s.SchoolID, &s.AdmissionNumber, &s.Status, &s.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("student not found")
	}
	return &s, err
}

func (ss *StudentService) GetStudentByAdmissionNumber(ctx context.Context, admissionNumber string) (*models.Student, error) {
	query := `SELECT id, first_name, last_name, email, phone, class_id, school_id, admission_number, status, created_at
		FROM students WHERE admission_number = $1 AND status != 'deleted'`
	var s models.Student
	err := ss.db.QueryRowContext(ctx, query, admissionNumber).Scan(
		&s.ID, &s.FirstName, &s.LastName, &s.Email, &s.Phone,
		&s.ClassID, &s.SchoolID, &s.AdmissionNumber, &s.Status, &s.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, errors.New("student not found")
	}
	return &s, err
}

func (ss *StudentService) UpdateStudent(ctx context.Context, student *models.Student) (*models.Student, error) {
	query := `UPDATE students SET first_name=$1, last_name=$2, email=$3, phone=$4, updated_at=NOW()
		WHERE id=$5 RETURNING id, first_name, last_name, email, phone, class_id, school_id, admission_number, status, created_at`
	err := ss.db.QueryRowContext(ctx, query,
		student.FirstName, student.LastName, student.Email, student.Phone, student.ID,
	).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Phone,
		&student.ClassID, &student.SchoolID, &student.AdmissionNumber, &student.Status, &student.CreatedAt)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func (ss *StudentService) TransferStudent(ctx context.Context, id, newClassID int, reason string) error {
	_, err := ss.db.ExecContext(ctx, `UPDATE students SET class_id=$1, updated_at=NOW() WHERE id=$2`, newClassID, id)
	return err
}

func (ss *StudentService) DeactivateStudent(ctx context.Context, id int) error {
	result, err := ss.db.ExecContext(ctx, `UPDATE students SET status='inactive', updated_at=NOW() WHERE id=$1`, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("student not found")
	}
	return nil
}
