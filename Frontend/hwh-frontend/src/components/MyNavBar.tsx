// import React from 'react'
import { Nav, Navbar, Container } from "react-bootstrap";
import { NavLink } from "react-router-dom";

export default function MyNavBar() {
  return (
    <>
      <Navbar expand="lg" className="bg-body-tertiary" fixed="top" bg="dark" data-bs-theme="dark">
        <Container>
          <Navbar.Brand href="#home">Homework Help</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
          <Nav className="me-auto" variant="underline">
            <Nav.Item>
              <NavLink style={{ textDecoration: "none", color: '#FFF' }} to="/login">Login</NavLink>
            </Nav.Item>
            <Nav.Item>
              <NavLink style={{ textDecoration: "none", color: '#FFF' }} to="/register">Registration</NavLink>
            </Nav.Item>
            
            {/* <Nav.Link href="#home"><Link to="/login">Login</Link></Nav.Link>
            <Nav.Link href="#link"><Link to="/register">Registration</Link></Nav.Link> */}
            {/* <NavDropdown title="Dropdown" id="basic-nav-dropdown">
            <NavDropdown.Item href="#action/3.1">Action</NavDropdown.Item>
            <NavDropdown.Item href="#action/3.2">
              Another action
            </NavDropdown.Item>
            <NavDropdown.Item href="#action/3.3">Something</NavDropdown.Item>
            <NavDropdown.Divider />
            <NavDropdown.Item href="#action/3.4">
              Separated link
            </NavDropdown.Item>
            </NavDropdown> */}
          </Nav>
          <Navbar.Collapse className="justify-content-end mt-3 mb-2">
            <Navbar.Text>
              Signed in as: <a href="#login">Robin Banks</a>
            </Navbar.Text>
          </Navbar.Collapse>
          </Navbar.Collapse>
        </Container>
      </Navbar>
      {/* <Nav variant="underline" defaultActiveKey="/home">
        <Nav.Item>
          <Nav.Link href="/home">Active</Nav.Link>
        </Nav.Item>
        <Nav.Item>
          <Nav.Link eventKey="link-1">Option 2</Nav.Link>
        </Nav.Item>
        <Nav.Item>
          <Nav.Link eventKey="disabled" disabled>
            Disabled
          </Nav.Link>
        </Nav.Item>
      </Nav> */}
    </>
  );
}
