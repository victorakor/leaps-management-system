
---

# 🧠 LEAPS SCHOOL OPERATING SYSTEM — FULL IMPLEMENTATION SPECIFICATION

## PROJECT NAME

**LEAPS — Leadership Preparatory Schools Digital Operating System**

📍 Makurdi, Benue State, Nigeria

---

# 1. SYSTEM DEFINITION

This system is a **full institutional operating system** for a school covering:

* Admissions lifecycle (online → physical → enrollment)
* Academic record management
* Automated grading and report generation
* Quiz / CBT system with anti-cheat enforcement
* Finance and bursar management system
* Timetable management
* Staff role management system
* Student lifecycle tracking
* School media/blog system
* Document generation system (PDF with branding)
* Audit logging system for all sensitive actions

It replaces all manual school operations.

---

# 2. CORE ARCHITECTURE PRINCIPLES

## 2.1 SINGLE SOURCE OF TRUTH

All data must originate from database records. No frontend calculations are trusted.

---

## 2.2 APPEND-ONLY ACADEMIC HISTORY

Never overwrite:

* results
* attendance
* class history
* admission decisions

Instead:

* create new versioned records with timestamps

---

## 2.3 ROLE-BASED ACCESS CONTROL (RBAC)

Every action must be validated by:

```text
Role → Permission → Resource → Action
```

No direct feature access without permission check.

---

## 2.4 AUDIT LOGGING (MANDATORY)

Every sensitive action stores:

* actor_id
* role
* timestamp
* IP address
* before/after state
* action type

---

# 3. TECH STACK (STRICT REQUIREMENTS)

## Backend

* Go (Golang)
* Gin (preferred) or Fiber
* PostgreSQL

## Frontend

* HTML5
* CSS3 (TailwindCSS or custom system)
* Vanilla JavaScript (no React/Vue required)
* Optional: Alpine.js

## PDF SYSTEM

* HTML → PDF rendering using:

  * Puppeteer (recommended)
  * OR chromedp (Go-native)
  * OR wkhtmltopdf fallback

## STORAGE

* Local storage OR S3-compatible object storage

---

# 4. GLOBAL SCHOOL BRANDING SYSTEM

All generated documents MUST include:

## SCHOOL IDENTITY

* Name: Leadership Preparatory Schools (LEAPS)
* Location: Makurdi, Benue State
* Motto: configurable
* Logo: uploaded dynamically by Owner

---

## DOCUMENT HEADER STANDARD

Every PDF MUST include:

* Logo (center/left aligned)
* School name (bold uppercase)
* Motto (italic optional)
* Address
* Contact info

Used in:

* admission letters
* report cards
* quiz certificates
* receipts
* official letters

---

# 5. ROLE SYSTEM (FULL DETAILED BEHAVIOR SPECIFICATION)

---

# 5.1 SCHOOL OWNER (PROPRIETOR)

## PURPOSE

Ultimate authority over entire school system.

---

## DASHBOARDS

* Global Admin Dashboard
* Admissions Control
* Finance Overview (all divisions)
* Academic Overview (all divisions)
* Staff Management
* System Configuration Panel
* Audit Logs Panel

---

## PERMISSIONS

### Academic Control

* Create sessions/terms
* Define grading system (A–F or weighted)
* Set promotion rules

### Admissions Control

* Open/close admission window
* Define eligible classes
* Override admission decisions
* Approve final admissions

### Staff Control

* Create roles
* Promote/demote staff
* Suspend users

### Finance Oversight

* View all transactions
* View debts
* View revenue reports

### System Control

* Upload logo
* Configure school branding
* Manage document templates

---

## RESTRICTIONS

* Cannot bypass audit logs
* Cannot access system code
* Cannot silently modify academic history

---

# 5.2 PRINCIPAL (SECONDARY HEAD)

## PURPOSE

Full control of secondary academic system.

---

## DASHBOARD

* Secondary Academic Dashboard
* Secondary Results Panel
* Secondary Quiz Management
* Secondary Timetable
* Secondary Admissions Review

---

## PERMISSIONS

* Approve secondary results
* Publish secondary results
* Manage secondary teachers
* Manage secondary quizzes
* Review admission applications (secondary)
* Post announcements

---

