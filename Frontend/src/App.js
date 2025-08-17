import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { Homepage } from "./js/home";
import { Dashboardpage } from "./js/dashboard";

// --- Your existing Home Page ---
function Home() {
  return <Homepage />;
}

// --- Dashboard Page ---
function Dashboard() {
  return <Dashboardpage />;
}

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/dashboard" element={<Dashboard />} />
      </Routes>
    </Router>
  );
}

export default App;
