# LEAPS Testing Guide

## Overview

This guide covers unit testing, integration testing, and end-to-end testing for the LEAPS API.

## Test Structure

```
leaps-backend/
├── user_service_test.go          # Unit tests
├── tests/
│   ├── unit/                     # Unit tests
│   ├── integration/              # Integration tests
│   └── e2e/                      # End-to-end tests
└── Makefile                      # Test commands
```

## Unit Testing

### Running Unit Tests

```bash
# Run all unit tests
make test-unit

# Run specific test
go test -v -run TestCreateUser ./...

# Run with coverage
make test-coverage
```

### Writing Unit Tests

Example unit test:

```go
func TestCreateUser(t *testing.T) {
    mockRepo := NewMockUserRepository()
    service := services.NewUserService(mockRepo)

    user := &models.User{
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john@example.com",
        Phone:     "08012345678",
        Role:      "teacher",
        SchoolID:  1,
        Status:    "active",
    }

    createdUser, err := service.CreateUser(context.Background(), user)
    if err != nil {
        t.Fatalf("CreateUser failed: %v", err)
    }

    if createdUser.ID == 0 {
        t.Error("User ID should not be 0")
    }
}
```

### Test Coverage

Generate coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Coverage targets:
- Services: 80%+
- Repositories: 75%+
- Controllers: 70%+
- Overall: 75%+

## Integration Testing

### Database Integration Tests

```bash
# Start test database
docker-compose -f docker-compose.test.yml up -d

# Run integration tests
go test -v -tags=integration ./tests/integration/...

# Stop test database
docker-compose -f docker-compose.test.yml down
```

### API Integration Tests

Test API endpoints:

```bash
# Start services
make up

# Run API tests
go test -v ./tests/integration/api_test.go

# Stop services
make down
```

## End-to-End Testing

### User Workflow Test

```bash
# 1. Create user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "phone": "08012345678",
    "role": "teacher",
    "school_id": 1
  }'

# 2. Get user
curl http://localhost:8080/api/users/1

# 3. Update user
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Jane",
    "last_name": "Smith"
  }'

# 4. Deactivate user
curl -X DELETE http://localhost:8080/api/users/1
```

### Student Enrollment Workflow

```bash
# 1. Create student
curl -X POST http://localhost:8080/api/students \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Alice",
    "last_name": "Johnson",
    "email": "alice@example.com",
    "class_id": 1,
    "school_id": 1
  }'

# 2. Get student
curl http://localhost:8080/api/students/1

# 3. Transfer student
curl -X POST http://localhost:8080/api/students/1/transfer \
  -H "Content-Type: application/json" \
  -d '{
    "new_class_id": 2,
    "reason": "Promotion"
  }'
```

### Result Submission Workflow

```bash
# 1. Submit scores
curl -X POST http://localhost:8080/api/results/scores \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "subject_id": 1,
    "assignment_avg": 15,
    "test_avg": 18,
    "exam_score": 75,
    "session_id": 1,
    "term_id": 1
  }'

# 2. Get results
curl http://localhost:8080/api/results/student/1

# 3. Publish results
curl -X POST http://localhost:8080/api/results/publish \
  -H "Content-Type: application/json" \
  -d '{
    "class_id": 1,
    "session_id": 1,
    "term_id": 1
  }'

# 4. Generate report card
curl http://localhost:8080/api/results/report-card/1?session_id=1&term_id=1
```

### Admission Workflow

```bash
# 1. Submit application
curl -X POST http://localhost:8080/api/admissions/apply \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "Bob",
    "last_name": "Smith",
    "email": "bob@example.com",
    "phone": "08012345678",
    "class_id": 1,
    "school_id": 1
  }'

# 2. Schedule interview
curl -X POST http://localhost:8080/api/admissions/schedule-interview \
  -H "Content-Type: application/json" \
  -d '{
    "application_id": 1,
    "interview_date": "2026-07-01",
    "interview_time": "10:00"
  }'

# 3. Approve application
curl -X POST http://localhost:8080/api/admissions/1/approve

# 4. Generate admission letter
curl http://localhost:8080/api/admissions/1/letter
```

### Payment Workflow

```bash
# 1. Create fee structure
curl -X POST http://localhost:8080/api/finance/fees \
  -H "Content-Type: application/json" \
  -d '{
    "class_id": 1,
    "session_id": 1,
    "amount": 50000,
    "description": "School fees"
  }'

# 2. Record payment
curl -X POST http://localhost:8080/api/finance/payments \
  -H "Content-Type: application/json" \
  -d '{
    "student_id": 1,
    "amount": 25000,
    "payment_method": "bank_transfer",
    "reference": "TRF/2026/000001"
  }'

# 3. Get balance
curl http://localhost:8080/api/finance/balance/1

# 4. Get debtors
curl http://localhost:8080/api/finance/debtors?school_id=1
```

## Performance Testing

### Load Testing with Apache Bench

```bash
# Test single endpoint
ab -n 1000 -c 10 http://localhost:8080/health

# Test with POST data
ab -n 1000 -c 10 -p data.json -T application/json \
  http://localhost:8080/api/users
```

### Load Testing with wrk

```bash
# Install wrk
brew install wrk

# Run load test
wrk -t4 -c100 -d30s http://localhost:8080/health
```

## Security Testing

### SQL Injection Testing

```bash
# Test vulnerable endpoint
curl "http://localhost:8080/api/users/email/test@example.com' OR '1'='1"

# Should return error, not data
```

### XSS Testing

```bash
# Test with XSS payload
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "first_name": "<script>alert(1)</script>",
    "last_name": "Test",
    "email": "test@example.com",
    "phone": "08012345678",
    "role": "teacher",
    "school_id": 1
  }'

# Should sanitize or reject
```

## Continuous Integration

### GitHub Actions Example

```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
    
    steps:
      - uses: actions/checkout@v2
      
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.21
      
      - name: Run tests
        run: go test -v -coverprofile=coverage.out ./...
      
      - name: Upload coverage
        uses: codecov/codecov-action@v2
```

## Test Checklist

- [ ] Unit tests for all services
- [ ] Unit tests for all repositories
- [ ] Integration tests for API endpoints
- [ ] End-to-end workflow tests
- [ ] Error handling tests
- [ ] Input validation tests
- [ ] Performance tests
- [ ] Security tests
- [ ] Database tests
- [ ] Authentication tests

## Best Practices

1. **Test Isolation**: Each test should be independent
2. **Mock External Dependencies**: Use mocks for databases, APIs
3. **Clear Test Names**: Use descriptive test names
4. **Arrange-Act-Assert**: Follow AAA pattern
5. **Test Edge Cases**: Test boundary conditions
6. **Keep Tests Fast**: Avoid slow operations
7. **Maintain Tests**: Update tests with code changes
8. **Coverage Goals**: Aim for 75%+ coverage

## Troubleshooting

### Tests Failing

1. Check logs:
```bash
docker-compose logs
```

2. Verify database:
```bash
docker-compose exec postgres psql -U postgres -d leaps_db
```

3. Run specific test:
```bash
go test -v -run TestName ./...
```

### Slow Tests

1. Profile tests:
```bash
go test -cpuprofile=cpu.prof -memprofile=mem.prof ./...
go tool pprof cpu.prof
```

2. Identify bottlenecks
3. Optimize or mock slow operations

