// import React from 'react'
import { SyntheticEvent, useState } from 'react';
import { Form, FloatingLabel, Button, Alert } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';

export default function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState("");

  const navigate = useNavigate();

  const handleSubmit = async (event: SyntheticEvent) => {
    event.preventDefault();
    setAlertVisible(false);
    if (password !== confirmPassword) {
      setAlertMessage("Passwords Do No Match!");
      setAlertVisible(true);
      return;
    }

    try {
      const response = await fetch(`http://localhost:8080/account/register`, {
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
        navigate(`/login`);
      } else if (content.message === "User Already Exists") {
        setAlertMessage(content.message);
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
        <h1 className="mb-4">Sign Up</h1>
        <FloatingLabel 
          controlId="floatingInput"
          label="Username"
          className="mb-3"
        >
          <Form.Control type="text" placeholder="Username" required 
            onChange={e => setUsername(e.target.value)}
            value={username}
          />
        </FloatingLabel>
        <FloatingLabel 
          controlId="floatingPassword" 
          label="Password (Optional)"
          className="mb-3"
        >
          <Form.Control type="password" placeholder="Password" 
            onChange={e => setPassword(e.target.value)}
            value={password}
          />
        </FloatingLabel>
        <FloatingLabel 
          controlId="floatingConfirmPassword" 
          label="Confirm Password (Optional)"
          className="mb-3"
        >
          <Form.Control type="password" placeholder="Confirm Password" 
            onChange={e => setConfirmPassword(e.target.value)}
            value={confirmPassword}
          />
        </FloatingLabel>
        <Button variant="primary" className="mb-2 w-100" type="submit">Sign Up</Button>
        <p>Already have an account? <Link to="/login">Login Now</Link></p>
      </Form>
    </>
  );
}