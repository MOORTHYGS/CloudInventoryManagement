import { useEffect, useState } from "react";

export function Dashboardpage() {
  const [message, setMessage] = useState("Loading...");

  useEffect(() => {
    const hash = window.location.hash.substring(1);
    const params = new URLSearchParams(hash);

    const accessToken = params.get("access_token");

    if (accessToken) {
      // Send token to backend -> set as secure HTTP-only cookie
      fetch("https://cloudinventorymanagement.onrender.com/api/set-token", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ access_token: accessToken }),
        credentials: "include", // <-- ensures cookies are stored
      }).then(() => {
        // Clean the URL (remove #access_token)
        window.history.replaceState({}, document.title, "/dashboard");

        // Fetch protected resource using cookie
        fetch("https://cloudinventorymanagement.onrender.com/api/dashboard", {
          credentials: "include",
        })
          .then((res) => res.json())
          .then((data) => setMessage(data.message)) // update React state
          .catch(() => setMessage("Failed to load dashboard data"));
      });
    } else {
      setMessage("No access token found. Redirecting...");
      window.location.href = "/";
    }
  }, []);

  return (
    <div>
      <h1>Dashboard</h1>
      <p>{message}</p>
    </div>
  );
}
