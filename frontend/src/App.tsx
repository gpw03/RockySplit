import { useEffect, useState } from "react";

function App() {
  // health will hold whatever the API returns (e.g., "ok")
  const [health, setHealth] = useState<string>("checking...");
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    // This runs once when the component mounts
    fetch("http://localhost:8080/api/healthz")
      .then((res) => {
        if (!res.ok) {
          throw new Error(`HTTP error ${res.status}`);
        }
        return res.text();
      })
      .then((text) => {
        setHealth(text);
      })
      .catch((err) => {
        console.error("Error calling /api/healthz:", err);
        setError(err.message);
      });
  }, []); // empty deps array = run once

  return (
    <div style={{ fontFamily: "system-ui", padding: "2rem" }}>
      <h1>My App</h1>
      <p>Backend health check:</p>
      {error ? (
        <p style={{ color: "red" }}>Error: {error}</p>
      ) : (
        <p>Status: {health}</p>
      )}
    </div>
  );
}

export default App;
