import React, { useEffect, useState, useRef } from "react";
import Header from "../common/Header";
import Footer from "../common/Footer";
import "../../css files/customer/Dashboard.css";
import { getAllTenants} from "../../../utils/endpoints";
import { saveTokenFromURL, checkAuth, setupInactivityLogout } from "../../../utils/auth";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircle, faBell } from "@fortawesome/free-solid-svg-icons";


const Dashboard = () => {
  const [loading, setLoading] = useState(true);
  const [tenants, setTenants] = useState([]);
  const didFetch = useRef(false);

  useEffect(() => {
    if (didFetch.current) return;
    didFetch.current = true;

    saveTokenFromURL();
    const auth = checkAuth();
    if (!auth) return;

    const cleanup = setupInactivityLogout(3); // 3 minutes inactivity logout

    const fetchTenantData = async () => {
      const data = await getAllTenants(auth.user_id, auth.token);
      setTenants(data);
      setLoading(false);
    };

    fetchTenantData();
    return cleanup;
  }, []);

  const getStatusColor = (status) => {
    switch (status.toLowerCase()) {
      case "active":
        return "text-success"; // green
      case "inactive":
        return "text-danger"; // red
      case "pending":
        return "text-warning"; // yellow
      default:
        return "text-secondary"; // gray
    }
  };

  return (
    <div className="d-flex flex-column min-vh-100">
      <Header />

      <div className="default-container">
        {/* Loading state */}
        {loading && (
          <div>
            <p>Loading Tenant Data...</p>
            <p>Welcome to CloudInventory!</p>
          </div>
        )}

        {/* Tenant cards */}
        {!loading && tenants.length > 0 && (
          <div className="tenant-card card-glass p-3 mb-3">
            <h2 className="title text-center mb-4">Tenants</h2>
            <div className="row g-4">
              {tenants.map((tenant) => (
                <div key={tenant.id} className="col-md-4">
                  <div
                    className="card shadow-sm cursor-pointer"
                    onClick={() => console.log("Selected Tenant:", tenant)}
                  >
                    <div className="card-body d-flex justify-content-between align-items-start">
                      <div>
                        <h5 className="card-title">{tenant.name}</h5>
                        <p className="card-text mb-1">
                          <strong>Domain:</strong> {tenant.domain}
                        </p>
                        <p className="card-text mb-1">
                          <strong>Status:</strong> {tenant.status}
                        </p>
                        <p className="card-text">
                          <strong>Created:</strong>{" "}
                          {new Date(tenant.created_at).toLocaleString()}
                        </p>
                      </div>

                      {/* Status and Message icons in a row */}
                      <div className="d-flex flex-row align-items-center gap-2">
                        <FontAwesomeIcon
                          icon={faCircle}
                          className={getStatusColor(tenant.status)}
                          size="sm"
                        />
                        <FontAwesomeIcon icon={faBell} className="text-primary" size="sm" />
                      </div>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Fallback if no tenants */}
        {!loading && tenants.length === 0 && (
          <p>No tenants found. Click "Add Tenant" to get started!</p>
        )}
      </div>

      <Footer />
    </div>
  );
};

export default Dashboard;
