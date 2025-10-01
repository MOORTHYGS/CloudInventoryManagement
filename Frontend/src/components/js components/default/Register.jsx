import React, { useState } from "react";
import "../../css files/default/Register.css";

export default function Register() {
  const STEPS = ["Company", "Inventory type", "Setup", "Payment"];
  const total = STEPS.length;

  const [step, setStep] = useState(1);
  const [loading, setLoading] = useState(false);
  const [paid, setPaid] = useState(false);
  const [fileName, setFileName] = useState(null);
  const [showLogin, setShowLogin] = useState(false); 

  const [form, setForm] = useState({
    companyName: "",
    email: "",
    phone: "",
    regId: "",
    address: "",
    city: "",
    state: "",
    country: "",
    website: "",
    inventoryType: "Enterprise",
    inventoryTypeOther: "",
    approxSKUs: "1-500",
    warehouses: "1",
    users: "5",
    contactName: "",
    contactRole: "",
    features: [],
    integrations: [],
    migrationNeeded: false,
    timeline: "",
    amount: 1
  });

  const [errors, setErrors] = useState({});
  
  const [loginForm, setLoginForm] = useState({
    tenantId: "",
    password: ""
  });
  const [loginErrors, setLoginErrors] = useState({});

  const pct = Math.round(((step - 1) / (total - 1)) * 100);

  function updateField(e) {
    const { name, value, type, checked, files } = e.target;

    if (type === "checkbox" && name === "migrationNeeded") {
      setForm((s) => ({ ...s, migrationNeeded: checked }));
      if (!checked) setFileName(null);
      return;
    }

    if (type === "file") {
      const f = files && files[0];
      setFileName(f ? f.name : null);
      return;
    }

    if (type === "checkbox") {
      if (name === "features") {
        setForm((s) => {
          const has = s.features.includes(value);
          return {
            ...s,
            features: has
              ? s.features.filter((x) => x !== value)
              : [...s.features, value],
          };
        });
        return;
      }

      if (name === "integrations") {
        setForm((s) => {
          const has = s.integrations.includes(value);
          return {
            ...s,
            integrations: has
              ? s.integrations.filter((x) => x !== value)
              : [...s.integrations, value],
          };
        });
        return;
      }
    }

    setForm((s) => ({ ...s, [name]: value }));
  }

  function handleLoginChange(e) {
    const { name, value } = e.target;
    setLoginForm((s) => ({ ...s, [name]: value }));
  }

  function validateLogin() {
    const e = {};
    if (!loginForm.tenantId.trim()) e.tenantId = "Tenant ID required";
    if (!loginForm.password) e.password = "Password required";
    setLoginErrors(e);
    return Object.keys(e).length === 0;
  }

  function submitLogin() {
    if (!validateLogin()) return;

    console.log("Logging in with:", loginForm);
    // TODO: call your backend login API here
    alert(`Tenant ${loginForm.tenantId} logged in!`);
    setShowLogin(false);
    setLoginForm({ tenantId: "", password: "" });
  }

  function renderErrors(name) {
    return errors[name] ? (
      <div className="form-text text-danger">{errors[name]}</div>
    ) : null;
  }

  function renderLoginErrors(name) {
    return loginErrors[name] ? (
      <div className="form-text text-danger">{loginErrors[name]}</div>
    ) : null;
  }

  function validateStep() {
    const e = {};

    if (step === 1) {
      if (!form.companyName.trim())
        e.companyName = "Company name is required";
      if (!/^\S+@\S+\.\S+$/.test(form.email))
        e.email = "Valid email required";
      if (!/^\+?[0-9()\-\s]{7,20}$/.test(form.phone))
        e.phone = "Valid phone required";
      if (!form.regId.trim())
        e.regId = "Registration / Tax ID is required";
    }

    if (step === 2) {
      if (!form.inventoryType)
        e.inventoryType = "Please select an inventory type";
      if (form.inventoryType === "Other" && !form.inventoryTypeOther.trim())
        e.inventoryTypeOther = "Please tell us the other inventory type";
    }

    if (step === 3) {
      if (!form.contactName.trim())
        e.contactName = "Primary contact name is required";
      if (!form.contactRole.trim())
        e.contactRole = "Primary contact role is required";
    }

    setErrors(e);
    return Object.keys(e).length === 0;
  }

  function next() {
    if (!validateStep()) {
      window.scrollTo({ top: 0, behavior: "smooth" });
      return;
    }
    setStep((s) => Math.min(total, s + 1));
  }

  function prev() {
    setStep((s) => Math.max(1, s - 1));
  }

  async function processPayment() {
    if (!validateStep()) return;

    setLoading(true);
    try {
      const res = await fetch("http://localhost:8080/register/payment", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });
      const order = await res.json();
      const options = {
        key: "rzp_test_REc8Mx2AfL87K5",
        amount: order.amount,
        currency: order.currency,
        name: form.companyName,
        description: "CloudInventory Registration Payment",
        order_id: order.id,
        handler: function (response) {
          console.log("✅ Payment Success:", response);
          setPaid(true);
        },
        prefill: {
          name: form.contactName,
          email: form.email,
          contact: form.phone,
        },
        theme: { color: "#0d6efd" },
      };
      const rzp = new window.Razorpay(options);
      rzp.open();
    } catch (err) {
      console.error("Payment error:", err);
      alert("Something went wrong with payment. Check console.");
    } finally {
      setLoading(false);
    }
  }

  return (
    <div>
      {/* Navbar with Login Popup Trigger */}
      <header>
        <nav className="navbar navbar-expand-lg navbar-dark nav container">
          <a className="brand navbar-brand d-flex align-items-center" href="/">
            <div className="logo" aria-hidden="true">
              <svg width="20px" height="20px" viewBox="0 0 24 24" fill="none">
                <path
                  d="M3 12l6-6 6 6 6-6"
                  stroke="white"
                  strokeOpacity="0.9"
                  strokeWidth="1.6"
                  strokeLinecap="round"
                  strokeLinejoin="round"
                />
              </svg>
            </div>
            <div>
              <div id="ctext" style={{ fontWeight: 800 }}>CloudInventory</div>
              <div
                style={{
                  fontSize: 12,
                  color: "var(--muted)",
                  marginTop: 2,
                }}
              >
                Smarter stock. Less guesswork.
              </div>
            </div>
          </a>
          <button
            className="navbar-toggler custom-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navMenu"
            aria-controls="navMenu"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navMenu">
            <ul className="navbar-nav ms-auto nav-links">
              <li className="nav-item"><a className="nav-link" href="#features">Features</a></li>
              <li className="nav-item"><a className="nav-link" href="#workflow">Workflow</a></li>
              <li className="nav-item"><a className="nav-link" href="#roadmap">Roadmap</a></li>
              <li className="nav-item"><a className="nav-link" href="#about">About</a></li>
              <li className="nav-item"><a className="nav-link" href="#contact">Contact us</a></li>
              <li className="btn btn-primary nav-link text-white"onClick={() => setShowLogin(true)}  style={{marginLeft: "1rem",borderRadius: "25px",padding: "0.5rem 1.2rem",}} >Login</li>
            </ul>
          </div>
        </nav>
      </header>

     {/* Main content */}
      <section style={{ flex: 1 }} className="container hero">
        <div className="card shadow-sm hero1">
          <div className="card-body p-4">
            <div className="d-flex justify-content-between align-items-center mb-3">
              <div>
                <h5 className="mb-0">CloudInventory — Company registration</h5>
              </div>
              <div className="text-end">
                <small className="text">Step {step} of {total}</small>
              </div>
            </div>

            <div className="progress mb-4" style={{ height: 10, borderRadius: 6 }}>
              <div
                className="progress-bar"
                role="progressbar"
                style={{ width: `${pct}%` }}
                aria-valuenow={pct}
                aria-valuemin="0"
                aria-valuemax="100"
              />
            </div>

            {paid ? (
              <div className="alert alert-success">
                <h6 className="mb-1">Demo payment successful</h6>
                <div className="small text-muted">
                  Thank you — we created a demo account for{" "}
                  <strong>{form.companyName || "your company"}</strong>.  
                  Our team will contact {form.contactName || "you"} with next steps.
                </div>
              </div>
            ) : null}

            {/* Step 1: Company details */}
            {step === 1 && (
              <div>
                <h6 className="mb-3">Company details</h6>
                <div className="row g-3">
                  <div className="col-md-6">
                    <label className="form-label">Company name *</label>
                    <input name="companyName" value={form.companyName} onChange={updateField} className="form-control" placeholder="Acme Foods Pvt Ltd" />
                    {renderErrors("companyName")}
                  </div>
                  <div className="col-md-6">
                    <label className="form-label">Official email *</label>
                    <input name="email" value={form.email} onChange={updateField} className="form-control" placeholder="hello@acme.com" />
                    {renderErrors("email")}
                  </div>

                  <div className="col-md-4">
                    <label className="form-label">Phone *</label>
                    <input name="phone" value={form.phone} onChange={updateField} className="form-control" placeholder="+91 90000 00000" />
                    {renderErrors("phone")}
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">Registered ID (GST / Tax ID) *</label>
                    <input name="regId" value={form.regId} onChange={updateField} className="form-control" placeholder="29ABCDE1234F2Z5" />
                    {renderErrors("regId")}
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">Website (optional)</label>
                    <input name="website" value={form.website} onChange={updateField} className="form-control" placeholder="https://your-company.com" />
                  </div>

                  <div className="col-md-12">
                    <label className="form-label">Address</label>
                    <input name="address" value={form.address} onChange={updateField} className="form-control" placeholder="Street, Building, Suite" />
                  </div>

                  <div className="col-md-4">
                    <label className="form-label">City</label>
                    <input name="city" value={form.city} onChange={updateField} className="form-control" />
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">State / Province</label>
                    <input name="state" value={form.state} onChange={updateField} className="form-control" />
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">Country</label>
                    <input name="country" value={form.country} onChange={updateField} className="form-control" placeholder="India" />
                  </div>
                </div>
              </div>
            )}

            {/* Step 2: Inventory type */}
            {step === 2 && (
              <div>
                <h6 className="mb-3">Inventory profile</h6>
                <div className="mb-3">
                  <label className="form-label d-block">Which category best describes your business? *</label>
                  {["Enterprise", "Supermarket", "Warehouse / Distribution", "Manufacturing", "Retail", "eCommerce", "Other"].map((t) => (
                    <div className="form-check form-check-inline" key={t}>
                      <input className="form-check-input" type="radio" id={`type-${t}`} name="inventoryType" value={t} checked={form.inventoryType === t} onChange={updateField} />
                      <label className="form-check-label" htmlFor={`type-${t}`}>{t}</label>
                    </div>
                  ))}
                  {renderErrors("inventoryType")}

                  {form.inventoryType === "Other" && (
                    <div className="mt-3">
                      <input name="inventoryTypeOther" value={form.inventoryTypeOther} onChange={updateField} className="form-control" placeholder="Please describe" />
                      {renderErrors("inventoryTypeOther")}
                    </div>
                  )}
                </div>

                <div className="row g-3">
                  <div className="col-md-4">
                    <label className="form-label">Approx. SKUs</label>
                    <select name="approxSKUs" value={form.approxSKUs} onChange={updateField} className="form-select">
                      <option>1-500</option>
                      <option>500-5,000</option>
                      <option>5,000-50,000</option>
                      <option>50,000+</option>
                    </select>
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">Warehouses / Locations</label>
                    <input name="warehouses" value={form.warehouses} onChange={updateField} type="number" min="1" className="form-control" />
                  </div>
                  <div className="col-md-4">
                    <label className="form-label">Active users (estimate)</label>
                    <input name="users" value={form.users} onChange={updateField} type="number" min="1" className="form-control" />
                  </div>
                </div>
              </div>
            )}

            {/* Step 3: Setup details */}
            {step === 3 && (
              <div>
                <h6 className="mb-3">Setup & preferences</h6>

                <div className="row g-3">
                  <div className="col-md-6">
                    <label className="form-label">Primary contact name *</label>
                    <input name="contactName" value={form.contactName} onChange={updateField} className="form-control" placeholder="Ravi Kumar" />
                    {renderErrors("contactName")}
                  </div>
                  <div className="col-md-6">
                    <label className="form-label">Role / Title *</label>
                    <input name="contactRole" value={form.contactRole} onChange={updateField} className="form-control" placeholder="Operations Manager" />
                    {renderErrors("contactRole")}
                  </div>

                  <div className="col-12">
                    <label className="form-label">Key features you'd like</label>
                    <div className="d-flex flex-wrap gap-2">
                      {["Reorder automation", "Forecasting & demand planning", "Barcode / RFID", "Multi-channel sales", "Analytics dashboard", "Custom integrations"].map((f) => (
                        <div className="form-check" style={{ minWidth: 220 }} key={f}>
                          <input className="form-check-input" type="checkbox" id={`f-${f}`} name="features" value={f} checked={form.features.includes(f)} onChange={updateField} />
                          <label className="form-check-label" htmlFor={`f-${f}`}>{f}</label>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div className="col-12">
                    <label className="form-label">Integrations (select those you need)</label>
                    <div className="d-flex flex-wrap gap-2">
                      {["SAP", "QuickBooks / Xero", "Shopify / WooCommerce", "Tally / ERP", "Custom API"].map((i) => (
                        <div className="form-check" style={{ minWidth: 180 }} key={i}>
                          <input className="form-check-input" type="checkbox" id={`i-${i}`} name="integrations" value={i} checked={form.integrations.includes(i)} onChange={updateField} />
                          <label className="form-check-label" htmlFor={`i-${i}`}>{i}</label>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div className="col-md-6">
                    <div className="form-check form-switch mt-2">
                      <input className="form-check-input" type="checkbox" id="migrationNeeded" name="migrationNeeded" checked={form.migrationNeeded} onChange={updateField} />
                      <label className="form-check-label" htmlFor="migrationNeeded">We need data migration (CSV/Excel)</label>
                    </div>
                    {form.migrationNeeded && (
                      <div className="mt-2">
                        <input type="file" className="form-control" onChange={updateField} accept=".csv,.xls,.xlsx" />
                        {fileName && <div className="form-text">Selected: {fileName}</div>}
                      </div>
                    )}
                  </div>

                  <div className="col-md-6">
                    <label className="form-label">Preferred go-live</label>
                    <select name="timeline" value={form.timeline} onChange={updateField} className="form-select">
                      <option value="">Select timeline (optional)</option>
                      <option>Within 2 weeks</option>
                      <option>Within 1 month</option>
                      <option>1-3 months</option>
                      <option>3+ months</option>
                    </select>
                  </div>
                </div>
              </div>
            )}

            {/* Step 4: Payment & summary */}
            {step === 4 && (
              <div>
                <div className="row">
                  <div className="col-lg-6">
                    <div className="border rounded p-3 mb-3">
                      <h6 className="mb-2">Summary</h6>
                      <div className="small text">Check the information below before demo payment</div>
                      <ul className="small mt-3 mb-0">
                        <li>Company: {form.companyName}</li>
                        <li>Email: {form.email}</li>
                        <li>Phone: {form.phone}</li>
                        <li>Type: {form.inventoryType}</li>
                        <li>Warehouses: {form.warehouses}</li>
                        <li>Users: {form.users}</li>
                      </ul>
                    </div>
                  </div>

                  <div className="col-lg-6">
                    <div className="border rounded p-3">
                      <h6 className="mb-3">payment Summary</h6>
                      <div className="alert alert-info small mb-3">
                          Pay only ₹1 to simulate a transaction and complete the registration
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            )}

            {/* Navigation buttons */}
            <div className="d-flex justify-content-between mt-4">
              <button className="btn btn-outline-secondary" onClick={prev} disabled={step === 1 || paid}>
                Back
              </button>
              {!paid && (
                <button className="btn btn-primary" onClick={step === total ? processPayment : next} disabled={loading}>
                  {step === total ? "Pay 1" : "Next"}
                </button>
              )}
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
      {/* LOGIN POPUP MODAL */}
      {showLogin && (
          <div
            className="modal d-block"
            tabIndex="-1"
            style={{
              backgroundColor: "rgba(31, 34, 59, 0.5)",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              height: "100vh",
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
                    <label className="form-label">Tenant ID</label>
                    <input
                      type="text"
                      className="form-control"
                      name="tenantId"
                      value={loginForm.tenantId}
                      onChange={handleLoginChange}
                      placeholder="Enter Tenant ID"
                    />
                    {renderLoginErrors("tenantId")}
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
    </div>
  );
}
