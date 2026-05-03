document.getElementById("signupForm").addEventListener("submit", async (e) => {
  e.preventDefault();

  const name = document.getElementById("name").value.trim();
  const email = document.getElementById("email").value.trim();
  const password = document.getElementById("password").value;
  const confirmPassword = document.getElementById("confirmPassword").value;

  // 🔐 Basic validation
  if (password !== confirmPassword) {
    alert("Passwords do not match ❌");
    return;
  }

  if (password.length < 6) {
    alert("Password must be at least 6 characters");
    return;
  }

  try {
    const res = await fetch("http://localhost:8080/register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        name,
        email,
        password
      })
    });

    const data = await res.json();

    if (res.ok) {
      alert("Account created 🎉");

      // Redirect to login
      window.location.href = "login.html";
    } else {
      alert(data.message || "Signup failed");
    }

  } catch (err) {
    console.error(err);
    alert("Server error");
  }
});