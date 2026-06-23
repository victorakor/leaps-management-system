package audit

import (
	"database/sql"
	"encoding/json"
	"time"

	"leaps/internal/models"

	"github.com/google/uuid"
)

type AuditLogger struct {
	db *sql.DB
}

func NewAuditLogger(db *sql.DB) *AuditLogger {
	return &AuditLogger{db: db}
}

func (al *AuditLogger) LogAction(userID, action, tableName, recordID, ipAddress string, beforeState, afterState interface{}) error {
	id := uuid.New().String()

	beforeJSON, _ := json.Marshal(beforeState)
	afterJSON, _ := json.Marshal(afterState)

	query := `
		INSERT INTO audit_logs (id, user_id, action, table_name, record_id, before_state, after_state, ip_address, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	_, err := al.db.Exec(query, id, userID, action, tableName, recordID, string(beforeJSON), string(afterJSON), ipAddress, time.Now())
	return err
}

func (al *AuditLogger) GetLogs(limit int) ([]models.AuditLog, error) {
	query := `
		SELECT id, user_id, action, table_name, record_id, before_state, after_state, ip_address, created_at
		FROM audit_logs
		ORDER BY created_at DESC
		LIMIT $1
	`

	rows, err := al.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []models.AuditLog
	for rows.Next() {
		var log models.AuditLog
		err := rows.Scan(&log.ID, &log.UserID, &log.Action, &log.TableName, &log.RecordID, &log.BeforeState, &log.AfterState, &log.IPAddress, &log.CreatedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}

	return logs, nil
}
