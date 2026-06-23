# LEAPS API Documentation

## Base URL
```
http://localhost:8080/api
```

## Health Check
```
GET /health
```

---

## USER MANAGEMENT ENDPOINTS

### Create User
```
POST /api/users
Content-Type: application/json

{
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "08012345678",
  "role": "teacher",
  "school_id": 1
}

Response: 201 Created
{
  "message": "User created successfully",
  "user": {
    "id": 1,
    "first_name": "John",
    "last_name": "Doe",
    "email": "john@example.com",
    "phone": "08012345678",
    "role": "teacher",
    "school_id": 1,
    "status": "active",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Get User by ID
```
GET /api/users/:id

Response: 200 OK
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "08012345678",
  "role": "teacher",
  "school_id": 1,
  "status": "active"
}
```

### Get User by Email
```
GET /api/users/email/:email

Response: 200 OK
{
  "id": 1,
  "first_name": "John",
  "last_name": "Doe",
  "email": "john@example.com",
  "phone": "08012345678",
  "role": "teacher",
  "school_id": 1,
  "status": "active"
}
```

### List Users
```
GET /api/users?school_id=1&page=1&limit=10

Response: 200 OK
{
  "users": [
    {
      "id": 1,
      "first_name": "John",
      "last_name": "Doe",
      "email": "john@example.com",
      "phone": "08012345678",
      "role": "teacher",
      "school_id": 1,
      "status": "active"
    }
  ],
  "page": 1,
  "limit": 10
}
```

### Update User
```
PUT /api/users/:id
Content-Type: application/json

{
  "first_name": "Jane",
  "last_name": "Smith",
  "phone": "08087654321",
  "role": "admin"
}

Response: 200 OK
{
  "message": "User updated successfully",
  "user": {
    "id": 1,
    "first_name": "Jane",
    "last_name": "Smith",
    "email": "john@example.com",
    "phone": "08087654321",
    "role": "admin",
    "school_id": 1,
    "status": "active"
  }
}
```

### Deactivate User
```
DELETE /api/users/:id

Response: 200 OK
{
  "message": "User deactivated successfully"
}
```

### Check User Exists
```
GET /api/users/exists/:email

Response: 200 OK
{
  "exists": true
}
```

---

## STUDENT MANAGEMENT ENDPOINTS

### Create Student
```
POST /api/students
Content-Type: application/json

{
  "first_name": "Alice",
  "last_name": "Johnson",
  "email": "alice@example.com",
  "phone": "08012345678",
  "class_id": 1,
  "school_id": 1
}

