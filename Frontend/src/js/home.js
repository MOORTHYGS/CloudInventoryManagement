import React, { useState } from "react";
import "../css/home.css"; // corrected import path

function googleLogin() {
  window.location.href = "https://cloudinventorymanagement.onrender.com/api/login/google";
}

function LoginModal({ closeModal, switchToSignup }) {
  return (
    <div id="modal" className="modal show">
      <div className="modal-content">
        <div className="modal-header">
          <h2 id="modal-title">Login</h2>
          <button id="close-modal" onClick={closeModal}>&times;</button>
        </div>
        <div className="modal-sub" id="modal-sub">
          Access your CloudInventory account
        </div>
        <form
          id="modal-form"
          onSubmit={e => {
            e.preventDefault();
            const data = { email: e.target.email.value };
            alert(JSON.stringify(data) + " Demo only. Integrate with your auth backend.");
            closeModal();
          }}
        >
          <div>
            <input type="email" name="email" placeholder="Email address" required />
            <button type="submit">Submit</button>
          </div>
          <div style={{ textAlign: "center", margin: "15px 0", position: "relative" }}>
            <span style={{ background: "#131111", padding: "0 5px", position: "relative", zIndex: 1 }}>or</span>
            <hr style={{ position: "absolute", top: "50%", left: 0, width: "100%", border: "none", borderTop: "1px solid #948585", zIndex: 0 }} />
          </div>
          <button
            type="button"
            id="google-login"
            onClick={googleLogin}
            style={{
              display: "flex",
              alignItems: "center",
              gap: "10px",
              justifyContent: "center",
              width: "100%",
              border: "1px solid #ccc",
              background: "#fff",
              padding: "8px",
              borderRadius: "5px",
              cursor: "pointer",
            }}
          >
            <img
              src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg"
              alt="Google"
              width={20}
              height={20}
            />
            <span>Sign in with Google</span>
          </button>
        </form>
        <div className="switch-sign" onClick={switchToSignup}>
          Create account
        </div>
      </div>
    </div>
  );
}

function SignupModal({ closeModal, switchToLogin }) {
  return (
    <div id="modal" className="modal show">
      <div className="modal-content">
        <div className="modal-header">
          <h2 id="modal-title">Create account</h2>
          <button id="close-modal" onClick={closeModal}>&times;</button>
        </div>
        <div className="modal-sub" id="modal-sub">
          Start your free trial — no credit card required
        </div>
        <form
          id="modal-form"
          onSubmit={e => {
            e.preventDefault();
            const data = { email: e.target.email.value };
            alert(JSON.stringify(data) + " Demo only. Integrate with your auth backend.");
            closeModal();
          }}
        >
          <div>
            <input type="email" name="email" placeholder="Email address" required />
            <button type="submit">Submit</button>
          </div>
          <div style={{ textAlign: "center", margin: "15px 0", position: "relative" }}>
            <span style={{ background: "#131111", padding: "0 5px", position: "relative", zIndex: 1 }}>or</span>
            <hr style={{ position: "absolute", top: "50%", left: 0, width: "100%", border: "none", borderTop: "1px solid #948585", zIndex: 0 }} />
          </div>
          <button
            type="button"
            id="google-login"
            onClick={googleLogin}
            style={{
              display: "flex",
              alignItems: "center",
              gap: "10px",
              justifyContent: "center",
              width: "100%",
              border: "1px solid #ccc",
              background: "#fff",
              padding: "8px",
              borderRadius: "5px",
              cursor: "pointer",
            }}
          >
            <img
              src="https://www.gstatic.com/firebasejs/ui/2.0.0/images/auth/google.svg"
              alt="Google"
              width={20}
              height={20}
            />
            <span>Sign up with Google</span>
          </button>
        </form>
        <div className="switch-sign" onClick={switchToLogin}>
          Back to login
        </div>
      </div>
    </div>
  );
}

function scrollToSection(id) {
  const el = document.getElementById(id);
  if (el) el.scrollIntoView({ behavior: "auto" });
}

