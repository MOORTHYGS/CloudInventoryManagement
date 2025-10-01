import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";
import { Chart, registerables } from "chart.js";
import "../../css files/common/common.css";
import { BACKEND_API_BASE_URL } from "../../../config";
import { Line, Bar, Pie } from 'react-chartjs-2';
import { Chart as ChartJS, CategoryScale, LinearScale, PointElement, LineElement, BarElement, Title, Tooltip, Legend, ArcElement } from 'chart.js';


import { Link } from "react-router-dom";

Chart.register(...registerables);

export function Homepage() {
  
ChartJS.register(CategoryScale, LinearScale, PointElement, LineElement, BarElement, Title, Tooltip, Legend, ArcElement);

const lineData = {
  labels: ["Jan", "Feb", "Mar", "Apr", "May", "Jun"],
  datasets: [
    {
      label: "Stock Level",
      data: [65, 59, 80, 81, 56, 55],
      borderColor: "#4fafc0", // Primary blue
      backgroundColor: "rgba(79, 175, 192, 0.2)",
      tension: 0.4,
    },
    {
      label: "Forecast",
      data: [60, 62, 78, 80, 55, 57],
      borderColor: "#7c3aed", // Secondary purple
      backgroundColor: "rgba(124, 58, 237, 0.2)",
      tension: 0.4,
    },
  ],
};

const barData = {
  labels: ["Warehouse A", "Warehouse B", "Warehouse C", "Warehouse D"],
  datasets: [
    {
      label: "Material Quantity",
      data: [120, 90, 70, 150],
      backgroundColor: ["#4fafc0", "#7c3aed", "#f59e0b", "#10b981"],
    },
  ],
};

const pieData = {
  labels: ["Material A", "Material B", "Material C", "Material D"],
  datasets: [
    {
      label: "Stock Share",
      data: [30, 25, 20, 25],
      backgroundColor: ["#4fafc0", "#7c3aed", "#f59e0b", "#ef4444"],
    },
  ],
};

const options = {
  responsive: true,
  plugins: {
    legend: {
      position: "bottom",
      labels: { color: "#e5e7eb" }, // Light gray for text
    },
    tooltip: {
      backgroundColor: "#111827",
      titleColor: "#f9fafb",
      bodyColor: "#f9fafb",
      borderColor: "#4fafc0",
      borderWidth: 1,
    },
  },
  scales: {
    x: { grid: { color: "#374151", borderColor: "#374151" }, ticks: { color: "#e5e7eb" } },
    y: { grid: { color: "#374151", borderColor: "#374151" }, ticks: { color: "#e5e7eb" } },
  },
};




  const [showLogin, setShowLogin] = useState(false); // popup modal
  const [loginForm, setLoginForm] = useState({
    username: "",
    password: ""
  });
  const [loginErrors, setLoginErrors] = useState({});


  const [currentIndex, setCurrentIndex] = useState(0);
  const [popup, setPopup] = React.useState({ show: false, message: "", type: "" });

  const tables = [
    ["Warehouse", "Material A", "Material B", "Material C"],
    ["Central DC - Chennai", "🔩 Steel Rods – 20%", "🪣 Cement Bags – 65%", "🎨 Paint Drums – 90%"],
    ["Regional Hub - Bangalore", "🚰 PVC Pipes – 55%", "🧱 Tiles – 30%", "🧱 Bricks – 70%"],
    ["Distribution Center - Hyderabad", "⚡ Copper Wire – 10%", "🪵 Wood Panels – 80%", "🪟 Glass Sheets – 50%"],
    ["North Zone - Delhi", "📄 Aluminium Sheets – 95%", "📌 Nails – 40%", "🔌 Cables – 25%"],
    ["West DC - Mumbai", "🧴 Plastic Granules – 60%", "🧷 Adhesives – 15%", "🛢 Lubricants – 85%"],
    ["South Hub - Coimbatore", "🔧 Iron Rods – 35%", "🏖 Sand – 90%", "🪨 Concrete Mix – 45%"],
    ["East Warehouse - Kolkata", "🧵 Fabric Rolls – 75%", "🧪 Dyes – 55%", "🧶 Thread Spools – 65%"],
  ];

  const colors = [
    "#3a1f68ff",
    "#4fafc0ff",
    "#eab65dff",
    "#055f41ff",
    "#7a1111ff",
    "#111e33ff",
    "#756d88ff",
    "#c973b9ff",
  ];

  // Auto carousel
  useEffect(() => {
    const interval = setInterval(() => {
      setCurrentIndex((prevIndex) => (prevIndex + 1) % tables.length);
    }, 3000);
    return () => clearInterval(interval);
  }, [tables.length]);


    function renderLoginErrors(name) {
    return loginErrors[name] ? (
      <div className="form-text text-danger">{loginErrors[name]}</div>
    ) : null;
  }

  function handleLoginChange(e) {
    const { name, value } = e.target;
    setLoginForm((s) => ({ ...s, [name]: value }));
  }

  function validateLogin() {
    const e = {};
    if (!loginForm.username.trim()) e.username = "User Name required";
    if (!loginForm.password) e.password = "Password required";
    setLoginErrors(e);
    return Object.keys(e).length === 0;
  }

async function submitLogin() {
  if (!validateLogin()) return;


  try {
    const response = await fetch(`${BACKEND_API_BASE_URL}/customers/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        user_name: loginForm.username,
        password: loginForm.password,
      }),
    });

    const data = await response.json();

    if (response.ok) {
      // Login successful

      setPopup({ show: true, message: `Welcome, ${loginForm.username}!`, type: "success" });
      setTimeout(() => {
        setShowLogin(false);
        setLoginForm({ username: "", password: "" });
        setPopup({ show: false, message: "", type: "" });
        window.location.href = `/dashboard/?data=${encodeURIComponent(JSON.stringify(data))}`;
      },1);
    } else {
      // Login failed
      setPopup({ show: true, message: data.error || "Login failed. Check credentials.", type: "error" });

      setTimeout(() => {
        setPopup({ show: false, message: "", type: "" });
      }, 2500);
    }
  } catch (error) {
    setPopup({ show: true, message: "Something went wrong. Try again later.", type: "error" });

    setTimeout(() => {
      setPopup({ show: false, message: "", type: "" });
    }, 2500);
  }
}


  return (
    <div>
      {/* Navbar */}
      <header>
        <nav className="navbar navbar-expand-lg navbar-dark nav container" >
          <a className="brand navbar-brand d-flex align-items-center" href="/">
            <div className="logo" aria-hidden="true">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none">
                <path d="M3 12l6-6 6 6 6-6" stroke="white" strokeOpacity="0.9" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round" />
              </svg>
            </div>
            <div>
              <div id="ctext" style={{ fontWeight: 800 }}>CloudInventory</div>
              <div style={{ fontSize: 12, color: "var(--muted)", marginTop: 2 }}>Smarter stock. Less guesswork.</div>
            </div>
          </a>
          <button className="navbar-toggler custom-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navMenu" aria-controls="navMenu" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
        <div className="collapse navbar-collapse" id="navMenu">
          <ul className="navbar-nav ms-auto nav-links">
            <li className="nav-item">
              <a className="nav-link" href="#features">Features</a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="#workflow">Workflow</a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="#roadmap">Roadmap</a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="#about">About</a>
            </li>
            <li className="nav-item">
              <a className="nav-link" href="#contact">Contact us</a>
            </li>
             <li className="nav-item">
              <Link to="/register" className="btn btn-primary me-2 mb-2" style={{ borderRadius: "25px", padding: "0.5rem 1.2rem" }}>
                Register
              </Link>
            </li>
            <li className="nav-item">
              <button className="btn btn-primary mb-2" style={{ borderRadius: "25px", padding: "0.5rem 1.2rem" }} onClick={() => setShowLogin(true)}>
                Login   
              </button>
            </li>
          </ul>
        </div>

        </nav>
      </header>

      {/* Hero */}
      <main className="container">
         {/* Hero */}
        <section className="hero" id="hero">
          <div className="row align-items-center">
            <div className="col-lg-7 col-md-12 mb-4">
              <h1 className="fade-up">Cloud Inventory Management Real-Time, Scalable, and Easy to Control</h1>
              <p className="lead">
                Monitor stock across warehouses, connect suppliers, automate reorders, and forecast demand — all from one powerful cloud dashboard.
              </p>

              <div className="kpis d-flex flex-wrap gap-4">
                <div className="kpi text-center"><strong>99.9%</strong><span>Uptime SLA</span></div>
                <div className="kpi text-center"><strong>40%</strong><span>Stockouts reduced</span></div>
                <div className="kpi text-center"><strong>Auto Reorder</strong><span>Smart thresholds</span></div>
                <div className="kpi text-center"><strong>Multi-Warehouse</strong><span>Central Control</span></div>
                <div className="kpi text-center"><strong>AI Forecast</strong><span>Demand Prediction</span></div>
                <div className="kpi text-center"><strong>ERP</strong><span>Seamless Integrations</span></div>
              </div>
            </div>

            <div className="col-lg-5 col-md-12">
              <div className="mock">
                <div className="screen">
                  <div className="topbar">
                    {tables.map((_, i) => (
                      <div
                        key={i}
                        className="dot"
                        style={{
                          background: i === currentIndex ? colors[i % colors.length] : "gray",
                          transform: i === currentIndex ? "scale(1.3)" : "scale(1)",
                          transition: "all 0.3s ease",
                        }}
                      ></div>
                    ))}
                  </div>
                  <div className="carousel">
                    {tables.map((table, index) => (
                      <div
                        className={`table ${index === currentIndex ? "active" : ""}`}
                        key={index}
                      >
                        {table.map((cellText, i) => (
                          <div className="cell" key={i}>{cellText}</div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </section>
        <div style={{ color: "var(--muted)", fontSize:12 }}>
              <p className="muted" style={{ margin: "8px 0 0" }}>“CloudInventory brings together warehouses, suppliers, and sales channels in one SaaS platform with microservices, real-time event streams, and secure role-based access designed to scale from SMBs to enterprises.”</p>
        </div>



        {/* Features + Chart */}
        <section className="hero" id="features">
          <div className="row align-items-center">
            <div className="col-lg-7 col-md-12 mb-4">
              <h1 >Real-time, Scalable Inventory Management</h1>
               <br /> 
              <p className="lead">Centralized data, automated reorders, forecasting and analytics — designed for multi-warehouse businesses.</p>
              <div className="mt-3 mb-2">
                <span className="lead">Real-time data,</span>
                <span className="lead"> Automated reorders,</span>
                <span className="lead"> Multi-warehouse</span>
              </div>
               <br /> 

              <div className="kpis d-flex flex-wrap">
                <div className="kpi "><span className="lead">Automated reordering with safety stock</span></div>
                <div className="kpi "><span className="lead">Custom reports, dashboards & scheduled exports</span></div>
                <div className="kpi "><span className="lead">Multi-channel order management (web, POS, marketplaces)</span></div>
                <div className="kpi "><span className="lead">Monthly forecast and recent sales history for a sample SKU. Adjust model inputs to reflect lead time & seasonality.</span></div>
              </div>
            </div>

            <div className="col-lg-5 col-md-12">
                <div className="col-md-7">
                  <div className="chart-container mb-4">
                    <Line data={lineData} options={options} />
                  </div>
                  <div className="chart-container mb-4">
                    <Bar data={barData} options={options} />
                  </div>
                  <div className="chart-container">
                    <Pie data={pieData} options={options} />
                  </div>
                </div>
            </div>

          </div>
        </section>



       <section id="workflow">
          <h2 style={{ marginBottom: 12 }}>Workflow — simple, observable</h2>
          <p className="muted" style={{ marginBottom: 12 }}>A clear flow from receiving to reorder — each step observable with automated actions and alerts.</p>

          <div className="workflow">
            <div className="step card">
              <div className="bubble">1</div>
              <h4>Receive</h4>
              <div className="small">Scan inbound stock, validate against purchase orders, record lot & expiry.</div>
            </div>
            <div className="step card">
              <div className="bubble">2</div>
              <h4>Store</h4>
              <div className="small">Smart bin placement with heatmaps to reduce picking time.</div>
            </div>
            <div className="step card">
              <div className="bubble">3</div>
              <h4>Pick & Pack</h4>
              <div className="small">Batch picks, mobile scanning and packing checklists.</div>
            </div>
            <div className="step card">
              <div className="bubble">4</div>
              <h4>Reorder</h4>
              <div className="small">Automatic purchase suggestions and supplier orchestration.</div>
            </div>
          </div>
        </section>


        <section id="roadmap" style={{ marginTop:25 }}>
          <h2 style={{ marginBottom: 12 }}>Roadmap</h2>
          <p className="muted" style={{ marginBottom: 18 }}>Where CloudInventory is headed — pragmatic milestones to expand reliability and integrations.</p>

          <div className="timeline">
            <div className="titem">
              <div className="dot"></div>
              <strong>Q3 2025 — Core v1</strong>
              <div className="small">Inventory tracking, multichannel sync, basic forecasting and mobile scanning.</div>
            </div>

            <div className="titem">
              <div className="dot"></div>
              <strong>Q4 2025 — Integrations</strong>
              <div className="small">ERP connectors, shipping APIs and supplier portal.</div>
            </div>

            <div className="titem">
              <div className="dot"></div>
              <strong>H1 2026 — Advanced analytics</strong>
              <div className="small">Demand signals, anomaly detection and automated reporting pipelines.</div>
            </div>

            <div className="titem">
              <div className="dot"></div>
              <strong>H2 2026 — Enterprise</strong>
              <div className="small">SAML SSO, multi-tenant policies and advanced SLA features.</div>
            </div>
          </div>
        </section>

        <section id="about" style={{ marginTop: 10 }}>
          <div className="container">
            <h2 className="mb-4 text-light">About</h2>
            <p className="text-secondary">
              CloudInventory was designed by operations and engineering teams who were tired of opaque stock levels and manual reconciliations. 
              Our mission: give teams one truth-of-stock, reduce waste and automate replenishment so businesses can focus on growth.
            </p>

            <div className="row g-3 mt-3">
              <div className="col-12 col-md-6">
                <div className="kpis d-flex flex-wrap gap-4">
                <div className="kpi text-center"><span>Encrypted at rest & transit, periodic pen-tests and role-based access.</span></div>
                <div className="kpi text-center"><span>Event-driven architecture, retries and near real-time replication across regions.</span></div>
              </div>
            </div>
          </div>
          </div>
        </section>


        <section id="contact" style={{ marginTop: 10 }}>
          <div className="container">
            <div className="row gy-4">
              {/* Company Info */}
              <div className="col-12 col-md-4">
                <h3 className="mb-3 text-light">CloudInventory</h3>
                <p className="text-secondary">
                  Empowering intelligent supply chain operations through real-time
                  visibility and predictive insights. Built for scale, driven by data.
                </p>
              </div>

              {/* Quick Links */}
              <div className="col-6 col-md-4">
                <h4 className="mb-3 text-light">Quick Links</h4>
                <ul className="list-unstyled text-secondary">
                  <li className="mb-2">
                    <a href="#about" className="text-secondary text-decoration-none">About</a>
                  </li>
                  <li className="mb-2">
                    <a href="#services" className="text-secondary text-decoration-none">Services</a>
                  </li>
                  <li className="mb-2">
                    <a href="#contact" className="text-secondary text-decoration-none">Contact</a>
                  </li>
                  <li className="mb-2">
                    <a href="#privacy" className="text-secondary text-decoration-none">Privacy Policy</a>
                  </li>
                </ul>
              </div>

              {/* Contact Info */}
              <div className="col-6 col-md-4">
                <h4 className="mb-3 text-light">Contact Us</h4>
                <p className="text-secondary mb-2">
                  📞 Phone: <a href="tel:8072655349" className="text-secondary text-decoration-none">8072655349</a>
                </p>
                <p className="text-secondary mb-2">
                  ✉️ Email: <a href="mailto:support@cloudinventory.com" className="text-secondary text-decoration-none">support@cloudinventory.com</a>
                </p>
                <p className="text-secondary">
                  📍 Address: Rayakottai, Hosur, TN
                </p>
              </div>

            </div>
          </div>
        </section>

       <footer >
        <div style={{borderTop: "1px dotted #334155",  marginTop: "0rem",  paddingTop: "0rem",  textAlign: "center",  fontSize: "0.9rem",  color: "#64748b"}} ></div>
        <div style={{ textAlign: "center", marginTop: "0.7rem", fontSize: "0.9rem", color: "#64748b", }}  >
          &copy; {new Date().getFullYear()} CloudInventory. All rights reserved.
        </div>
        <div  style={{    borderTop: "1px dotted #334155",    marginTop: "0.9rem",  paddingTop: "1.5rem",  textAlign: "center",  fontSize: "0.9rem",  color: "#64748b"  }}
        ></div>
      </footer>


      {showLogin && (
        <div
          className="modal d-block"
          tabIndex="-1"
          style={{
            backgroundColor: "rgba(3, 7, 33, 0.5)",
            display: "flex",
            alignItems: "center",
            justifyContent: "center",
            height: "100vh",
            marginTop:"150px"
          }}
        >
          <div className="modal-dialog">
            <div className="modal-content">
              <div className="modal-header">
                <h5 className="modal-title">Login</h5>
                <button type="button" className="btn-close" onClick={() => setShowLogin(false)}></button>
              </div>
              <div className="modal-body">
                <div className="mb-3">
                  <label className="form-label">User Name</label>
                  <input
                    type="text"
                    className="form-control"
                    name="username"
                    value={loginForm.customerId}
                    onChange={handleLoginChange}
                    placeholder="Enter User Name"
                  />
                  {renderLoginErrors("User Name")}
                </div>
                <div className="mb-3">
                  <label className="form-label">Password</label>
                  <input
                    type="password"
                    className="form-control"
                    name="password"
                    value={loginForm.password}
                    onChange={handleLoginChange}
                    placeholder="Enter password"
                  />
                  {renderLoginErrors("password")}
                </div>
              </div>
              <div className="modal-footer">
                <button className="btn btn-secondary" onClick={() => setShowLogin(false)}>Close</button>
                <button className="btn btn-primary" onClick={submitLogin}>Login</button>
              </div>
            </div>
          </div>
        </div>
      )}


      {popup.show && (
        <div
          style={{
            position: "fixed",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            padding: "20px 30px",
            backgroundColor: popup.type === "success" ? "#4caf50" : "#f44336",
            color: "#fff",
            borderRadius: "8px",
            boxShadow: "0 4px 12px rgba(0,0,0,0.3)",
            zIndex: 9999,
            minWidth: "250px",
            textAlign: "center",
            fontWeight: "bold",
            fontSize: "16px",
          }}
        >
          {popup.message}
        </div>
      )}



      </main>
    </div>
  );
}
