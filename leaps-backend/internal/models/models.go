package models

import "time"

//
// ============================
// CORE SCHOOL
// ============================
//

type School struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Acronym   string    `json:"acronym"`
	Motto     string    `json:"motto"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email"`
	LogoURL   string    `json:"logo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//
// ============================
// USERS & AUTH SYSTEM
// ============================
//

type User struct {
	ID           int       `json:"id"`
	SchoolID     int       `json:"school_id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	PasswordHash string    `json:"-"`
	Role         string    `json:"role"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Role struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Scope string `json:"scope"`
}

type Permission struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RolePermission struct {
	RoleID       string `json:"role_id"`
	PermissionID string `json:"permission_id"`
}

//
// ============================
// ACADEMIC STRUCTURE
// ============================
//

type Session struct {
	ID        string    `json:"id"`
	SchoolID  string    `json:"school_id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type Term struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	Name      string    `json:"name"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	IsActive  bool      `json:"is_active"`
}

type Class struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Division  string `json:"division"` // nursery | primary | secondary
	SessionID string `json:"session_id"`
}

type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//
// ============================
// STUDENTS SYSTEM
// ============================
//

type Student struct {
	ID              string    `json:"id"`
	UserID          string    `json:"user_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Email           string    `json:"email"`
	Phone           string    `json:"phone"`
	ClassID         string    `json:"class_id"`
	SchoolID        int       `json:"school_id"`
	AdmissionNumber string    `json:"admission_number"`
	CurrentClassID  string    `json:"current_class_id"`
	AdmissionDate   time.Time `json:"admission_date"`
	Status          string    `json:"status"` // active, graduated, transferred
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type StudentClassHistory struct {
	ID        string    `json:"id"`
	StudentID string    `json:"student_id"`
	ClassID   string    `json:"class_id"`
	SessionID string    `json:"session_id"`
	TermID    string    `json:"term_id"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

//
// ============================
// TEACHERS SYSTEM
// ============================
//

type Teacher struct {
	ID         string `json:"id"`
	UserID     string `json:"user_id"`
	StaffID    string `json:"staff_id"`
	Department string `json:"department"`
}

type TeacherClassAssignment struct {
	ID        string `json:"id"`
	TeacherID string `json:"teacher_id"`
	ClassID   string `json:"class_id"`
	SubjectID string `json:"subject_id"`
	SessionID string `json:"session_id"`
	TermID    string `json:"term_id"`
}

//
// ============================
// ACADEMIC SCORES SYSTEM
// ============================
//

type ScoreEntry struct {
	ID        string  `json:"id"`
	StudentID string  `json:"student_id"`
	SubjectID string  `json:"subject_id"`
	TeacherID string  `json:"teacher_id"`
	SessionID string  `json:"session_id"`
	TermID    string  `json:"term_id"`

	Assignment1 int     `json:"assignment_1"`
	Assignment2 int     `json:"assignment_2"`
	Test1       int     `json:"test_1"`
	Test2       int     `json:"test_2"`
	Test3       int     `json:"test_3"`
	Exam        float64 `json:"exam"`
	CA1         float64 `json:"ca1"`
	CA2         float64 `json:"ca2"`
	Total       float64 `json:"total"`
	Grade       string  `json:"grade"`
	Remark      string  `json:"remark"`

	Status string `json:"status"` // draft, submitted, approved
}

type ComputedResult struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	SubjectID string `json:"subject_id"`
	SessionID string `json:"session_id"`
	TermID    string `json:"term_id"`

	TotalScore      int    `json:"total_score"`
	Grade           string `json:"grade"`
	SubjectPosition int    `json:"subject_position"`
	ClassPosition   int    `json:"class_position"`
}

