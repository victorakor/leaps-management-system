package services

import (
	"context"
	"database/sql"
	"errors"

	"leaps/internal/models"
)

type FinanceService struct {
	db *sql.DB
}

func NewFinanceService(db *sql.DB) *FinanceService {
	return &FinanceService{db: db}
}

func (fs *FinanceService) CreateFeeStructure(ctx context.Context, fee *models.Fee) (*models.Fee, error) {
	if fee.Amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	query := `INSERT INTO fees (class_id, term_id, amount, category) VALUES ($1,$2,$3,$4) RETURNING id`
	err := fs.db.QueryRowContext(ctx, query, fee.ClassID, fee.TermID, fee.Amount, fee.Category).Scan(&fee.ID)
	return fee, err
}

func (fs *FinanceService) RecordPayment(ctx context.Context, payment *models.Payment) (*models.Payment, error) {
	if payment.StudentID == "" {
		return nil, errors.New("student_id is required")
	}
	query := `INSERT INTO payments (student_id, fee_id, amount, method, paid_at) VALUES ($1,$2,$3,$4,NOW()) RETURNING id, paid_at`
	err := fs.db.QueryRowContext(ctx, query, payment.StudentID, payment.FeeID, payment.Amount, payment.Method).
		Scan(&payment.ID, &payment.PaidAt)
	return payment, err
}

func (fs *FinanceService) GetStudentBalance(ctx context.Context, studentID string) (map[string]interface{}, error) {
	var totalFees, totalPaid float64
	fs.db.QueryRowContext(ctx, `SELECT COALESCE(SUM(f.amount),0) FROM fees f JOIN students s ON s.class_id=f.class_id WHERE s.id=$1`, studentID).Scan(&totalFees)
	fs.db.QueryRowContext(ctx, `SELECT COALESCE(SUM(amount),0) FROM payments WHERE student_id=$1`, studentID).Scan(&totalPaid)
	return map[string]interface{}{
		"student_id":  studentID,
		"total_fees":  totalFees,
		"total_paid":  totalPaid,
		"balance_due": totalFees - totalPaid,
	}, nil
}

func (fs *FinanceService) GetPaymentHistory(ctx context.Context, studentID string) ([]*models.Payment, error) {
	rows, err := fs.db.QueryContext(ctx, `SELECT id, student_id, fee_id, amount, method, paid_at FROM payments WHERE student_id=$1 ORDER BY paid_at DESC`, studentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var payments []*models.Payment
	for rows.Next() {
		var p models.Payment
		rows.Scan(&p.ID, &p.StudentID, &p.FeeID, &p.Amount, &p.Method, &p.PaidAt)
		payments = append(payments, &p)
	}
	return payments, rows.Err()
}

func (fs *FinanceService) GetDebtors(ctx context.Context, schoolID string, page, limit int) ([]map[string]interface{}, error) {
	_ = schoolID
	return []map[string]interface{}{}, nil
}

func (fs *FinanceService) GetFinancialSummary(ctx context.Context, schoolID, sessionID string) (map[string]interface{}, error) {
	var totalCollected float64
	fs.db.QueryRowContext(ctx, `SELECT COALESCE(SUM(amount),0) FROM payments`).Scan(&totalCollected)
	return map[string]interface{}{
		"school_id":       schoolID,
		"session_id":      sessionID,
		"total_collected": totalCollected,
	}, nil
}

// Satisfy unused import
var _ = errors.New
