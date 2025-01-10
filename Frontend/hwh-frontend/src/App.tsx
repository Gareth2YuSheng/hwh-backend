import "./App.css";
//For Bootstrap
import "../node_modules/bootstrap/dist/css/bootstrap.min.css";

import { useEffect } from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import MyNavBar from "./components/MyNavBar";
import Home from "./pages/Home";
import Login from "./pages/Login";
import Register from "./pages/Register";
import CreateThread from "./pages/CreateThread";
import ThreadDetails from "./pages/ThreadDetails";

function App() {
  const location = useLocation();

  useEffect(() => {
    if (location.pathname === "/") {
      document.body.style.placeItems = "start"
    } else {
      document.body.style.placeItems = "center"
    }
  }, [location]);

  return (
    <div className="App">
      <MyNavBar />
      <Routes>
        <Route path="/" Component={Home} />
        <Route path="/login" Component={Login} />
        <Route path="/register" Component={Register} />
        <Route path="/createThread" Component={CreateThread} />
        <Route path="/threadDetails/:threadId" Component={ThreadDetails} />
        <Route path="/updateThread/:threadId" Component={CreateThread} />
      </Routes>        
    </div>
  );
}

export default App;
