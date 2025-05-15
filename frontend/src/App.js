import React, { useState } from 'react';
import { getToken, setToken } from './auth';

function getRole() {
  return localStorage.getItem("role") || "user";
}
function setRole(role) {
  localStorage.setItem("role", role);
}

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [imageFile, setImageFile] = useState(null);
  const [imageID, setImageID] = useState("");
  const [ghibliImages, setGhibliImages] = useState([]);
  const [orderId, setOrderId] = useState("");
  const [orderInfo, setOrderInfo] = useState(null);
  const [uploading, setUploading] = useState(false);
  const [jobStatus, setJobStatus] = useState("");
  const [lastHash, setLastHash] = useState("");

  const gatewayBase = "http://localhost:8080/api"; // Adjust in Docker env

  const signup = async () => {
    const res = await fetch(`${gatewayBase}/auth/signup`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    });
    if(res.ok) {
      alert("Signed up successfully");
    } else {
      alert("Sign up failed");
    }
  }

  const login = async () => {
    const res = await fetch(`${gatewayBase}/auth/login`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    });
    if(res.ok) {
      const data = await res.json();
      setToken(data.token);
      // Decode JWT to get role
      const payload = JSON.parse(atob(data.token.split('.')[1]));
      setRole(payload.role);
      alert("Logged in!");
    } else {
      alert("Login failed");
    }
  }

  const uploadImage = async () => {
    if(!imageFile) return;
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }
    if(uploading) return;
    setUploading(true);
    setJobStatus("");
    const res = await fetch(`${gatewayBase}/upload`, {
      method: "POST",
      headers: { "Authorization": `Bearer ${token}`, "X-User": email },
      body: imageFile
    });
    setUploading(false);
    if(res.ok) {
      const data = await res.json();
      setImageID(data.imageID);
      setLastHash(data.imageID);
      alert("Image uploaded");
      // Poll for job status
      pollJobStatus(data.imageID);
    } else if(res.status === 409) {
      alert("Duplicate job");
    } else {
      alert("Upload failed");
    }
  }

  const pollJobStatus = (hash) => {
    setJobStatus("Processing...");
    let tries = 0;
    const poll = async () => {
      tries++;
      const res = await fetch(`http://localhost:8083/job-status?hash=${hash}`);
      if(res.ok) {
        const data = await res.json();
        if(data.status.startsWith("done:")) {
          setJobStatus("Done");
          setGhibliImages([{url: data.status.replace("done:", "")}]);
          return;
        } else {
          setJobStatus(data.status);
        }
      }
      if(tries < 20) setTimeout(poll, 2000);
      else setJobStatus("Timeout");
    };
    poll();
  };

  const generateFree = async () => {
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }

    const url = `${gatewayBase}/ghibli/free/${imageID}`;
    const res = await fetch(url, {
      method: "GET",
      headers: { "Authorization": `Bearer ${token}` }
    });
    const data = await res.json();
    setGhibliImages(data);
  }

  const generatePaid = async () => {
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }

    const url = `${gatewayBase}/ghibli/paid/${imageID}`;
    const res = await fetch(url, {
      method: "GET",
      headers: { "Authorization": `Bearer ${token}` }
    });
    const data = await res.json();
    setGhibliImages(data);
  }

  const createOrder = async (orderType) => {
    // orderType = "TEN_IMAGES" or "3D_MODEL"
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }

    const res = await fetch(`${gatewayBase}/orders`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${token}`,
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        email,
        orderType,
        address: "123 Demo St",
        imageRef: orderType === "3D_MODEL" ? "CHOSEN_IMAGE_URL" : ""
      })
    });
    if(res.ok) {
      const data = await res.json();
      setOrderId(data.ID);
      alert(`Order created with ID: ${data.ID}`);
    } else {
      alert("Order creation failed");
    }
  }

  const payOrder = async () => {
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }

    const res = await fetch(`${gatewayBase}/orders/${orderId}/pay`, {
      method: "POST",
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });
    if(res.ok) {
      const data = await res.json();
      alert(`Order paid. Current status: ${data.Status}`);
    } else {
      alert("Payment failed");
    }
  }

  const getOrderInfo = async () => {
    const token = getToken();
    if(!token) {
      alert("Please login first");
      return;
    }

    const res = await fetch(`${gatewayBase}/orders/${orderId}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${token}`
      }
    });
    if(res.ok) {
      const data = await res.json();
      setOrderInfo(data);
    } else {
      alert("Failed to get order info");
    }
  }

  return (
    <div style={{ padding: 20 }}>
      <h1>Get-Ur-Ghibli</h1>
      <div>
        <h2>Auth</h2>
        <input type="email" placeholder="Email" value={email} onChange={e=>setEmail(e.target.value)} />
        <input type="password" placeholder="Password" value={password} onChange={e=>setPassword(e.target.value)} />
        <button onClick={signup}>Sign Up</button>
        <button onClick={login}>Log In</button>
      </div>

      <div>
        <h2>Upload</h2>
        <input type="file" onChange={e => setImageFile(e.target.files[0])} />
        <button onClick={uploadImage} disabled={uploading}>Upload</button>
        {jobStatus && <p>Job Status: {jobStatus}</p>}
      </div>

      <div>
        <h2>Generate Ghibli</h2>
        <p>Image ID: {imageID}</p>
        <button onClick={generateFree}>Generate Free Image</button>
        <button onClick={generatePaid}>Generate 10 Paid Images</button>
        {ghibliImages.length > 0 && (
          <ul>
            {ghibliImages.map((img, i) => (
              <li key={i}>
                <a href={img.url} target="_blank" rel="noreferrer">{img.url}</a>
              </li>
            ))}
          </ul>
        )}
      </div>

      <div>
        <h2>Orders</h2>
        <button onClick={()=>createOrder("TEN_IMAGES")}>Create Order for 10 Images (Rs.100)</button>
        <button onClick={()=>createOrder("3D_MODEL")}>Create Order for 3D Model (Rs.2000)</button>
        {orderId && <p>Order ID: {orderId}</p>}
        <button onClick={payOrder}>Pay for Order</button>
        <button onClick={getOrderInfo}>Get Order Info</button>
        {orderInfo && (
          <pre>{JSON.stringify(orderInfo, null, 2)}</pre>
        )}
      </div>

      {getRole() === "admin" && (
        <div style={{background: '#eee', padding: 10, marginTop: 20}}>
          <h2>Admin Panel</h2>
          <p>Admin-only features go here.</p>
        </div>
      )}
    </div>
  );
}

export default App;
