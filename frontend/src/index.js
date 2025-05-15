// frontend/src/index.js
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App"; // or "./App.jsx" if your file is named App.jsx

// Attach React to the <div id="root"> in public/index.html
const rootElement = document.getElementById("root");
const root = ReactDOM.createRoot(rootElement);

// Render the top-level <App /> component within React.StrictMode
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