## RESTRICTIONS

* No primary access
* No finance modification
* No global system control

---

# 5.3 HEAD TEACHER (PRIMARY/NURSERY)

Same as Principal but scoped ONLY to:

* Nursery
* Primary

---

# 5.4 VICE PRINCIPAL / ASSISTANT HEAD

## PURPOSE

Academic support + moderation

---

## PERMISSIONS

* Review results before publication
* Assist quiz moderation
* Monitor academic integrity
* Support timetable management

---

## RESTRICTIONS

* Cannot publish results alone
* Cannot override principal/owner

---

# 5.5 BURSAR (PRIMARY / SECONDARY)

## PURPOSE

Financial operations only

---

## PERMISSIONS

* Create fee structures
* Record payments
* Generate receipts
* Track debts
* Financial reporting

---

## RESTRICTIONS

* No academic access
* No quiz access
* No admission control

---

# 5.6 TEACHER

## PURPOSE

Academic data entry only

---

## DASHBOARD

* My Classes
* Score Entry Panel
* Attendance Panel
* Assignment Panel
* Quiz Contribution Panel

---

## PERMISSIONS

* Enter raw scores only
* Upload assignments
* Mark attendance
* Submit quiz questions

---

## STRICT RULE

Teachers MUST NOT:

* compute totals
* compute grades
* compute positions
* publish results

---

# 5.7 CLASS TEACHER (DESIGNATION)

## EXTRA PERMISSIONS

* Generate class reports
* Communicate with parents
* Monitor attendance anomalies

---

# 5.8 EXAM / ACADEMIC OFFICER

## PURPOSE

Academic validation authority

---

## PERMISSIONS

* Approve quiz questions
* Moderate exam difficulty
* Validate results before publication
* Ensure grading integrity

---

# 5.9 ADMISSIONS OFFICER

## PURPOSE

Admission pipeline execution

---

## PERMISSIONS

* Review applications
* Generate interview schedules
* Create appointment slips
* Verify documents
* Recommend admission approval

---

## RESTRICTIONS

* Cannot approve final admission alone
* Cannot create students manually

---

# 5.10 STUDENT (SECONDARY)

## DASHBOARD

* Results
* Quizzes
* Timetable
* Assignments

---

## PERMISSIONS

* Take quizzes
* View results
* View announcements

---

# 5.11 PUPIL (PRIMARY/NURSERY)

Simplified version of student role:

* quizzes
* results
* assignments
* announcements

---

# 5.12 PARENT / GUARDIAN

## PURPOSE

Observation + accountability

---

## PERMISSIONS

* View child results
* View attendance
* View fees
* Download admission letters
* Print reports
* Receive notifications

---

# 6. ADMISSION SYSTEM (FULL WORKFLOW)

---

## 6.1 APPLICATION WINDOW CONTROL

Owner defines:

* open date
* close date
* allowed classes

---

## 6.2 APPLICATION FORM

Parent submits:

* child info
* class selection
* parent info
* contact details

System generates:

```text
APP/2026/000145
```

---

## 6.3 INTERVIEW SCHEDULING ENGINE

Rules:

* Monday–Friday only
* 9:00 AM – 2:00 PM only
* slot capacity enforced
* no duplicate bookings

Output:

Printable appointment slip.

---

## 6.4 PHYSICAL VERIFICATION

Admissions officer:

* verifies documents
* confirms attendance
* marks status

---

## 6.5 FINAL ADMISSION

Decision:

* APPROVED / REJECTED

---

## 6.6 ADMISSION ID GENERATION

```text
SCH/2026/PRI/0001
```

Rules:

* sequential
* unique
* transaction-safe

---

## 6.7 STUDENT CREATION

On approval:

* create student account
* assign class
* activate portal

---

## 7. ACADEMIC RESULT ENGINE

---

## 7.1 INPUT RULE

Teachers input ONLY raw scores:

* assignments
* tests
* exams

---

## 7.2 SYSTEM CALCULATIONS

System computes:

* totals
* grades
* subject position
* class position
* averages

---

## 7.3 GRADING SCALE

A: 70–100
B: 60–69
C: 50–59
D: 45–49
E: 40–44
F: 0–39

---

## 7.4 RESULT STATES

