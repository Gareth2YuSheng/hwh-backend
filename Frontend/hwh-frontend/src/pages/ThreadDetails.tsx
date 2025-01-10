import { useEffect, useState } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
//Redux
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { fetchThreadDetails } from "../store/threadSlice";

import { Spinner, Card, Badge, Button, Modal, Alert } from "react-bootstrap";
import CommentIcon from '@mui/icons-material/Comment';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

import formatDateFromUTC from "../helpers/formatDateFromUTC";
import Cookies from "js-cookie";
import DisplayComments from '../components/DisplayComments';

export default function ThreadDetails() {
  const { user } = useSelector((state: RootState) => state.user);
  const { thread, isLoading } = useSelector((state: RootState) => state.thread);
  const dispatch = useDispatch<AppDispatch>();

  const { threadId } = useParams();

  const navigate = useNavigate();

  const [modalDeleteThreadVisible, setModalDeleteThreadVisible] = useState(false);
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState("");
  const [alertVariant, setAlertVariant] = useState("success");
  const [disableDeleteButton, setDisableDeleteButton] = useState(false);

  const token = Cookies.get("hwh-jwt");

  useEffect(() => {
    //Check if user is logged in
    if(token === undefined) {
      navigate("/login");
      return;
    }
    if (user === null) {
      navigate("/");
      return;
    }
    dispatch(fetchThreadDetails({token, threadId}));
  }, [])

  const deleteThead = async () => {
    if (user !== null && thread !== null && user.userId !== thread.authorId) {
      return;
    }
    setDisableDeleteButton(true);
    try {
      const response = await fetch(`http://localhost:8080/thread/${threadId}/delete`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        }
      });
      const content = await response.json();
      setModalDeleteThreadVisible(false);
      if (content.success) {
        setAlertVariant("success");
        setAlertMessage("Thread Deleted Successfully");
        setAlertVisible(true);
        setTimeout(() => navigate(`/`), 1000);
      } else if (content.message.includes("Failed to Delete Thread")) {
        setAlertMessage("Unable to Delete Thread, Something Went Wrong");
        setAlertVariant("danger"); 
        setAlertVisible(true); 
        setDisableDeleteButton(false);
      } else {
        setAlertMessage("Something Went Wrong, Try Again Later");
        setAlertVariant("danger");
        setAlertVisible(true);
        setDisableDeleteButton(false);
      }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  const handleShowDeleteThreadModal = () => {
    setModalDeleteThreadVisible(true);
  };

  const handleCloseDeleteThreadModal = () => {
    setModalDeleteThreadVisible(false);
  };

  if (isLoading) {
    return <Spinner animation="border" />;
  }

  return (
    <>
      <Modal show={modalDeleteThreadVisible} onHide={handleCloseDeleteThreadModal}>
        <Modal.Header closeButton>
          <Modal.Title>Delete Thread</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          Are you sure you want to delete this thread?
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleCloseDeleteThreadModal}>Cancel</Button>
          <Button variant="danger" disabled={disableDeleteButton} onClick={deleteThead}>Delete Thead</Button>
        </Modal.Footer>
      </Modal>
      
      {alertVisible && alertMessage !== "" && <Alert variant={alertVariant}>{alertMessage}</Alert>}

      {thread !== null ? <>
      <Card style={{}} className="shadow-sm mb-4">
        <Card.Body>
          <div className="d-flex justify-content-between align-items-stretch flex-wrap">
            <Card.Text style={{marginBottom:7}}>
              <Badge bg="success" className="" style={{marginRight:10}}>{thread.tagName}</Badge>
              <b>{thread.author}</b> - {thread.createdAt === thread.updatedAt ? formatDateFromUTC(thread.createdAt) : "Updated " + formatDateFromUTC(thread.updatedAt)}
            </Card.Text>
            {user !== null && <div className="mb-2">
              {user.userId === thread.authorId && <Button className="mx-2" onClick={() => navigate(`/updateThread/${threadId}`)} ><EditIcon/></Button>}
              {user.role === "Admin" && <Button variant="danger" onClick={handleShowDeleteThreadModal}><DeleteIcon/></Button>}
            </div>}
          </div>          
          <Card.Title><h2>{thread.title}</h2></Card.Title>
          <Card.Text>
            {thread.content}
          </Card.Text>
          <Card.Text>
            <CommentIcon style={{marginRight:"5px", marginTop:"-3px"}}/>{thread.commentCount}
          </Card.Text>
        </Card.Body>
      </Card>
      <DisplayComments />
      </> : 
      <Card style={{}} className="shadow-sm">
        <Card.Body>
          <Card.Title>No Thread Found</Card.Title>
        </Card.Body>
      </Card>}
    </>
  );
};
