// import React from 'react'
import { SyntheticEvent, useState } from 'react';
import { Form, FloatingLabel, Button } from 'react-bootstrap';

export default function Register() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  const submit = async (event: SyntheticEvent) => {
    event.preventDefault();
    console.log(username, password, confirmPassword)

    const response = await fetch("http://localhost:8080/account/register", {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({
        "username": username,
        "password": password
      })
    });

    const content = await response.json();
    console.log(content)
  };

  return (
    <>
      <Form style={{minWidth: 400}} onSubmit={submit}>
        <h1 className="mb-3">Sign Up</h1>
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
        <Button variant="primary" type="submit">Submit</Button>
      </Form>
    </>
  );
}