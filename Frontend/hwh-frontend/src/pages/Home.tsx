// import React from 'react'
import { useEffect } from "react";
import { useNavigate } from "react-router-dom";

import Cookies from "js-cookie";

export default function Home() {
  const navigate = useNavigate();

  useEffect(() => {
    //Check if user is logged in
    const token = Cookies.get("hwh-jwt");
    if(token === undefined) {
      navigate("/login");
      return;
    }
    (
      async () => {
        console.log("awaiting User Data"); //remove later
        const response = await fetch(`http://localhost:8080/account/user`, {
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          }
        });
        const content = await response.json();
        console.log(content.data.user); //remove later
        //Store userData in Redux
      }
    )();
    
  });

  return (
    <div>Home</div>
  )
}
