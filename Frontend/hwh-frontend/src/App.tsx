import "./App.css";
//For Bootstrap
import "../node_modules/bootstrap/dist/css/bootstrap.min.css";

import { useEffect } from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import MyNavBar from "./components/MyNavBar";
import Home from "./pages/Home";
import Login from "./pages/Login";
import CreateThread from "./pages/CreateThread";
import ThreadDetails from "./pages/ThreadDetails";
import CreateComment from "./pages/CreateComment";

function App() {
  const location = useLocation();

  useEffect(() => {
    if (location.pathname === "/" 
      || location.pathname.startsWith("/threadDetails") 
      || location.pathname.startsWith("/updateThread")
      || location.pathname.startsWith("/createComment")
      || location.pathname.startsWith("/updateComment")) {
      document.body.style.placeItems = "start";
      document.body.style.paddingTop = "100px";
    } else {
      document.body.style.placeItems = "center";
      document.body.style.paddingTop = "0px";
    }
  }, [location]);

  return (
    <div className="App">
      <MyNavBar />
      <Routes>
        <Route path="/" Component={Home} />
        <Route path="/login" element={<Login mode="LOGIN"/>} />
        <Route path="/register" element={<Login mode="REGISTER"/>} />
        <Route path="/createThread" element={<CreateThread mode="CREATE"/>} />
        <Route path="/threadDetails" Component={ThreadDetails} />
        <Route path="/updateThread" element={<CreateThread mode="UPDATE"/>} />
        <Route path="/createComment" element={<CreateComment mode="CREATE"/>} />
        <Route path="/updateComment" element={<CreateComment mode="UPDATE"/>} />
      </Routes>        
    </div>
  );
}

export default App;
