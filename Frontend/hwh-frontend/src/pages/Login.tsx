// import React from 'react'
import { SyntheticEvent, useState } from "react";
import { Form, FloatingLabel, Button, Alert } from "react-bootstrap";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

import Cookies from "js-cookie";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState("");

  const navigate = useNavigate();

  const handleSubmit = async (event: SyntheticEvent) => {
    event.preventDefault();
    setAlertVisible(false);

    try {
      const response = await fetch(`http://localhost:8080/account/login`, {
        method: "POST",
        headers: {"Content-Type": "application/json"},
        body: JSON.stringify({
          "username": username,
          "password": password
        })
      });
      const content = await response.json();
      console.log(content) //remove later
      if (content.success) {
        const token = content.data.accessToken;
        Cookies.set("hwh-jwt", token, { expires: 2 }); //secure: true once backend uses https
        navigate(`/`);
      } else if (content.message === "Login Failed: Incorrect Username or Password") {
        setAlertMessage("Incorrect Username or Password");
        setAlertVisible(true);  
      }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  return (
    <>
      {alertVisible && alertMessage !== "" && <Alert variant="danger">{alertMessage}</Alert>}
      <Form style={{}} onSubmit={handleSubmit}>
        <h1 className="mb-4">Login</h1>
        <FloatingLabel 
          controlId="floatingInput"
          label="Username"
          className="mb-3"
        >
          <Form.Control type="text" placeholder="Username" required
            value={username} 
            onChange={e => setUsername(e.target.value)}
            />
        </FloatingLabel>
        <FloatingLabel 
          controlId="floatingPassword" 
          label="Password"
          className="mb-3"
        >
          <Form.Control type="password" placeholder="Password"
            value={password} 
            onChange={e => setPassword(e.target.value)}
            />
        </FloatingLabel>
        <Button variant="primary" className="mb-2 w-100" type="submit">Login</Button>
        <p>Don't have an account? <Link to="/register">Sign Up</Link></p>
      </Form>
    </>
  );
}