* DRAFT
* PENDING APPROVAL
* PUBLISHED
* LOCKED

LOCKED = immutable

---

## 7.5 REPORT CARD GENERATION

Must include:

* school branding header
* student details
* academic table
* attendance
* psychomotor traits
* affective traits
* remarks
* QR verification
* document ID
* signature area

---

# 8. QUIZ / CBT SYSTEM

---

## FEATURES

* holiday quizzes
* class quizzes
* exams

---

## ANTI-CHEAT SYSTEM

Detect:

* tab switching
* copy/paste
* fullscreen exit
* refresh attempts

Enforce:

* fullscreen mode
* timer lock
* question shuffle
* option shuffle

---

# 9. FINANCE SYSTEM

Separate per division:

* Primary bursar
* Secondary bursar

Functions:

* fees
* receipts
* debts
* scholarships
* reports

---

# 10. BLOG / MEDIA SYSTEM

Instagram-style feed:

* images/videos
* school events
* announcements

---

# 11. TIMETABLE SYSTEM

* class timetable
* exam timetable
* teacher timetable

Includes:

* conflict detection
* publishing workflow

---

# 12. HISTORICAL DATA SYSTEM

Must preserve:

* student movement between classes
* teacher assignments
* result history
* attendance history

NO DELETIONS.

---

# 13. UI / UX REQUIREMENTS

---

## DESIGN STYLE

* modern SaaS school dashboard
* glassmorphism
* soft shadows
* card-based layout

---

## ANIMATIONS

* page transitions (fade/slide)
* modal scaling animations
* hover glow effects
* sidebar animations
* skeleton loading states

---

## RESPONSIVENESS

* mobile-first
* tablet optimized
* desktop full dashboard grid

---

# 14. SECURITY REQUIREMENTS

* JWT authentication
* RBAC middleware enforcement
* audit logs everywhere
* result locking enforcement
* admission verification system
* QR code verification system

---

# 15. DATABASE PRINCIPLES

* UUID primary keys
* sequential public IDs for humans
* strict relational integrity
* no destructive updates
* full history tracking tables

---

# 16. FINAL SYSTEM OUTPUT DEFINITION

This system must function as:

> A full digital school operating system replacing all administrative, academic, financial, and communication workflows of LEAPS with automated, secure, and auditable processes.

---

# 🗄️ 1. FULL POSTGRESQL DATABASE SCHEMA (LEAPS)

## 🔑 GLOBAL DESIGN RULES

* All tables use `UUID` primary keys
* All tables include:

  * `created_at`
  * `updated_at`
  * `deleted_at` (soft delete)
* All sensitive actions logged in `audit_logs`
* Academic records are **append-only (no overwrite)**

---

# 🏫 1. SCHOOL CORE TABLES

## schools

```sql
CREATE TABLE schools (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    acronym VARCHAR(20),
    motto TEXT,
    address TEXT,
    phone VARCHAR(50),
    email VARCHAR(255),
    logo_url TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

---

## sessions

```sql
CREATE TABLE sessions (
    id UUID PRIMARY KEY,
    school_id UUID REFERENCES schools(id),
    name VARCHAR(20), -- 2026/2027
    is_active BOOLEAN DEFAULT false,
    created_at TIMESTAMP
);
```

---

## terms

```sql
CREATE TABLE terms (
    id UUID PRIMARY KEY,
    session_id UUID REFERENCES sessions(id),
    name VARCHAR(20), -- First Term
    start_date DATE,
    end_date DATE,
    is_active BOOLEAN
);
```

---

# 👥 2. USER & ROLE SYSTEM

## users

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    school_id UUID REFERENCES schools(id),
    full_name VARCHAR(255),
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(50),
    password_hash TEXT,
    role_id UUID,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP
);
```

---

## roles

```sql
CREATE TABLE roles (
    id UUID PRIMARY KEY,
    name VARCHAR(100),
    scope VARCHAR(50) -- owner, admin, teacher, student, parent
);
```

---

## permissions

```sql
CREATE TABLE permissions (
    id UUID PRIMARY KEY,
    name VARCHAR(100) -- manage_students, publish_results
);
```

---

## role_permissions

