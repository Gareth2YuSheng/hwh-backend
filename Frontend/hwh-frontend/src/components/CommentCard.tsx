import { useState, useEffect } from "react";
//Redux
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

import { Card, Badge, Button } from "react-bootstrap"
import { Comment } from "../store/commentsSlice";
import DoneIcon from '@mui/icons-material/Done';
import ThumbUpIcon from '@mui/icons-material/ThumbUp';
import ThumbDownIcon from '@mui/icons-material/ThumbDown';
import ThumbUpOffAltIcon from '@mui/icons-material/ThumbUpOffAlt';
import ThumbDownOffAltIcon from '@mui/icons-material/ThumbDownOffAlt';
import EditIcon from '@mui/icons-material/Edit';
import DeleteIcon from '@mui/icons-material/Delete';

import formatDateFromUTC from "../helpers/formatDateFromUTC";

interface Props {
  comment: Comment;
  index: number;
  handleShowDeleteModal?: () => void;
  updateCommentNavigation?: (n: number) => void;
}

export default function CommentCard({ comment, index, handleShowDeleteModal, updateCommentNavigation } : Props) {
  const { user } = useSelector((state: RootState) => state.user);

  const [voteStatus, setVoteStatus] = useState(0);

  useEffect(() => {
    if (comment.vote) {
      setVoteStatus(comment.vote.voteValue);
    }
  }, []);

  const handleUpvote = () => {
    if (voteStatus > 0) {
      unvoteComment();
      setVoteStatus(0);
    } else {
      upvoteComment();
      setVoteStatus(1);
    }
  };

  const handleDownvote = () => {
    if (voteStatus < 0) {
      unvoteComment();
      setVoteStatus(0);
    } else {
      downvoteComment();
      setVoteStatus(-1);
    }
  };

  const upvoteComment = async () => {
    console.log("Upvoting Comment")
  };

  const downvoteComment = async () => {
    console.log("Downvoting Comment")
  };

  const unvoteComment = async () => {
    console.log("Unvoting Comment")
  };
  
  return (
    <>
      <Card key={comment.commentId} style={{}} className="mt-2 shadow-sm">
        <Card.Body>
          <div className="d-flex justify-content-between align-items-stretch flex-wrap">
            <Card.Text style={{ marginBottom: 7 }}>
              <b>{comment.author}</b> • {comment.createdAt === comment.updatedAt ? 
                formatDateFromUTC(comment.createdAt) : "Edited " + formatDateFromUTC(comment.updatedAt)}
            </Card.Text>
            {user !== null && <div className="mb-2">
              {user.userId === comment.authorId && <Button className="mx-2" onClick={() => updateCommentNavigation?.(index)} ><EditIcon/></Button>}
              {(user.role === "Admin" || user.userId === comment.authorId) && <Button variant="danger" onClick={() => {}}><DeleteIcon/></Button>}
            </div>}
          </div>
          <Card.Text className="d-flex mb-2" style={{alignItems:"center"}}>
            {comment.isAnswer && <Badge bg="success" style={{ marginRight: "5px"}}><DoneIcon/></Badge>}
            {comment.content}
          </Card.Text>
          <div className="d-flex mt-3" style={{ alignItems:"center"}}>
            <Button className="bg-transparent border-0 text-dark p-0" onClick={handleUpvote}>{voteStatus === 1 ? <ThumbUpIcon /> : <ThumbUpOffAltIcon/>}</Button>
            <p className="mx-2 my-0">{comment.voteCount}</p>
            <Button className="bg-transparent border-0 text-dark p-0" onClick={handleDownvote}>{voteStatus === -1 ? <ThumbDownIcon/> : <ThumbDownOffAltIcon/>}</Button>
          </div>
        </Card.Body>
      </Card>
    </>
  );
};
