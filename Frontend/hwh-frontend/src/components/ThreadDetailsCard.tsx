import { useEffect, useState } from 'react';
//Redux
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

import { Card, Badge, Button, Image } from "react-bootstrap";
import CommentIcon from '@mui/icons-material/Comment';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

import formatDateFromUTC from "../helpers/formatDateFromUTC";

interface Props {
  viewingDetails: boolean;
  handleShowDeleteThreadModal?: () => void;
  updateThreadNavigation?: () => void;
  createCommentNavigation?: () => void;
}

export default function ThreadDetailsCard({ 
  viewingDetails, 
  handleShowDeleteThreadModal, 
  updateThreadNavigation,
  createCommentNavigation
} : Props) {
  const { thread } = useSelector((state: RootState) => state.thread);
  const { user } = useSelector((state: RootState) => state.user);

  const [timestamp, setTimestamp] = useState("");

  useEffect(() => {
    if (thread) {
      if (thread.createdAt && thread.updatedAt && thread.createdAt !== thread.updatedAt) {
        setTimestamp("Edited " + formatDateFromUTC(thread.updatedAt));
      } else {
        setTimestamp(formatDateFromUTC(thread.createdAt));
      } 
      console.log(thread)
    }    
  }, [thread]);

  return (
    <>
    {/* <div>nigga</div> */}
      {thread ? <Card className="shadow-sm mb-4">
        <Card.Body>
          <div className="d-flex justify-content-between align-items-stretch flex-wrap">
            <Card.Text style={{marginBottom:7}}>
              <Badge bg="success" className="" style={{marginRight:10, fontSize:"13px"}}>{thread.tagName}</Badge>
              <b>{thread.author}</b> • {timestamp}
            </Card.Text>
            {user !== null && viewingDetails && <div className="mb-2">
              {user.userId === thread.authorId && <Button className="mx-2" onClick={updateThreadNavigation} ><EditIcon/></Button>}
              {(user.role === "Admin" || user.userId === thread.authorId) && <Button variant="danger" onClick={handleShowDeleteThreadModal}><DeleteIcon/></Button>}
            </div>}
          </div>          
          <Card.Title><h2>{thread.title}</h2></Card.Title>
          <Card.Text className="mb-4">
            {thread.content}
          </Card.Text>
          {thread.imageURL && <div style={{justifyContent:"center", display:"flex"}}>
              <Image style={{marginBottom:"20px", maxWidth:"100%", maxHeight:"450px"}} src={thread.imageURL} />
            </div>}
          {viewingDetails && <Card.Text>
            <CommentIcon style={{marginRight:"5px", marginTop:"-3px"}}/>{thread.commentCount}
            <Button variant="warning" className="mx-3" onClick={createCommentNavigation}>Add a Comment</Button>
          </Card.Text>}
        </Card.Body>
      </Card> : 
      <Card style={{}} className="shadow-sm">
        <Card.Body>
          <Card.Title>No Thread Found</Card.Title>
        </Card.Body>
      </Card>}
    </>
  );
};
