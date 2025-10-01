import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Homepage } from "./components/js components/default/home";
import Register from "./components/js components/default/Register";
import Dashboard from "./components/js components/customer/dashboard";

// 404 Page Component
function NotFound() {
  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>404</h1>
      <h2>Page Not Found</h2>
      <p>Oops! The page you are looking for doesn’t exist.</p>
      <a href="/" style={{ color: "blue", textDecoration: "underline" }}>
        Go back to Home
      </a>
    </div>
  );
}

function Home() {
  return <Homepage />;
}
function RegisterPage() {
  return <Register />;
}

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/dashboard" element={<Dashboard />} />

        {/* Catch-all for undefined routes */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </Router>
  );
}

export default App;
