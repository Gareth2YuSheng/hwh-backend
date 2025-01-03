// import React from 'react'
import { Form, FloatingLabel, Button } from 'react-bootstrap';

export default function Login() {
  return (
    <>
      <Form>
        <h1 className="mb-3">Login</h1>
        <FloatingLabel 
          controlId="floatingInput"
          label="Username"
          className="mb-3"
        >
          <Form.Control type="text" placeholder="Username" />
        </FloatingLabel>
        <FloatingLabel 
          controlId="floatingPassword" 
          label="Password"
          className="mb-3"
        >
          <Form.Control type="password" placeholder="Password" />
        </FloatingLabel>
        <Button variant="primary" type="submit">Submit</Button>
      </Form>
    </>
  );
}
