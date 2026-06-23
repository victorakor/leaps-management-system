window.addEventListener("load", async () => {
  try {
    const user = await AuthAPI.getCurrentUser();
    if (user.success) {
      const fullName = user.data.full_name || "User";
      document.getElementById("userName").textContent = fullName;
      // Set initials in avatar
      const initials = fullName.split(" ").map(n => n[0]).join("").toUpperCase().slice(0, 2);
      document.getElementById("userAvatar").textContent = initials;
    }
  } catch (error) {
    console.log("User not authenticated, redirecting to login");
    window.location.href = "login.html";
  }

  // Load default page — pass null since there's no click event on initial load
  loadPage("overview", null);
});

function logout() {
  if (confirm("Are you sure you want to logout?")) {
    localStorage.removeItem("token");
    localStorage.removeItem("user");
    window.location.href = "login.html";
  }
}