```sql
CREATE TABLE role_permissions (
    role_id UUID REFERENCES roles(id),
    permission_id UUID REFERENCES permissions(id),
    PRIMARY KEY(role_id, permission_id)
);
```

---

# 🎓 3. STUDENTS SYSTEM

## students

```sql
CREATE TABLE students (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    admission_number VARCHAR(50) UNIQUE,
    current_class_id UUID,
    admission_date DATE,
    status VARCHAR(20) -- active, graduated, transferred
);
```

---

## student_class_history

```sql
CREATE TABLE student_class_history (
    id UUID PRIMARY KEY,
    student_id UUID,
    class_id UUID,
    session_id UUID,
    term_id UUID,
    start_date DATE,
    end_date DATE
);
```

---

## classes

```sql
CREATE TABLE classes (
    id UUID PRIMARY KEY,
    name VARCHAR(50), -- Primary 5, JSS1
    division VARCHAR(20), -- primary, secondary, nursery
    session_id UUID
);
```

---

# 👨‍🏫 4. TEACHERS SYSTEM

## teachers

```sql
CREATE TABLE teachers (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    staff_id VARCHAR(50),
    department VARCHAR(100)
);
```

---

## teacher_class_assignments

```sql
CREATE TABLE teacher_class_assignments (
    id UUID PRIMARY KEY,
    teacher_id UUID,
    class_id UUID,
    subject VARCHAR(100),
    session_id UUID,
    term_id UUID
);
```

---

# 📝 5. ACADEMIC RESULTS SYSTEM

## subjects

```sql
CREATE TABLE subjects (
    id UUID PRIMARY KEY,
    name VARCHAR(100)
);
```

---

## score_entries (RAW INPUT ONLY)

```sql
CREATE TABLE score_entries (
    id UUID PRIMARY KEY,
    student_id UUID,
    subject_id UUID,
    teacher_id UUID,
    session_id UUID,
    term_id UUID,
    assignment_1 INT,
    assignment_2 INT,
    test_1 INT,
    test_2 INT,
    test_3 INT,
    exam INT,
    status VARCHAR(20) -- draft, submitted, approved
);
```

---

## computed_results (SYSTEM GENERATED)

```sql
CREATE TABLE computed_results (
    id UUID PRIMARY KEY,
    student_id UUID,
    subject_id UUID,
    total_score INT,
    grade VARCHAR(5),
    subject_position INT,
    class_position INT,
    session_id UUID,
    term_id UUID
);
```

---

## report_cards

```sql
CREATE TABLE report_cards (
    id UUID PRIMARY KEY,
    student_id UUID,
    session_id UUID,
    term_id UUID,
    file_url TEXT,
    result_id VARCHAR(50),
    is_locked BOOLEAN DEFAULT true
);
```

---

# 🧪 6. QUIZ SYSTEM

## quizzes

```sql
CREATE TABLE quizzes (
    id UUID PRIMARY KEY,
    title VARCHAR(255),
    class_id UUID,
    created_by UUID,
    duration_minutes INT,
    start_time TIMESTAMP,
    end_time TIMESTAMP
);
```

---

## quiz_questions

```sql
CREATE TABLE quiz_questions (
    id UUID PRIMARY KEY,
    quiz_id UUID,
    question TEXT,
    option_a TEXT,
    option_b TEXT,
    option_c TEXT,
    option_d TEXT,
    correct_answer VARCHAR(1)
);
```

---

## quiz_attempts

```sql
CREATE TABLE quiz_attempts (
    id UUID PRIMARY KEY,
    quiz_id UUID,
    student_id UUID,
    score INT,
    started_at TIMESTAMP,
    submitted_at TIMESTAMP,
    status VARCHAR(20)
);
```

---

# 💰 7. FINANCE SYSTEM

## fees

```sql
CREATE TABLE fees (
    id UUID PRIMARY KEY,
    class_id UUID,
    session_id UUID,
    amount DECIMAL
);
```

---

## payments

```sql
CREATE TABLE payments (
    id UUID PRIMARY KEY,
    student_id UUID,
    amount DECIMAL,
    method VARCHAR(50),
    receipt_no VARCHAR(50),
    paid_at TIMESTAMP
);
```

---

# 📢 8. ADMISSIONS SYSTEM

## applications

