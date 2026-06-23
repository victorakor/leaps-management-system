document.getElementById("loginForm").addEventListener("submit", async (e) => {
  e.preventDefault();

  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;
  const errorEl = document.getElementById("error");

  try {
    errorEl.textContent = "";
    const response = await AuthAPI.login(email, password);

    if (response.success) {
      localStorage.setItem("token", response.data.token);
      localStorage.setItem("user", JSON.stringify(response.data.user));
      window.location.href = "dashboard.html";
    } else {
      errorEl.textContent = response.error || "Login failed";
    }
  } catch (error) {
    errorEl.textContent = error.message || "Login failed";
  }
});

function logout() {
  localStorage.removeItem("token");
  localStorage.removeItem("user");
  window.location.href = "login.html";
}