Response: 201 Created
{
  "message": "Student created successfully",
  "student": {
    "id": 1,
    "first_name": "Alice",
    "last_name": "Johnson",
    "email": "alice@example.com",
    "phone": "08012345678",
    "admission_number": "SCH/2026/PRI/0001",
    "class_id": 1,
    "school_id": 1,
    "status": "active",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Get Student by ID
```
GET /api/students/:id

Response: 200 OK
{
  "id": 1,
  "first_name": "Alice",
  "last_name": "Johnson",
  "email": "alice@example.com",
  "phone": "08012345678",
  "admission_number": "SCH/2026/PRI/0001",
  "class_id": 1,
  "school_id": 1,
  "status": "active"
}
```

### Get Student by Admission Number
```
GET /api/students/admission/:admission_number

Response: 200 OK
{
  "id": 1,
  "first_name": "Alice",
  "last_name": "Johnson",
  "email": "alice@example.com",
  "admission_number": "SCH/2026/PRI/0001",
  "class_id": 1,
  "school_id": 1,
  "status": "active"
}
```

### Update Student
```
PUT /api/students/:id
Content-Type: application/json

{
  "first_name": "Alicia",
  "last_name": "Johnson",
  "email": "alicia@example.com",
  "phone": "08087654321"
}

Response: 200 OK
{
  "message": "Student updated successfully",
  "student": {
    "id": 1,
    "first_name": "Alicia",
    "last_name": "Johnson",
    "email": "alicia@example.com",
    "phone": "08087654321",
    "admission_number": "SCH/2026/PRI/0001",
    "class_id": 1,
    "school_id": 1,
    "status": "active"
  }
}
```

### Transfer Student
```
POST /api/students/:id/transfer
Content-Type: application/json

{
  "new_class_id": 2,
  "reason": "Promotion to next class"
}

Response: 200 OK
{
  "message": "Student transferred successfully"
}
```

### Deactivate Student
```
DELETE /api/students/:id

Response: 200 OK
{
  "message": "Student deactivated successfully"
}
```

---

## ACADEMIC RESULTS ENDPOINTS

### Submit Scores
```
POST /api/results/scores
Content-Type: application/json

{
  "student_id": 1,
  "subject_id": 1,
  "assignment_avg": 15,
  "test_avg": 18,
  "exam_score": 75,
  "session_id": 1,
  "term_id": 1
}

Response: 201 Created
{
  "message": "Scores submitted successfully",
  "result": {
    "id": 1,
    "student_id": 1,
    "subject_id": 1,
    "ca_score": 16.5,
    "exam_score": 75,
    "total_score": 91.5,
    "grade": "A",
    "status": "draft",
    "session_id": 1,
    "term_id": 1,
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Get Student Results
```
GET /api/results/student/:student_id

Response: 200 OK
{
  "results": [
    {
      "id": 1,
      "student_id": 1,
      "subject_id": 1,
      "ca_score": 16.5,
      "exam_score": 75,
      "total_score": 91.5,
      "grade": "A",
      "status": "published",
      "session_id": 1,
      "term_id": 1
    }
  ]
}
```

### Publish Results
```
POST /api/results/publish
Content-Type: application/json

{
  "class_id": 1,
  "session_id": 1,
  "term_id": 1
}

Response: 200 OK
{
  "message": "Results published successfully"
}
```

### Lock Results
```
POST /api/results/lock
Content-Type: application/json

{
  "class_id": 1,
  "session_id": 1,
  "term_id": 1
}

Response: 200 OK
{
  "message": "Results locked successfully"
}
```

### Generate Report Card
```
GET /api/results/report-card/:student_id?session_id=1&term_id=1

Response: 200 OK
{
  "student_id": 1,
  "student_name": "Alice Johnson",
  "admission_number": "SCH/2026/PRI/0001",
  "class": "Primary 1",
  "session": "2025/2026",
  "term": "First Term",
  "results": [
    {
      "subject": "Mathematics",
      "ca_score": 16.5,
      "exam_score": 75,
      "total_score": 91.5,
      "grade": "A"
    }
  ],
  "generated_at": "2026-06-23T10:00:00Z"
}
```

---

## ADMISSIONS ENDPOINTS

### Submit Application
```
POST /api/admissions/apply
Content-Type: application/json

{
  "first_name": "Bob",
  "last_name": "Smith",
  "email": "bob@example.com",
  "phone": "08012345678",
  "class_id": 1,
  "school_id": 1
}

Response: 201 Created
{
  "message": "Application submitted successfully",
  "application": {
    "id": 1,
    "application_number": "APP/2026/000001",
    "first_name": "Bob",
    "last_name": "Smith",
    "email": "bob@example.com",
    "phone": "08012345678",
    "class_id": 1,
    "school_id": 1,
    "status": "pending",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Schedule Interview
```
POST /api/admissions/schedule-interview
Content-Type: application/json

{
  "application_id": 1,
  "interview_date": "2026-07-01T00:00:00Z",
  "interview_time": "10:00"
}

Response: 201 Created
{
  "message": "Interview scheduled successfully",
  "appointment": {
    "id": 1,
    "application_id": 1,
    "interview_date": "2026-07-01",
    "interview_time": "10:00",
    "status": "scheduled",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Approve Application
```
POST /api/admissions/:id/approve

Response: 200 OK
{
  "message": "Application approved successfully"
}
```

### Reject Application
```
POST /api/admissions/:id/reject
Content-Type: application/json

{
  "reason": "Does not meet admission criteria"
}

Response: 200 OK
{
  "message": "Application rejected successfully"
}
```

### Get Applications
```
GET /api/admissions?status=pending&page=1&limit=10

Response: 200 OK
{
  "applications": [
    {
      "id": 1,
      "application_number": "APP/2026/000001",
      "first_name": "Bob",
      "last_name": "Smith",
      "email": "bob@example.com",
      "phone": "08012345678",
      "class_id": 1,
      "school_id": 1,
      "status": "pending",
      "created_at": "2026-06-23T10:00:00Z"
    }
  ],
  "page": 1,
  "limit": 10
}
```

### Generate Admission Letter
```
GET /api/admissions/:id/letter

Response: 200 OK
{
  "application_id": 1,
  "applicant_name": "Bob Smith",
  "class": "Primary 1",
  "admission_date": "2026-09-01",
  "admission_number": "SCH/2026/PRI/0002",
  "letter_content": "..."
}
```

---

## QUIZ/CBT ENDPOINTS

### Create Quiz
```
POST /api/quizzes
Content-Type: application/json

{
  "title": "Mathematics Quiz 1",
  "description": "Basic arithmetic",
  "class_id": 1,
  "subject_id": 1,
  "duration": 30,
  "total_marks": 50,
  "pass_mark": 25
}

Response: 201 Created
{
  "message": "Quiz created successfully",
  "quiz": {
    "id": 1,
    "title": "Mathematics Quiz 1",
    "description": "Basic arithmetic",
    "class_id": 1,
    "subject_id": 1,
    "duration": 30,
    "total_marks": 50,
    "pass_mark": 25,
    "status": "draft",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Add Question
```
POST /api/quizzes/:quiz_id/questions
Content-Type: application/json

{
  "question_text": "What is 2 + 2?",
  "option_a": "3",
  "option_b": "4",
  "option_c": "5",
  "option_d": "6",
  "correct_option": "B",
  "marks": 5
}

Response: 201 Created
{
  "message": "Question added successfully",
  "question": {
    "id": 1,
    "quiz_id": 1,
    "question_text": "What is 2 + 2?",
    "option_a": "3",
    "option_b": "4",
    "option_c": "5",
    "option_d": "6",
    "correct_option": "B",
    "marks": 5,
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Start Quiz
```
POST /api/quizzes/:quiz_id/start
Content-Type: application/json

{
  "student_id": 1
}

Response: 201 Created
{
  "message": "Quiz started successfully",
  "attempt": {
    "id": 1,
    "quiz_id": 1,
    "student_id": 1,
    "start_time": "2026-06-23T10:00:00Z",
    "end_time": "2026-06-23T10:30:00Z",
    "status": "in_progress"
  }
}
```

### Submit Quiz
```
POST /api/quizzes/attempts/:attempt_id/submit
Content-Type: application/json

{
  "answers": {
    "1": "B",
    "2": "A",
    "3": "C"
  }
}

Response: 200 OK
{
  "message": "Quiz submitted successfully",
  "result": {
    "attempt_id": 1,
    "quiz_id": 1,
    "student_id": 1,
    "score": 45,
    "total_marks": 50,
    "percentage": 90,
    "status": "completed",
    "submitted_at": "2026-06-23T10:30:00Z"
  }
}
```

### Get Quiz Results
```
GET /api/quizzes/results/:student_id

Response: 200 OK
{
  "results": [
    {
      "quiz_id": 1,
      "quiz_title": "Mathematics Quiz 1",
      "score": 45,
      "total_marks": 50,
      "percentage": 90,
      "status": "completed",
      "submitted_at": "2026-06-23T10:30:00Z"
    }
  ]
}
```

### Get Quiz
```
GET /api/quizzes/:id

Response: 200 OK
{
  "id": 1,
  "title": "Mathematics Quiz 1",
  "description": "Basic arithmetic",
  "class_id": 1,
  "subject_id": 1,
  "duration": 30,
  "total_marks": 50,
  "pass_mark": 25,
  "status": "active",
  "questions": [
    {
      "id": 1,
      "question_text": "What is 2 + 2?",
      "option_a": "3",
      "option_b": "4",
      "option_c": "5",
      "option_d": "6",
      "marks": 5
    }
  ]
}
```

---

## FINANCE ENDPOINTS

### Create Fee Structure
```
POST /api/finance/fees
Content-Type: application/json

{
  "class_id": 1,
  "session_id": 1,
  "amount": 50000,
  "description": "School fees for 2025/2026 session"
}

Response: 201 Created
{
  "message": "Fee structure created successfully",
  "fee": {
    "id": 1,
    "class_id": 1,
    "session_id": 1,
    "amount": 50000,
    "description": "School fees for 2025/2026 session",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Record Payment
```
POST /api/finance/payments
Content-Type: application/json

{
  "student_id": 1,
  "amount": 25000,
  "payment_method": "bank_transfer",
  "reference": "TRF/2026/000001"
}

Response: 201 Created
{
  "message": "Payment recorded successfully",
  "receipt": {
    "receipt_number": "RCP/2026/000001",
    "student_id": 1,
    "amount": 25000,
    "payment_method": "bank_transfer",
    "reference": "TRF/2026/000001",
    "status": "completed",
    "created_at": "2026-06-23T10:00:00Z"
  }
}
```

### Get Student Balance
```
GET /api/finance/balance/:student_id

Response: 200 OK
{
  "student_id": 1,
  "student_name": "Alice Johnson",
  "total_fees": 50000,
  "total_paid": 25000,
  "balance": 25000,
  "status": "partial"
}
```

### Get Payment History
```
GET /api/finance/payments/:student_id

Response: 200 OK
{
  "payments": [
    {
      "receipt_number": "RCP/2026/000001",
      "amount": 25000,
      "payment_method": "bank_transfer",
      "reference": "TRF/2026/000001",
      "status": "completed",
      "created_at": "2026-06-23T10:00:00Z"
    }
  ]
}
```

### Get Debtors
```
GET /api/finance/debtors?school_id=1&page=1&limit=10

Response: 200 OK
{
  "debtors": [
    {
      "student_id": 1,
      "student_name": "Alice Johnson",
      "admission_number": "SCH/2026/PRI/0001",
      "class": "Primary 1",
      "total_fees": 50000,
      "total_paid": 25000,
      "outstanding": 25000
    }
  ],
  "page": 1,
  "limit": 10
}
```

### Get Financial Summary
```
GET /api/finance/summary?school_id=1&session_id=1

Response: 200 OK
{
  "school_id": 1,
  "session_id": 1,
  "total_students": 100,
  "total_fees": 5000000,
  "total_paid": 3500000,
  "total_outstanding": 1500000,
  "collection_rate": 70,
  "debtors_count": 30
}
```

---

## Error Responses

### 400 Bad Request
```json
{
  "error": "Invalid request parameters"
}
```

### 404 Not Found
```json
{
  "error": "Resource not found"
}
```

### 500 Internal Server Error
```json
{
  "error": "Internal server error"
}
```

---

## Status Codes

- **200 OK** - Request successful
- **201 Created** - Resource created successfully
- **204 No Content** - Request successful, no content to return
- **400 Bad Request** - Invalid request parameters
- **401 Unauthorized** - Authentication required
- **403 Forbidden** - Access denied
- **404 Not Found** - Resource not found
- **500 Internal Server Error** - Server error

