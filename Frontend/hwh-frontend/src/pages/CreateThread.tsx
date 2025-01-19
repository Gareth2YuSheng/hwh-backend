import { useState, useEffect, SyntheticEvent } from "react";
import { Form, FloatingLabel, Button, Alert, Modal, Badge, Image } from "react-bootstrap";
import { useNavigate } from "react-router-dom";

//Redux
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { fetchTagData } from "../store/tagSlice";

import Cookies from "js-cookie";
import TagsDropdown from "../components/TagsDropdown";

interface Props {
  mode: "CREATE" | "UPDATE"
}

export default function CreateThread({ mode }: Props) {
  //Redux
  const { thread } = useSelector((state: RootState) => state.thread);
  const dispatch = useDispatch<AppDispatch>();

  // const { threadId } = useParams();
  
  const navigate = useNavigate();

  const [alertVisible, setAlertVisible] = useState(false);
  const [alertVariant, setAlertVariant] = useState("success");
  const [alertMessage, setAlertMessage] = useState("");
  const [title, setTitle] = useState("");
  const [threadContent, setContent] = useState("");
  const [disableSubmitBtn, setDisableSubmitBtn] = useState(false);
  const [image, setImage] = useState<File | null>(null);
  //For Tag Dropdown
  const [selectedTagId, setSelectedTagId] = useState("");
  const [showCreateTagModal, setShowCreateTagModal] = useState(false);
  const [newTagName, setNewTagName] = useState("");
  const [modalAlertVisible, setModalAlertVisible] = useState(false);
  const [modalAlertMessage, setModalAlertMessage] = useState("");
  const [modalAlertVariant, setModalAlertVariant] = useState("success");
  const [disableModalSubmitBtn, setDisableModalSubmitBtn] = useState(false)

  const token = Cookies.get("hwh-jwt");

  useEffect(() => {
    //Check if user is logged in
    if(token === undefined) {
      navigate("/login");
      return;
    }
    if (mode === "CREATE") {
      getTags();
    } else if (mode === "UPDATE" && thread) {
      setTitle(thread.title);
      setContent(thread.content);
    } else if (mode === "UPDATE" && thread === null) {
      navigate(-1);
    }
  }, []);

  const getTags = () => {
    console.log("Fetching Tags");
    dispatch(fetchTagData(token));
  }

  const handleSubmit = async (event: SyntheticEvent) => {
    event.preventDefault();
    setAlertVisible(false);    
    console.log("Title:",title, "| Content:", threadContent,"| Tag:", selectedTagId);
    if (title === "" || threadContent === "") {
      setAlertMessage("Thread Title or Content cannot be empty!");
      setAlertVariant("danger");
      setAlertVisible(true);
      return;
    }
    if (mode === "CREATE" && selectedTagId === "") {
      setAlertMessage("A Tag is Required for the thread!");
      setAlertVariant("danger");
      setAlertVisible(true);
      return;
    }
    setDisableSubmitBtn(true);
    if (mode === "UPDATE") {
      try {
        const response = await fetch(`http://localhost:8080/thread/${thread?.threadId}/update`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          },
          body: JSON.stringify({
            "title": title,
            "content": threadContent
          })
        });
        const content = await response.json();
        if (content.success) {
          setAlertMessage("Thread Updated Successfully");
          setAlertVariant("success")
          setAlertVisible(true);
          setTimeout(() => navigate(-1), 1000);
        } else if (content.message.includes("Failed to Update Thread")) {
          setAlertMessage("Unable to Update Thread, Something Went Wrong");
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
      const formData = new FormData();
      formData.append("title", title)
      formData.append("content", threadContent)
      formData.append("tagId", selectedTagId)
      if (image) {
        formData.append("image", image);
      }      
      try {
        const response = await fetch(`http://localhost:8080/thread/create`, {
          method: "POST",
          headers: {
            "Authorization": `Bearer ${token}`
          },
          body: formData
        });
        const content = await response.json();
        if (content.success) {
          setAlertMessage("Thread Created Successfully");
          setAlertVariant("success")
          setAlertVisible(true);
          setTimeout(() => navigate(`/`), 1000);
        } else if (content.message.includes("Failed to Create Thread")) {
          setAlertMessage("Unable to Create Thread, Something Went Wrong");
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

  const handleCreateNewTag = async (event: SyntheticEvent) => {
    event.preventDefault();
    setDisableModalSubmitBtn(true);
    setModalAlertVisible(false);
    console.log("New Tag name:",newTagName);
    if (newTagName === "") {
      setModalAlertVariant("danger");
      setModalAlertMessage("New Tag Name Cannot be Empty!");
      setModalAlertVisible(true);
      return;
    }

    try {
      const response = await fetch(`http://localhost:8080/tag/create`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
          "name": newTagName
        })
      });
      const content = await response.json();
      if (content.success) {
        setModalAlertVariant("success");
        setModalAlertMessage("Tag Created Successfully, please refresh the page the see the new tag");
        setModalAlertVisible(true);
        setTimeout(() => handleCloseCreateTagModal(), 1000);
      } else if (content.message.includes("Failed to Create Tag")) {
        setModalAlertVariant("danger");
        setModalAlertMessage("Unable to Create Tag, Something Went Wrong");
        setModalAlertVisible(true);  
        setDisableModalSubmitBtn(false);
      }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  const handleShowCreateTagModal = () => {
    setShowCreateTagModal(true);
  };

  const handleCloseCreateTagModal = () => {
    setShowCreateTagModal(false);
  };

  return (
    <>
      {alertVisible && alertMessage !== "" && <Alert variant={alertVariant}>{alertMessage}</Alert>}

      {/* Create Tag Modal */}
      <Modal show={showCreateTagModal} onHide={handleCloseCreateTagModal}>
        <Modal.Header closeButton>
          <Modal.Title>Create New Tag</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {modalAlertVisible && modalAlertMessage !== "" && <Alert variant={modalAlertVariant}>{modalAlertMessage}</Alert>}
          <FloatingLabel 
            controlId="floatingInput"
            label="New Tag Name"
            className="mb-3"
          >
            <Form.Control type="text" placeholder="Name" 
              value={newTagName} 
              onChange={e => setNewTagName(e.target.value)}
              />
          </FloatingLabel>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseCreateTagModal}>Cancel</Button>
          <Button variant="success" type="submit" disabled={disableModalSubmitBtn} onClick={handleCreateNewTag}>Create Tag</Button>
        </Modal.Footer>
      </Modal>

      <Form style={{}} onSubmit={handleSubmit} autoComplete="off">
        <div className="d-flex mt-4 justify-content-between align-items-stretch flex-wrap">
          {mode === "UPDATE" ? <>
            <h2 className="mb-3">Update Thread</h2>
            <Badge style={{ alignSelf:"baseline", fontSize:"18px", padding:"13px" }} bg="success">{thread?.tagName}</Badge>
          </> : <>
            <h2 className="mb-3">New Thread</h2>
            <TagsDropdown selectTagFn={setSelectedTagId} forFilter={false} openModalFn={handleShowCreateTagModal} />
          </>}
        </div>
        
        <FloatingLabel 
          controlId="floatingInput"
          label="Title"
          className="mb-3"
        >
          <Form.Control type="text" placeholder="Title" required 
            value={title} 
            onChange={e => setTitle(e.target.value)}
            />
        </FloatingLabel>
        <FloatingLabel 
          controlId="floatingInput" 
          label="Content"
          className="mb-3"
        >
          <Form.Control as="textarea" style={{minHeight:200}} placeholder="Content" required
            value={threadContent} 
            onChange={e => setContent(e.target.value)}
            />
        </FloatingLabel>
        {mode === "CREATE" ? <Form.Control className="mb-3" type="file" placeholder="Image" accept="image/*"
          onChange={e => {
            if (e.target instanceof HTMLInputElement) {
              if (e.target.files != null && e.target.files.length > 0) {
                setImage(e.target.files[0]);
              }              
            }
          }} /> : <div style={{justifyContent:"center", display:"flex"}}>
            <Image style={{marginBottom:"20px", maxHeight:"400px", maxWidth:"100%"}} src={thread?.imageURL} />
          </div>}

        <Button variant="warning" className="mb-2 w-100" type="submit" disabled={disableSubmitBtn}>{mode === "UPDATE" ? "Update Thread" : "Create Thread"}</Button>
      </Form>
    </>
  );
};