type ReportCard struct {
	ID        string    `json:"id"`
	StudentID string    `json:"student_id"`
	SessionID string    `json:"session_id"`
	TermID    string    `json:"term_id"`

	FileURL   string    `json:"file_url"`
	ResultID  string    `json:"result_id"`
	IsLocked  bool      `json:"is_locked"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//
// ============================
// QUIZ SYSTEM
// ============================
//

type Quiz struct {
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	SubjectID       string    `json:"subject_id"`
	ClassID         string    `json:"class_id"`
	CreatedBy       string    `json:"created_by"`
	Duration        int       `json:"duration"`
	DurationMinutes int       `json:"duration_minutes"`
	StartTime       time.Time `json:"start_time"`
	EndTime         time.Time `json:"end_time"`
	IsActive        bool      `json:"is_active"`
	CreatedAt       time.Time `json:"created_at"`
}

type Question struct {
	ID            string `json:"id"`
	QuizID        string `json:"quiz_id"`
	Text          string `json:"text"`
	Question      string `json:"question"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	Answer        string `json:"answer"`
	CorrectAnswer string `json:"correct_answer"`
}

type QuizQuestion struct {
	ID            string `json:"id"`
	QuizID        string `json:"quiz_id"`
	Question      string `json:"question"`
	OptionA       string `json:"option_a"`
	OptionB       string `json:"option_b"`
	OptionC       string `json:"option_c"`
	OptionD       string `json:"option_d"`
	CorrectAnswer string `json:"correct_answer"`
}

type Attempt struct {
	ID          string    `json:"id"`
	QuizID      string    `json:"quiz_id"`
	StudentID   string    `json:"student_id"`
	Score       float64   `json:"score"`
	StartedAt   time.Time `json:"started_at"`
	EndedAt     time.Time `json:"ended_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	Status      string    `json:"status"`
}

type QuizAttempt struct {
	ID          string    `json:"id"`
	QuizID      string    `json:"quiz_id"`
	StudentID   string    `json:"student_id"`
	Score       int       `json:"score"`
	StartedAt   time.Time `json:"started_at"`
	SubmittedAt time.Time `json:"submitted_at"`
	Status      string    `json:"status"`
}

//
// ============================
// FINANCE SYSTEM
// ============================
//

type Fee struct {
	ID        string    `json:"id"`
	ClassID   string    `json:"class_id"`
	SessionID string    `json:"session_id"`
	TermID    string    `json:"term_id"`
	Amount    float64   `json:"amount"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Payment struct {
	ID        string    `json:"id"`
	StudentID string    `json:"student_id"`
	FeeID     string    `json:"fee_id"`
	Amount    float64   `json:"amount"`
	Method    string    `json:"method"`
	ReceiptNo string    `json:"receipt_no"`
	PaidAt    time.Time `json:"paid_at"`
}

//
// ============================
// ADMISSION SYSTEM
// ============================
//

type Application struct {
	ID                string    `json:"id"`
	ApplicationNumber string    `json:"application_number"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	ClassID           string    `json:"class_id"`
	DesiredClass      string    `json:"desired_class"`
	Status            string    `json:"status"`
	AppliedAt         time.Time `json:"applied_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type Appointment struct {
	ID            string    `json:"id"`
	ApplicationID string    `json:"application_id"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	Date          string    `json:"date"`
	Time          string    `json:"time"`
	Status        string    `json:"status"`
	Notes         string    `json:"notes"`
}

type AdmissionAppointment struct {
	ID            string    `json:"id"`
	ApplicationID string    `json:"application_id"`
	ScheduledAt   time.Time `json:"scheduled_at"`
	Date          string    `json:"date"`
	Time          string    `json:"time"`
	Status        string    `json:"status"`
	Notes         string    `json:"notes"`
}

//
// ============================
// AUDIT SYSTEM
// ============================
//

type AuditLog struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Action      string    `json:"action"`
	TableName   string    `json:"table_name"`
	RecordID    string    `json:"record_id"`
	BeforeState string    `json:"before_state"`
	AfterState  string    `json:"after_state"`
	IPAddress   string    `json:"ip_address"`
	CreatedAt   time.Time `json:"created_at"`
}

//
// ============================
// AUTH DTOs
// ============================
//

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}
