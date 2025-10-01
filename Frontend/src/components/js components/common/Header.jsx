import React from "react";
import "../../css files/default/Header.css";

const Header = () => {
  return (
        <header>
        <nav className="navbar navbar-expand-lg navbar-dark nav container">
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
                <a className="nav-link" href="#features">Dashboard</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#workflow">Inventiory</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#roadmap">Orders</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#about">Suppliers</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#contact">Analysis</a>
              </li>
            </ul>

            <ul className="navbar-nav ms-auto nav-links">
            <li className="nav-item">
                <a className="nav-link d-flex align-items-center" href="/notifications">
                <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" fill="currentColor" className="bi bi-bell me-1" viewBox="0 0 16 16" >
                    <path d="M8 16a2 2 0 0 0 1.985-1.75H6.015A2 2 0 0 0 8 16zm.104-14.804a1 1 0 0 1 1.792.934A5.002 5.002 0 0 1 13 7c0 1.098.37 2.17 1.071 3H1.929C2.63 9.17 3 8.098 3 7a5.002 5.002 0 0 1 3.104-4.87z"/>
                </svg>
                Notifications
                </a>
            </li>
            <li className="nav-item dropdown">
            <a
                className="nav-link d-flex align-items-center dropdown-toggle"
                href="#q"
                role="button"
                id="accountDropdown"
                data-bs-toggle="dropdown"
                aria-expanded="false"
            >
                <svg
                xmlns="http://www.w3.org/2000/svg"
                width="20"
                height="20"
                fill="currentColor"
                className="bi bi-person-circle me-1"
                viewBox="0 0 16 16"
                >
                <path d="M13 8a5 5 0 1 1-10 0 5 5 0 0 1 10 0z"/>
                <path fillRule="evenodd" d="M8 15A7 7 0 1 0 8 1a7 7 0 0 0 0 14zm0-6a3 3 0 1 1 0-6 3 3 0 0 1 0 6z"/>
                </svg>
                Account
            </a>
            <ul className="dropdown-menu dropdown-menu-end" aria-labelledby="accountDropdown">
                <li>
                  <a className="dropdown-item d-flex align-items-center" href="/profile">
                    <i className="fas fa-user me-2"></i> Profile
                  </a>
                </li>
                <li>
                  <a className="dropdown-item d-flex align-items-center" href="/settings">
                    <i className="fas fa-cog me-2"></i> Settings
                  </a>
                </li>
                <li><hr className="dropdown-divider" /></li>
                <li>
                  <button className="dropdown-item d-flex align-items-center" href="/logout">
                    <i className="fas fa-sign-out-alt me-2"></i> Logout
                  </button>
                </li>
            </ul>
            </li>
            </ul>
          </div>
        </nav>
      </header>
  );
};

export default Header;
