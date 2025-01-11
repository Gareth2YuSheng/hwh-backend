import { useEffect, useState } from 'react';
import { useParams, useNavigate, NavLink } from 'react-router-dom';
//Redux
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { fetchThreadDetails } from "../store/threadSlice";
import { fetchCommentData } from '../store/commentsSlice';
import { fetchUserData } from '../store/userSlice';

import { Spinner, Button, Modal, Alert, Pagination } from "react-bootstrap";
import DisplayComments from '../components/DisplayComments';

import Cookies from "js-cookie";
import ThreadDetailsCard from '../components/ThreadDetailsCard';

export default function ThreadDetails() {
  const { user } = useSelector((state: RootState) => state.user);
  const { thread, isLoading } = useSelector((state: RootState) => state.thread);
  const { error, totalComments } = useSelector((state: RootState) => state.comment);
  const dispatch = useDispatch<AppDispatch>();

  // const { threadId } = useParams();

  const navigate = useNavigate();

  const [modalDeleteThreadVisible, setModalDeleteThreadVisible] = useState(false);
  const [alertVisible, setAlertVisible] = useState(false);
  const [alertMessage, setAlertMessage] = useState("");
  const [alertVariant, setAlertVariant] = useState("success");
  const [disableDeleteButton, setDisableDeleteButton] = useState(false);
  const [page, setPage] = useState(1);
  const [count, setCount] = useState(5);
  const [paginationItems, setPaginationItems] = useState<JSX.Element[]>([]);

  const token = Cookies.get("hwh-jwt");

  useEffect(() => {
    //Check if user is logged in
    if(token === undefined) {
      navigate("/login");
      return;
    }
    if (thread === null) {
      navigate(-1);
      return;
    }
    if (user === null) {
      dispatch(fetchUserData(token)); 
    } 
    dispatch(fetchThreadDetails({token, threadId: thread?.threadId}));
  }, []);

  useEffect(() => {
    getComments();    
  },[page, count]);

  useEffect(() => {
    if (error) {
      setAlertVariant("danger");
      setAlertMessage(error);
      setAlertVisible(true);
    }    
  }, [error]);

  useEffect(() => {
    let newPaginationItems = [];
    for (let num = 1; num < Math.ceil(totalComments/count)+1; num++) {
      newPaginationItems.push(
        <Pagination.Item key={num} active={num === page} onClick={() => setPage(num)}>
          {num}
        </Pagination.Item>);
    }
    setPaginationItems(newPaginationItems);
  },[totalComments, page]);

  const deleteThead = async () => {
    if (user !== null && thread !== null && user.userId !== thread.authorId) {
      return;
    }
    setDisableDeleteButton(true);
    try {
      const response = await fetch(`http://localhost:8080/thread/${thread?.threadId}/delete`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        }
      });
      const content = await response.json();
      console.log(content)
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

  const getComments = () => {
    console.log("Fetching Posts - Page:", page, "Count:", count)
    dispatch(fetchCommentData({ token, page, count, threadId: thread?.threadId }));
  };

  const handleShowDeleteThreadModal = () => {
    setModalDeleteThreadVisible(true);
  };

  const handleCloseDeleteThreadModal = () => {
    setModalDeleteThreadVisible(false);
  };

  if (isLoading) {
    return <div className="d-flex justify-content-center align-items-center" style={{ height: '100vh' }}>
      <Spinner animation="border" />
    </div>;
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

      <ThreadDetailsCard viewingDetails 
        handleShowDeleteThreadModal={handleShowDeleteThreadModal} 
        updateThreadNavigation={() => navigate(`/updateThread`)}
        createCommentNavigation={() => navigate(`/createComment`)} /> 

      <DisplayComments navigate={navigate} />
      <div className="mt-4 d-flex justify-content-center">
        <Pagination>{paginationItems}</Pagination>
      </div>
    </>
  );
};