```sql
CREATE TABLE applications (
    id UUID PRIMARY KEY,
    application_number VARCHAR(50),
    full_name VARCHAR(255),
    desired_class VARCHAR(50),
    status VARCHAR(20),
    created_at TIMESTAMP
);
```

---

## admission_appointments

```sql
CREATE TABLE admission_appointments (
    id UUID PRIMARY KEY,
    application_id UUID,
    date DATE,
    time TIME,
    status VARCHAR(20)
);
```

---

# 🧾 9. AUDIT LOGS

```sql
CREATE TABLE audit_logs (
    id UUID PRIMARY KEY,
    user_id UUID,
    action VARCHAR(255),
    table_name VARCHAR(100),
    record_id UUID,
    before_state JSONB,
    after_state JSONB,
    ip_address VARCHAR(50),
    created_at TIMESTAMP
);
```

---

# 🧱 2. GO BACKEND STRUCTURE (PRODUCTION ARCHITECTURE)

```text
leaps-backend/
│
├── cmd/
│   └── server/main.go
│
├── config/
│   ├── db.go
│   ├── env.go
│
├── internal/
│   ├── models/
│   ├── controllers/
│   ├── services/
│   ├── repositories/
│   ├── middleware/
│   ├── routes/
│   ├── utils/
│   ├── auth/
│   ├── audit/
│   ├── admissions/
│   ├── academics/
│   ├── finance/
│   ├── quizzes/
│   ├── reports/
│   ├── users/
│   └── school/
│
├── pkg/
│   ├── logger/
│   ├── response/
│   ├── validator/
│
├── storage/
│
├── templates/
│   ├── reports/
│   ├── admission/
│
├── migrations/
├── go.mod
└── go.sum
```

---

# 🌐 API ENDPOINT STRUCTURE (REAL ROUTES)

## AUTH

```
POST   /api/auth/login
POST   /api/auth/register
POST   /api/auth/logout
GET    /api/auth/me
```

---

## STUDENTS

```
GET    /api/students
POST   /api/students
GET    /api/students/:id
PUT    /api/students/:id
```

---

## ADMISSIONS

```
POST   /api/admissions/apply
GET    /api/admissions
POST   /api/admissions/schedule
POST   /api/admissions/approve
POST   /api/admissions/reject
```

---

## ACADEMICS

```
POST   /api/scores/submit
GET    /api/results/student/:id
POST   /api/results/publish
POST   /api/results/lock
```

---

## QUIZZES

```
POST   /api/quizzes
POST   /api/quizzes/:id/start
POST   /api/quizzes/:id/submit
GET    /api/quizzes/:id/results
```

---

## FINANCE

```
POST   /api/fees
POST   /api/payments
GET    /api/payments/student/:id
```

---

## REPORTS

```
GET    /api/reports/student/:id
POST   /api/reports/generate
GET    /api/reports/download/:id
```

---

# 🎨 3. UI WIREFRAMES (ROLE DASHBOARDS)

---

# 🟣 OWNER DASHBOARD

```
+--------------------------------------------------+
| LEAPS ADMIN DASHBOARD                            |
+--------------------------------------------------+

[ Admissions ] [ Finance ] [ Academics ] [ Staff ]

[ SCHOOL OVERVIEW CARD ]
- Total Students
- Revenue
- Active Teachers
- Pending Admissions

[ SYSTEM SETTINGS ]
- Logo Upload
- Session Control
- Grading System

[ AUDIT LOGS PANEL ]
```

---

# 🔵 PRINCIPAL DASHBOARD

```
+---------------- SECONDARY CONTROL ----------------+

[ Results ] [ Quizzes ] [ Timetable ] [ Admissions ]

[ RESULT APPROVAL PANEL ]
- Pending Results
- Approve / Reject

[ QUIZ CONTROL PANEL ]
- Approve Questions
- Publish Quiz

[ STUDENT PERFORMANCE OVERVIEW ]
```

---

# 🟢 HEAD TEACHER DASHBOARD

```
+---------------- PRIMARY CONTROL ----------------+

[ Results ] [ Quizzes ] [ Classes ]

[ CLASS OVERVIEW ]
- Performance
- Attendance
- Weak Students

[ RESULT REVIEW PANEL ]
```

