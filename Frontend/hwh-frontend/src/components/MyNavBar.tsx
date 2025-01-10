import { Nav, Navbar, Container } from "react-bootstrap";
import { NavLink } from "react-router-dom";
//Redux
import { useSelector, useDispatch } from 'react-redux';
import { RootState, AppDispatch } from '../store/store';
import { clearUser } from "../store/userSlice";

import AutoStoriesIcon from '@mui/icons-material/AutoStories';

import Cookies from "js-cookie";

import "../styles/MyNavBarStyles.css";

export default function MyNavBar() {
  //Redux
  const { user } = useSelector((state: RootState) => state.user);
  const dispatch = useDispatch<AppDispatch>();

  const handleLogout = () => {
    Cookies.remove("hwh-jwt");
    dispatch(clearUser());
  };

  return (
    <>
      <Navbar expand="lg" className="bg-body-tertiary" fixed="top" bg="dark" data-bs-theme="dark">
        <Container>
          <Navbar.Brand style={{ fontWeight: 500, verticalAlign: "baseline", fontSize: 25 }}><AutoStoriesIcon style={{fontSize:35, margin: 10}} />Homework Help</Navbar.Brand>
          <Navbar.Toggle aria-controls="basic-navbar-nav" />
          <Navbar.Collapse id="basic-navbar-nav">
            <Nav className="mx-3" variant="underline">
              {user === null ? <>
                <Nav.Item>
                  <NavLink className="myLink" to="/login">Login</NavLink>
                </Nav.Item>
                <Nav.Item>
                  <NavLink className="myLink" to="/register">Registration</NavLink>
                </Nav.Item>
              </> : <>
                <Nav.Item>
                  <NavLink className="myLink" to="/">Home</NavLink>
                </Nav.Item>
              </>}
              
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
            {user !== null && <Navbar.Collapse className="justify-content-end mt-3 mb-2">
              <Navbar.Text className="m-2">
                Signed in as: <NavLink className="myLink" to="/">{user.username}</NavLink>
              </Navbar.Text>
              <Nav.Item className="m-2">
                <NavLink className="myLink" to="/login" onClick={handleLogout} >Logout</NavLink>
              </Nav.Item>
            </Navbar.Collapse>}
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
