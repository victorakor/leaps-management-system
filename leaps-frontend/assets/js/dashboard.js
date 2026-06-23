window.addEventListener("load", async () => {
  try {
    const user = await AuthAPI.getCurrentUser();
    if (user.success) {
      document.getElementById("userName").textContent = user.data.full_name || "User";
    }
  } catch (error) {
    console.log("User not authenticated, redirecting to login");
    window.location.href = "login.html";
  }

  // Load default page
  loadPage("overview");
});

function logout() {
  if (confirm("Are you sure you want to logout?")) {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    window.location.href = "login.html";
  }
}