---

# 🟡 TEACHER DASHBOARD

```
+---------------- TEACHER PANEL ----------------+

[ My Classes ]
[ Score Entry ]
[ Attendance ]
[ Assignments ]

[ SCORE INPUT TABLE ]
Student | CA | Test | Exam
```

---

# 🟠 BURSAR DASHBOARD

```
+---------------- FINANCE ----------------+

[ Fees Setup ]
[ Payments ]
[ Receipts ]
[ Debtors ]

[ FINANCIAL SUMMARY ]
- Paid
- Pending
- Total Revenue
```

---

# 🧪 STUDENT DASHBOARD

```
+---------------- STUDENT PORTAL ----------------+

[ Results ]
[ Quizzes ]
[ Assignments ]
[ Attendance ]

[ PERFORMANCE CARD ]
```

---

# 👨‍👩‍👧 PARENT DASHBOARD

```
+---------------- PARENT PORTAL ----------------+

[ Child Results ]
[ Attendance ]
[ Fees ]
[ Admission Letter ]

[ DOWNLOAD REPORT CARD ]
```

---

# 🧱 LEAPS FRONTEND UI SYSTEM (STARTER KIT)

## 📁 PROJECT STRUCTURE

```text
frontend/
│
├── index.html
├── login.html
├── dashboard.html
│
├── assets/
│   ├── css/
│   │   ├── base.css
│   │   ├── dashboard.css
│   │   ├── auth.css
│   │   └── animations.css
│   │
│   ├── js/
│   │   ├── app.js
│   │   ├── api.js
│   │   ├── auth.js
│   │   ├── router.js
│   │   ├── dashboard.js
│   │   └── utils.js
│   │
│   └── images/
│       └── logo.png
│
└── pages/
    ├── owner/
    ├── principal/
    ├── teacher/
    ├── student/
    ├── parent/
```

---

# 🎨 1. BASE CSS SYSTEM (design foundation)

## assets/css/base.css

```css
:root {
  --primary: #1e3a8a;
  --secondary: #fbbf24;
  --bg: #f8fafc;
  --card: #ffffff;
  --text: #0f172a;
  --muted: #64748b;
  --radius: 12px;
}

* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  font-family: "Inter", sans-serif;
}

body {
  background: var(--bg);
  color: var(--text);
}

a {
  text-decoration: none;
  color: inherit;
}

button {
  cursor: pointer;
  border: none;
}
```

---

# ✨ 2. ANIMATIONS SYSTEM

## assets/css/animations.css

```css
.fade-in {
  animation: fadeIn 0.4s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.slide-in {
  animation: slideIn 0.3s ease;
}

@keyframes slideIn {
  from { transform: translateX(-20px); }
  to { transform: translateX(0); }
}
```

---

# 🔐 3. LOGIN PAGE

## login.html

```html
<!DOCTYPE html>
<html>
<head>
  <title>LEAPS Login</title>
  <link rel="stylesheet" href="assets/css/auth.css">
</head>

<body class="auth-body">

  <div class="login-card fade-in">
    <img src="assets/images/logo.png" class="logo"/>

    <h2>LEAPS Portal</h2>

    <input id="email" placeholder="Email"/>
    <input id="password" type="password" placeholder="Password"/>

    <button onclick="login()">Login</button>

    <p id="error"></p>
  </div>

  <script src="assets/js/auth.js"></script>
</body>
</html>
```

---

## auth.css

```css
.auth-body {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(120deg, #1e3a8a, #0f172a);
}

.login-card {
  background: white;
  padding: 30px;
  width: 350px;
  border-radius: 12px;
  text-align: center;
}

.login-card input {
  width: 100%;
  padding: 10px;
  margin: 8px 0;
}
```

---

## auth.js

```javascript
async function login() {
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;

  const res = await fetch("/api/auth/login", {
    method: "POST",
    headers: {"Content-Type":"application/json"},
    body: JSON.stringify({ email, password })
  });

  const data = await res.json();

  if (data.token) {
    localStorage.setItem("token", data.token);
    window.location.href = "dashboard.html";
  } else {
    document.getElementById("error").innerText = "Login failed";
  }
}
```

---

# 🧭 4. MAIN DASHBOARD SHELL

