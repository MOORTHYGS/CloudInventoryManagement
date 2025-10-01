// src/components/Footer.js
import React from "react";

const Footer = () => {
  return (
    <footer className="text-light py-0 mt-auto">
      <div className="container d-flex justify-content-between align-items-center">
        <div>
          &copy; {new Date().getFullYear()} CloudInventory. All rights reserved.
        </div>
      </div>
    </footer>
  );
};

export default Footer;
