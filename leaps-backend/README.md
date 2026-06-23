# LEAPS Backend - Go Server

Leadership Preparatory Schools Digital Operating System Backend

## Setup Instructions

### Prerequisites
- Go 1.21+
- PostgreSQL 12+
- Git

### Installation

1. **Clone and navigate to backend**
```bash
cd leaps-backend
```

2. **Install dependencies**
```bash
go mod download
```

3. **Setup PostgreSQL Database**
```bash
createdb leaps
psql leaps < migrations/001_init_schema.sql
```

4. **Configure environment**
```bash
cp .env.example .env
# Edit .env with your database credentials
```

5. **Run the server**
```bash
go run cmd/server/main.go
```

Server will start on `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /api/auth/login` - User login
- `POST /api/auth/register` - User registration
- `GET /api/auth/me` - Get current user

### Students
- `GET /api/students` - List all students
- `POST /api/students` - Create student
- `GET /api/students/:id` - Get student details
- `PUT /api/students/:id` - Update student

### Admissions
- `POST /api/admissions/apply` - Submit application
- `GET /api/admissions` - List applications
- `POST /api/admissions/schedule` - Schedule interview
- `POST /api/admissions/approve` - Approve admission
- `POST /api/admissions/reject` - Reject admission

### Academic Results
- `POST /api/scores/submit` - Submit scores
- `GET /api/results/student/:id` - Get student results
- `POST /api/results/publish` - Publish results
- `POST /api/results/lock` - Lock results

### Quizzes
- `POST /api/quizzes` - Create quiz
- `POST /api/quizzes/:id/start` - Start quiz
- `POST /api/quizzes/:id/submit` - Submit quiz
- `GET /api/quizzes/:id/results` - Get quiz results

### Finance
- `POST /api/fees` - Create fee structure
- `POST /api/payments` - Record payment
- `GET /api/payments/student/:id` - Get student payments

### Reports
- `GET /api/reports/student/:id` - Get student report
- `POST /api/reports/generate` - Generate report
- `GET /api/reports/download/:id` - Download report

## Project Structure

```
leaps-backend/
├── cmd/server/main.go          # Entry point
├── config/                      # Configuration
├── internal/
│   ├── models/                  # Data models
│   ├── controllers/             # Request handlers
│   ├── services/                # Business logic
│   ├── repositories/            # Data access
│   ├── middleware/              # Auth & RBAC
│   ├── routes/                  # Route definitions
│   ├── auth/                    # JWT & password
│   ├── audit/                   # Audit logging
│   ├── admissions/              # Admission logic
│   ├── academics/               # Academic logic
│   ├── finance/                 # Finance logic
│   ├── quizzes/                 # Quiz logic
│   ├── reports/                 # Report generation
│   └── utils/                   # Utilities
├── pkg/
│   ├── logger/                  # Logging
│   ├── response/                # Response formatting
│   └── validator/               # Input validation
├── migrations/                  # Database migrations
├── templates/                   # PDF templates
└── storage/                     # File storage
```

## Database Schema

All tables use UUID primary keys with soft deletes. Key tables:

- `users` - User accounts with roles
- `students` - Student records
- `classes` - Class definitions
- `score_entries` - Raw score input
- `computed_results` - Calculated grades
- `quizzes` - Quiz definitions
- `payments` - Payment records
- `audit_logs` - Action audit trail

## Security Features

- JWT authentication
- Role-based access control (RBAC)
- Password hashing with bcrypt
- Audit logging for all sensitive actions
- SQL injection prevention
- CORS enabled

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o leaps-server cmd/server/main.go
```

## License

Proprietary - Leadership Preparatory Schools