## dashboard.html

```html
<!DOCTYPE html>
<html>
<head>
  <title>LEAPS Dashboard</title>
  <link rel="stylesheet" href="assets/css/dashboard.css">
</head>

<body>

<div class="layout">

  <!-- SIDEBAR -->
  <div class="sidebar slide-in">
    <h2>LEAPS</h2>

    <a onclick="loadPage('overview')">Overview</a>
    <a onclick="loadPage('students')">Students</a>
    <a onclick="loadPage('admissions')">Admissions</a>
    <a onclick="loadPage('results')">Results</a>
    <a onclick="loadPage('finance')">Finance</a>
    <a onclick="loadPage('quizzes')">Quizzes</a>
    <a onclick="logout()">Logout</a>
  </div>

  <!-- MAIN CONTENT -->
  <div class="main">
    <div id="content" class="fade-in"></div>
  </div>

</div>

<script src="assets/js/api.js"></script>
<script src="assets/js/router.js"></script>
<script src="assets/js/dashboard.js"></script>

</body>
</html>
```

---

## dashboard.css

```css
.layout {
  display: flex;
}

.sidebar {
  width: 220px;
  background: var(--primary);
  color: white;
  height: 100vh;
  padding: 20px;
}

.sidebar a {
  display: block;
  padding: 10px;
  margin: 5px 0;
  color: white;
  border-radius: 8px;
}

.sidebar a:hover {
  background: rgba(255,255,255,0.1);
}

.main {
  flex: 1;
  padding: 20px;
}

.card {
  background: white;
  padding: 15px;
  border-radius: 10px;
  margin-bottom: 10px;
}
```

---

# 🌐 5. API LAYER (FRONTEND TO BACKEND)

## api.js

```javascript
const API = "http://localhost:8080/api";

function getToken() {
  return localStorage.getItem("token");
}

async function request(url, method = "GET", body = null) {
  const res = await fetch(API + url, {
    method,
    headers: {
      "Content-Type": "application/json",
      "Authorization": "Bearer " + getToken()
    },
    body: body ? JSON.stringify(body) : null
  });

  return await res.json();
}
```

---

# 🧭 6. SIMPLE ROUTER (NO FRAMEWORK)

## router.js

```javascript
async function loadPage(page) {
  const content = document.getElementById("content");

  switch(page) {

    case "overview":
      content.innerHTML = `
        <div class="card fade-in">
          <h2>School Overview</h2>
          <p>Welcome to LEAPS Dashboard</p>
        </div>`;
      break;

    case "students":
      const students = await request("/students");

      content.innerHTML = `
        <h2>Students</h2>
        ${students.map(s => `
          <div class="card">
            ${s.full_name}
          </div>
        `).join("")}
      `;
      break;

    case "admissions":
      content.innerHTML = `<h2>Admissions Panel</h2>`;
      break;

    case "results":
      content.innerHTML = `<h2>Results Engine</h2>`;
      break;

    case "finance":
      content.innerHTML = `<h2>Finance System</h2>`;
      break;

    case "quizzes":
      content.innerHTML = `<h2>Quiz System</h2>`;
      break;
  }
}
```

---

# 🧠 7. DASHBOARD LOGIC

## dashboard.js

```javascript
window.onload = () => {
  loadPage("overview");
};

function logout() {
  localStorage.removeItem("token");
  window.location.href = "login.html";
}
```

---

# 🧪 WHAT THIS FRONTEND GIVES YOU

## ✔ Fully working system shell

## ✔ Role-based UI ready structure

## ✔ API integration layer

## ✔ Dashboard navigation system

## ✔ Animation-ready UI

## ✔ Extensible for all modules

---

# 🚀 HOW THIS CONNECTS TO YOUR BACKEND

Your Go backend will expose:

* `/api/auth/login`
* `/api/students`
* `/api/admissions`
* `/api/results`
* `/api/quizzes`
* `/api/finance`

Frontend already expects these.

---


# 🧠 FINAL SYSTEM SUMMARY

This system is:

> A full institutional operating system for LEAPS that integrates admissions, academics, quizzes, finance, reporting, and media into a unified secure RBAC-controlled platform with automated computation and full auditability.

# END OF FULL SPECIFICATION

---