export function Homepage() {
  // ✅ Hooks moved inside the component
  const [modalOpen, setModalOpen] = useState(false);
  const [modalKind, setModalKind] = useState("login");

  function openModal(kind) {
    setModalKind(kind);
    setModalOpen(true);
  }

  function closeModal() {
    setModalOpen(false);
  }

  return (
    <div>
      <header>
        <div className="nav container">
          <div className="brand">
            <div className="logo" aria-hidden="true">
              <svg width="20" height="20" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M3 12l6-6 6 6 6-6" stroke="white" strokeOpacity="0.9" strokeWidth="1.6" strokeLinecap="round" strokeLinejoin="round"/></svg>
            </div>
            <div>
              <div style={{ fontWeight: 800 }}>CloudInventory</div>
              <div style={{ fontSize: 12, color: "var(--muted)", marginTop: 2 }}>Smarter stock. Less guesswork.</div>
            </div>
          </div>
          <nav style={{ display: "flex", gap: 8, alignItems: "center" }}>
            <a href="#features">Features</a>
            <a href="#workflow">Workflow</a>
            <a href="#roadmap">Roadmap</a>
            <a href="#about">About</a>
            <div className="cta">
              <button className="btn secondary" onClick={() => openModal('login')}>Login</button>
              <button className="btn" onClick={() => openModal('signup')}>Sign up</button>
            </div>
          </nav>
        </div>
      </header>
      <main className="container">

        <section className="hero">
          <div>
            <h1 className="fade-up">Cloud inventory management built for speed, visibility and scale</h1>
            <p className="lead">Track stock across warehouses, unify supplier feeds, automate reordering and get predictive insights — all from a single cloud dashboard.</p>
            <div style={{ display: "flex", gap: 12, alignItems: "center" }}>
              <button className="btn" onClick={() => scrollToSection('features')}>Get started — it's free</button>
              <button className="btn secondary" onClick={() => alert('Demo coming soon')}>Watch demo</button>
            </div>

            <div className="kpis">
              <div className="kpi">
                <strong>99.9%</strong>
                <span className="small">Uptime SLA</span>
              </div>
              <div className="kpi">
                <strong>40%</strong>
                <span className="small">Avg. stockouts reduced</span>
              </div>
              <div className="kpi">
                <strong>Auto Reorder</strong>
                <span className="small">Rules, thresholds & schedules</span>
              </div>
            </div>

            <div style={{ marginTop: 28, color: "var(--muted)", fontSize: 13 }}>
              <strong>Details</strong>
              <p className="muted" style={{ margin: "8px 0 0" }}>CloudInventory is a SaaS-first inventory platform that unifies warehouse data, supplier feeds and sales channels. Built with microservices, real-time event pipelines and role-based access — ideal for SMBs to enterprise teams.</p>
            </div>
          </div>

          <aside className="mock fade-up" aria-hidden="true">
            <div className="screen">
              <div className="topbar"><div className="dot"></div><div className="dot" style={{ opacity: 0.8 }}></div><div className="dot" style={{ opacity: 0.6 }}></div></div>
              <div style={{ display: "flex", gap: 10, alignItems: "center", marginTop: 8 }}>
                <div style={{ flex: 1 }}>
                  <div style={{ display: "flex", gap: 8, alignItems: "center" }}>
                    <div style={{ width: 18, height: 18, borderRadius: 4, background: "linear-gradient(90deg,var(--accent-1),var(--accent-2))" }}></div>
                    <div style={{ fontSize: 12, color: "var(--muted)" }}>Warehouse — East</div>
                  </div>
                  <div style={{ fontWeight: 700, marginTop: 6 }}>In stock: 1,482</div>
                </div>
                <div style={{ textAlign: "right" }}>
                  <div style={{ fontSize: 11, color: "var(--muted)" }}>Backorders</div>
                  <div style={{ fontWeight: 700 }}>34</div>
                </div>
              </div>
              <div className="table">
                <div className="cell">SKU: 1001<br /><small className="small">Available</small></div>
                <div className="cell">SKU: 1023<br /><small className="small">Allocated</small></div>
                <div className="cell">SKU: 1108<br /><small className="small">Incoming</small></div>
                <div className="cell">SKU: 1320<br /><small className="small">Reserved</small></div>
              </div>
            </div>
          </aside>
        </section>

        <section id="features">
          <div style={{ display: "flex", alignItems: "center", justifyContent: "space-between", marginBottom: 18 }}>
            <h2>Features that remove friction</h2>
            <div className="muted">From essentials to advanced automation — built to scale.</div>
          </div>

          <div className="features grid">
            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>AI</div>
                <div>
                  <h3>Demand forecasting</h3>
                  <div className="small">Short-term and seasonal forecasts using sales history and trends.</div>
                </div>
              </div>
            </div>

            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>⟳</div>
                <div>
                  <h3>Auto reorder</h3>
                  <div className="small">Custom reorder rules, supplier lead time and safety stock calculations.</div>
                </div>
              </div>
            </div>

            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>🔗</div>
                <div>
                  <h3>Integrations</h3>
                  <div className="small">Connect ERP, e‑commerce, shipping carriers and barcode scanners.</div>
                </div>
              </div>
            </div>

            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>🔒</div>
                <div>
                  <h3>Role-based access</h3>
                  <div className="small">Fine-grained permissions and audit logs for compliance.</div>
                </div>
              </div>
            </div>

            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>📊</div>
                <div>
                  <h3>Live analytics</h3>
                  <div className="small">Turn raw events into dashboards, alerts and scheduled reports.</div>
                </div>
              </div>
            </div>

            <div className="card">
              <div style={{ display: "flex", alignItems: "center", gap: 12 }}>
                <div className="bubble" aria-hidden>⚡</div>
                <div>
                  <h3>Edge & offline</h3>
                  <div className="small">Local scanning clients sync to cloud for low-latency operations.</div>
                </div>
              </div>
            </div>

          </div>
        </section>

        <section id="workflow">
          <h2 style={{ marginBottom: 12 }}>Workflow — simple, observable</h2>
          <p className="muted" style={{ marginBottom: 18 }}>A clear flow from receiving to reorder — each step observable with automated actions and alerts.</p>

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

        <section id="roadmap">
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

        <section id="about">
          <h2>About</h2>
          <div className="about">
            <div>
              <p className="muted">CloudInventory was designed by operations and engineering teams who were tired of opaque stock levels and manual reconciliations. Our mission: give teams one truth-of-stock, reduce waste and automate replenishment so businesses can focus on growth.</p>

              <div style={{ display: "flex", gap: 12, marginTop: 18 }}>
                <div className="card" style={{ flex: 1 }}>
                  <h4>Secure by design</h4>
                  <div className="small">Encrypted at rest & transit, periodic pen-tests and role-based access.</div>
                </div>
                <div className="card" style={{ flex: 1 }}>
                  <h4>Reliable</h4>
                  <div className="small">Event-driven architecture, retries and near real-time replication across regions.</div>
                </div>
              </div>
            </div>

            <div style={{ textAlign: "right" }}>
              <div style={{ fontWeight: 700, fontSize: 20 }}>Trusted by operations teams</div>
              <div className="muted" style={{ marginTop: 10 }}>Start with a free trial, scale to enterprise.</div>
              <div style={{ marginTop: 20 }}><button className="btn" onClick={() => scrollToSection('open-signup')}>Start free trial</button></div>
            </div>
          </div>
        </section>

        <footer>
          <div style={{ display: "flex", justifyContent: "space-between", alignItems: "center", gap: 12, flexWrap: "wrap" }}>
            <div>
              <strong>CloudInventory</strong>
              <div className="muted">© 2025 CloudInventory Inc.</div>
            </div>
            <div className="muted">Contact: hello@cloudinventory.example · +91 98765 43210</div>
          </div>
        </footer>

      </main>
      {modalOpen && (
      modalKind === 'signup'
        ? <SignupModal closeModal={closeModal} switchToLogin={() => setModalKind('login')} />
        : <LoginModal closeModal={closeModal} switchToSignup={() => setModalKind('signup')} />
    )}

    </div>
  );
}
