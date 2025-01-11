import { useState, useEffect, SyntheticEvent } from "react";
import { Form, FloatingLabel, Button, Alert } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

//Redux
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

import Cookies from "js-cookie";
import ThreadDetailsCard from "../components/ThreadDetailsCard";

interface Props {
  mode: "CREATE" | "UPDATE"
}

export default function CreateComment({ mode }: Props) {
  //Redux
  const { comment } = useSelector((state: RootState) => state.comment);
  const { thread } = useSelector((state: RootState) => state.thread);
  
  const navigate = useNavigate();

  const [alertVisible, setAlertVisible] = useState(false);
  const [alertVariant, setAlertVariant] = useState("success");
  const [alertMessage, setAlertMessage] = useState("");
  const [commentContent, setContent] = useState("");
  const [disableSubmitBtn, setDisableSubmitBtn] = useState(false);

  const token = Cookies.get("hwh-jwt");

  useEffect(() => {
      //Check if user is logged in
      if(token === undefined) {
        navigate("/login");
        return;
      }
      if (thread === null) {
        navigate("/");
        return;
      }
      console.log("Comment:",comment)
      if (mode === "UPDATE" && comment) {
        setContent(comment.content);
      } else if (mode === "UPDATE" && comment === null) {
        navigate(-1);
      }
    }, []);

  const handleSubmit = async (event: SyntheticEvent) => {
    event.preventDefault();
    setAlertVisible(false);    
    console.log("Content:", commentContent);
    if (commentContent === "") {
      setAlertMessage("Comment Content cannot be empty!");
      setAlertVariant("danger");
      setAlertVisible(true);
      return;
    }
    setDisableSubmitBtn(true);
    if (mode === "UPDATE") {
      try {
        const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/update`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          },
          body: JSON.stringify({
            "content": commentContent
          })
        });
        const content = await response.json();
        if (content.success) {
          setAlertMessage("Comment Updated Successfully");
          setAlertVariant("success");
          setAlertVisible(true);
          setTimeout(() => navigate(-1), 1000);
        } else if (content.message.includes("Failed to Update Thread")) {
          setAlertMessage("Unable to Update Comment, Something Went Wrong");
          setAlertVariant("danger");
          setAlertVisible(true);
          setDisableSubmitBtn(false);
        } else {
          setAlertMessage("Something Went Wrong, Try Again Later");
          setAlertVariant("danger");
          setAlertVisible(true);
          setDisableSubmitBtn(false);
        }
      } catch (err) {
        console.log("Error:", err);
      }
    } else if (mode === "CREATE") {
      try {
        const response = await fetch(`http://localhost:8080/comment/${thread?.threadId}/create`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          },
          body: JSON.stringify({
            "content": commentContent
          })
        });
        const content = await response.json();
        if (content.success) {
          setAlertMessage("Comment Created Successfully");
          setAlertVariant("success")
          setAlertVisible(true);
          setTimeout(() => navigate(`/`), 1000);
        } else if (content.message.includes("Failed to Create Comment")) {
          setAlertMessage("Unable to Create Comment, Something Went Wrong");
          setAlertVariant("danger");
          setAlertVisible(true);
          setDisableSubmitBtn(false);
        } else {
          setAlertMessage("Something Went Wrong, Try Again Later");
          setAlertVariant("danger");
          setAlertVisible(true);
          setDisableSubmitBtn(false);
        }
      } catch (err) {
        console.log("Error:", err);
      }
    }
  };

  return (
    <>
      {alertVisible && alertMessage !== "" && <Alert variant={alertVariant}>{alertMessage}</Alert>}

      <ThreadDetailsCard viewingDetails={false}/>

      <Form style={{}} onSubmit={handleSubmit} autoComplete="off">
        <div className="d-flex mt-4 justify-content-between align-items-stretch flex-wrap">
          {mode === "UPDATE" ?
            <h2 className="mb-3">Update Comment</h2>
            : <h2 className="mb-3">New Comment</h2>}
        </div>
        
        <FloatingLabel 
          controlId="floatingInput" 
          label="Content"
          className="mb-3"
        >
          <Form.Control as="textarea" style={{minHeight:200}} placeholder="Content" required
            value={commentContent} 
            onChange={e => setContent(e.target.value)}
            />
        </FloatingLabel>
        <Button variant="warning" className="mb-2 w-100" type="submit" disabled={disableSubmitBtn}>{mode === "UPDATE" ? "Update Comment" : "Create Comment"}</Button>
      </Form>
    </>
  );
};
