async function loadPage(page) {
  const content = document.getElementById("content");
  const pageTitle = document.getElementById("pageTitle");

  // Update active nav item
  document.querySelectorAll(".nav-item").forEach((item) => {
    item.classList.remove("active");
  });
  event.target.closest(".nav-item").classList.add("active");

  try {
    switch (page) {
      case "overview":
        pageTitle.textContent = "Dashboard Overview";
        content.innerHTML = await getOverviewPage();
        break;

      case "students":
        pageTitle.textContent = "Students";
        content.innerHTML = await getStudentsPage();
        break;

      case "admissions":
        pageTitle.textContent = "Admissions";
        content.innerHTML = await getAdmissionsPage();
        break;

      case "results":
        pageTitle.textContent = "Academic Results";
        content.innerHTML = await getResultsPage();
        break;

      case "quizzes":
        pageTitle.textContent = "Quizzes & CBT";
        content.innerHTML = await getQuizzesPage();
        break;

      case "finance":
        pageTitle.textContent = "Finance Management";
        content.innerHTML = await getFinancePage();
        break;

      case "reports":
        pageTitle.textContent = "Reports";
        content.innerHTML = await getReportsPage();
        break;

      default:
        content.innerHTML = "<p>Page not found</p>";
    }
  } catch (error) {
    content.innerHTML = `<div class="card"><p class="error-message">Error loading page: ${error.message}</p></div>`;
  }
}

async function getOverviewPage() {
  return `
    <div class="grid">
      <div class="stat-card hover-glow">
        <div class="stat-icon">👥</div>
        <div class="stat-label">Total Students</div>
        <div class="stat-value">1,245</div>
      </div>
      <div class="stat-card hover-glow">
        <div class="stat-icon">💰</div>
        <div class="stat-label">Revenue</div>
        <div class="stat-value">₦2.5M</div>
      </div>
      <div class="stat-card hover-glow">
        <div class="stat-icon">👨‍🏫</div>
        <div class="stat-label">Active Teachers</div>
        <div class="stat-value">85</div>
      </div>
      <div class="stat-card hover-glow">
        <div class="stat-icon">📝</div>
        <div class="stat-label">Pending Admissions</div>
        <div class="stat-value">23</div>
      </div>
    </div>

    <div class="card mt-20">
      <h3>Recent Activities</h3>
      <table>
        <thead>
          <tr>
            <th>Activity</th>
            <th>User</th>
            <th>Date</th>
            <th>Status</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Result Published</td>
            <td>Mr. Okafor</td>
            <td>2026-06-23</td>
            <td><span style="color: var(--success);">✓ Completed</span></td>
          </tr>
          <tr>
            <td>Student Admitted</td>
            <td>Admin</td>
            <td>2026-06-22</td>
            <td><span style="color: var(--success);">✓ Completed</span></td>
          </tr>
          <tr>
            <td>Payment Recorded</td>
            <td>Bursar</td>
            <td>2026-06-21</td>
            <td><span style="color: var(--success);">✓ Completed</span></td>
          </tr>
        </tbody>
      </table>
    </div>
  `;
}

async function getStudentsPage() {
  return `
    <div class="card">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h3>Student Directory</h3>
        <button class="btn btn-primary">+ Add Student</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>Admission No.</th>
            <th>Full Name</th>
            <th>Class</th>
            <th>Status</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>SCH/2026/PRI/0001</td>
            <td>Chioma Okafor</td>
            <td>Primary 5</td>
            <td><span style="color: var(--success);">Active</span></td>
            <td><button class="btn btn-primary" style="padding: 5px 10px; font-size: 12px;">View</button></td>
          </tr>
          <tr>
            <td>SCH/2026/SEC/0045</td>
            <td>Tunde Adeyemi</td>
            <td>JSS 2</td>
            <td><span style="color: var(--success);">Active</span></td>
            <td><button class="btn btn-primary" style="padding: 5px 10px; font-size: 12px;">View</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  `;
}

async function getAdmissionsPage() {
  return `
    <div class="card">
      <h3>Admission Applications</h3>
      <div style="margin-top: 20px;">
        <div class="card" style="background: var(--bg);">
          <div style="display: flex; justify-content: space-between;">
            <div>
              <p class="text-muted">Application #APP/2026/000145</p>
              <p><strong>Amara Nwosu</strong></p>
              <p class="text-muted">Desired Class: Primary 4</p>
            </div>
            <div style="text-align: right;">
              <p style="color: var(--warning); font-weight: bold;">PENDING</p>
              <button class="btn btn-primary" style="padding: 8px 15px; font-size: 12px; margin-top: 10px;">Review</button>
            </div>
          </div>
        </div>
      </div>
    </div>
  `;
}

async function getResultsPage() {
  return `
    <div class="card">
      <h3>Academic Results Management</h3>
      <div style="margin-top: 20px;">
        <div class="form-group">
          <label>Select Class</label>
          <select>
            <option>Primary 5</option>
            <option>Primary 6</option>
            <option>JSS 1</option>
            <option>JSS 2</option>
            <option>JSS 3</option>
            <option>SSS 1</option>
            <option>SSS 2</option>
            <option>SSS 3</option>
          </select>
        </div>
        <button class="btn btn-primary">Load Results</button>
      </div>
    </div>
  `;
}

async function getQuizzesPage() {
  return `
    <div class="card">
      <div style="display: flex; justify-content: space-between; align-items: center; margin-bottom: 20px;">
        <h3>Quiz Management</h3>
        <button class="btn btn-primary">+ Create Quiz</button>
      </div>
      <table>
        <thead>
          <tr>
            <th>Quiz Title</th>
            <th>Class</th>
            <th>Duration</th>
            <th>Status</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Mathematics Quiz 1</td>
            <td>JSS 2</td>
            <td>30 mins</td>
            <td><span style="color: var(--success);">Active</span></td>
            <td><button class="btn btn-primary" style="padding: 5px 10px; font-size: 12px;">Manage</button></td>
          </tr>
        </tbody>
      </table>
    </div>
  `;
}

async function getFinancePage() {
  return `
    <div class="grid">
      <div class="stat-card hover-glow">
        <div class="stat-icon">💵</div>
        <div class="stat-label">Total Paid</div>
        <div class="stat-value">₦1.8M</div>
      </div>
      <div class="stat-card hover-glow">
        <div class="stat-icon">⏳</div>
        <div class="stat-label">Pending</div>
        <div class="stat-value">₦700K</div>
      </div>
      <div class="stat-card hover-glow">
        <div class="stat-icon">⚠️</div>
        <div class="stat-label">Debtors</div>
        <div class="stat-value">45</div>
      </div>
    </div>

    <div class="card mt-20">
      <h3>Recent Payments</h3>
      <table>
        <thead>
          <tr>
            <th>Student</th>
            <th>Amount</th>
            <th>Method</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>Chioma Okafor</td>
            <td>₦50,000</td>
            <td>Bank Transfer</td>
            <td>2026-06-23</td>
          </tr>
        </tbody>
      </table>
    </div>
  `;
}

async function getReportsPage() {
  return `
    <div class="card">
      <h3>Generate Reports</h3>
      <div style="margin-top: 20px;">
        <div class="form-group">
          <label>Report Type</label>
          <select>
            <option>Student Report Card</option>
            <option>Class Performance</option>
            <option>Financial Report</option>
            <option>Attendance Report</option>
          </select>
        </div>
        <button class="btn btn-primary">Generate Report</button>
      </div>
    </div>
  `;
}
