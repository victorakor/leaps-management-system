const API_BASE = "https://leaps-management-system-production.up.railway.app/api";

function getToken() {
  return localStorage.getItem("token");
}

async function request(endpoint, method = "GET", body = null) {
  const headers = {
    "Content-Type": "application/json",
  };

  const token = getToken();
  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const options = {
    method,
    headers,
  };

  if (body) {
    options.body = JSON.stringify(body);
  }

  try {
    const response = await fetch(API_BASE + endpoint, options);
    const data = await response.json();

    if (!response.ok) {
      throw new Error(data.error || "Request failed");
    }

    return data;
  } catch (error) {
    console.error("API Error:", error);
    throw error;
  }
}

// Auth API
const AuthAPI = {
  login: (email, password) => request("/auth/login", "POST", { email, password }),
  register: (user) => request("/auth/register", "POST", user),
  getCurrentUser: () => request("/auth/me"),
};

// Students API
const StudentsAPI = {
  getAll: () => request("/students"),
  getById: (id) => request(`/students/${id}`),
  create: (student) => request("/students", "POST", student),
  update: (id, student) => request(`/students/${id}`, "PUT", student),
};

// Admissions API
const AdmissionsAPI = {
  apply: (application) => request("/admissions/apply", "POST", application),
  getAll: () => request("/admissions"),
  schedule: (appointment) => request("/admissions/schedule", "POST", appointment),
  approve: (id) => request(`/admissions/${id}/approve`, "POST"),
  reject: (id) => request(`/admissions/${id}/reject`, "POST"),
};

// Results API
const ResultsAPI = {
  submitScores: (scores) => request("/scores/submit", "POST", scores),
  getStudentResults: (studentId) => request(`/results/student/${studentId}`),
  publish: (resultId) => request("/results/publish", "POST", { result_id: resultId }),
  lock: (resultId) => request("/results/lock", "POST", { result_id: resultId }),
};

// Quizzes API
const QuizzesAPI = {
  create: (quiz) => request("/quizzes", "POST", quiz),
  start: (quizId) => request(`/quizzes/${quizId}/start`, "POST"),
  submit: (quizId, answers) => request(`/quizzes/${quizId}/submit`, "POST", answers),
  getResults: (quizId) => request(`/quizzes/${quizId}/results`),
};

// Finance API
const FinanceAPI = {
  createFee: (fee) => request("/fees", "POST", fee),
  recordPayment: (payment) => request("/payments", "POST", payment),
  getStudentPayments: (studentId) => request(`/payments/student/${studentId}`),
};

// Reports API
const ReportsAPI = {
  getStudentReport: (studentId) => request(`/reports/student/${studentId}`),
  generate: (reportData) => request("/reports/generate", "POST", reportData),
  download: (reportId) => request(`/reports/download/${reportId}`),
};