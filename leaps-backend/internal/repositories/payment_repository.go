package repositories

import (
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (pr *PaymentRepository) Create(payment *models.Payment) error {
	query := `
		INSERT INTO payments (id, student_id, amount, method, receipt_no, paid_at, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
	`

	_, err := pr.db.Exec(query, payment.ID, payment.StudentID, payment.Amount, payment.Method, payment.ReceiptNo, payment.PaidAt)
	return err
}

func (pr *PaymentRepository) GetByID(id string) (*models.Payment, error) {
	query := `
		SELECT id, student_id, amount, method, receipt_no, paid_at
		FROM payments
		WHERE id = $1
	`

	var payment models.Payment
	err := pr.db.QueryRow(query, id).Scan(
		&payment.ID, &payment.StudentID, &payment.Amount, &payment.Method, &payment.ReceiptNo, &payment.PaidAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("payment not found")
		}
		return nil, err
	}

	return &payment, nil
}

func (pr *PaymentRepository) GetByStudent(studentID string) ([]models.Payment, error) {
	query := `
		SELECT id, student_id, amount, method, receipt_no, paid_at
		FROM payments
		WHERE student_id = $1
		ORDER BY paid_at DESC
	`

	rows, err := pr.db.Query(query, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID, &payment.StudentID, &payment.Amount, &payment.Method, &payment.ReceiptNo, &payment.PaidAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, rows.Err()
}

func (pr *PaymentRepository) GetTotalByStudent(studentID string, sessionID string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(amount), 0)
		FROM payments
		WHERE student_id = $1 AND EXTRACT(YEAR FROM paid_at) = EXTRACT(YEAR FROM NOW())
	`

	var total float64
	err := pr.db.QueryRow(query, studentID).Scan(&total)
	return total, err
}

func (pr *PaymentRepository) GetTotalBySession(sessionID string) (float64, error) {
	query := `
		SELECT COALESCE(SUM(amount), 0)
		FROM payments
		WHERE EXTRACT(YEAR FROM paid_at) = EXTRACT(YEAR FROM NOW())
	`

	var total float64
	err := pr.db.QueryRow(query).Scan(&total)
	return total, err
}

func (pr *PaymentRepository) GetDebtors(sessionID string) ([]map[string]interface{}, error) {
	query := `
		SELECT s.id, u.full_name, (f.amount - COALESCE(SUM(p.amount), 0)) as balance
		FROM students s
		JOIN users u ON s.user_id = u.id
		JOIN fees f ON s.current_class_id = f.class_id
		LEFT JOIN payments p ON s.id = p.student_id
		WHERE f.session_id = $1
		GROUP BY s.id, u.full_name, f.amount
		HAVING (f.amount - COALESCE(SUM(p.amount), 0)) > 0
		ORDER BY balance DESC
	`

	rows, err := pr.db.Query(query, sessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var debtors []map[string]interface{}
	for rows.Next() {
		var id, name string
		var balance float64
		err := rows.Scan(&id, &name, &balance)
		if err != nil {
			return nil, err
		}
		debtors = append(debtors, map[string]interface{}{
			"student_id": id,
			"name":       name,
			"balance":    balance,
		})
	}

	return debtors, rows.Err()
}
