import { BACKEND_API_BASE_URL } from "../../src/config";

//Function for customers to get all tenants informations
export async function getAllTenants(customersId, token) {
  try {
    const response = await fetch(`${BACKEND_API_BASE_URL}/customers/${customersId}/tenants/`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data; // Return the tenant data
  } catch (err) {
    console.error("Failed to fetch tenants:", err);
    return null; // or return [] if you want an empty array on failure
  }
}


//Function for customers to get perticular tenant information 
export async function getTenants(customersId, token ,tenantId) {
  try {
    const response = await fetch(`${BACKEND_API_BASE_URL}/customers/${customersId}/tenants/${tenantId}`, {
      headers: { Authorization: `Bearer ${token}` },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    return data; // Return the tenant data
  } catch (err) {
    console.error("Failed to fetch tenants:", err);
    return null; // or return [] if you want an empty array on failure
  }
}