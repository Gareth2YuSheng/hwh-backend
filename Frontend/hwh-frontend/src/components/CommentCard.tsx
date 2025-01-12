import { useState, useEffect } from "react";
//Redux
import { useSelector, useDispatch } from 'react-redux';
import { RootState, AppDispatch } from '../store/store';
import { voteComment, markAnswerComment } from "../store/commentsSlice";

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
  handleDeleteComment?: () => void;
  updateCommentNavigation?: (n: number) => void;
  token: string | undefined;
}

export default function CommentCard({ comment, index, handleDeleteComment, updateCommentNavigation, token } : Props) {
  const { user } = useSelector((state: RootState) => state.user);
  const { thread } = useSelector((state: RootState) => state.thread);
  const dispatch = useDispatch<AppDispatch>();

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
      dispatch(voteComment({ index, amt: -1 }));
    } else {
      if (voteStatus === 0) {
        dispatch(voteComment({ index, amt: 1 }));
      } else if (voteStatus < 0) {
        dispatch(voteComment({ index, amt: 2 }));
      }
      upvoteComment();
      setVoteStatus(1);
    }
  };

  const handleDownvote = () => {
    //If comment was already downvoted, unvote it
    if (voteStatus < 0) {
      unvoteComment();
      setVoteStatus(0);
      dispatch(voteComment({ index, amt: 1 }));
    } else {
    //If comment has not been downvoted, downvote it
      //If comment was not voted before, voteCount -1
      if (voteStatus === 0) {
        dispatch(voteComment({ index, amt: -1 }));
      }      
      //Else if comment was upvoted before, voteCount -2
      else if (voteStatus > 0) {
        dispatch(voteComment({ index, amt: -2 }));
      }
      downvoteComment();
      setVoteStatus(-1);
    }
  };

  const upvoteComment = async () => {
    console.log("Upvoting Comment")
    try {
      const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/vote`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
          "voteType": "up"
        })
      });
      const content = await response.json();
      console.log(content)
      // if (content.success) {
      // } else {
      // }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  const downvoteComment = async () => {
    console.log("Downvoting Comment")
    try {
      const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/vote`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        },
        body: JSON.stringify({
          "voteType": "down"
        })
      });
      const content = await response.json();
      console.log(content)
      // if (content.success) {
      // } else {
      // }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  const unvoteComment = async () => {
    console.log("Unvoting Comment")
    try {
      const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/unvote`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`
        }
      });
      const content = await response.json();
      console.log(content)
      // if (content.success) {
      // } else {
      // }
    } catch (err) {
      console.log("Error:", err);
    }
  };

  const markCommentAsAnswer = async () => {
    console.log("Marking comment as ans")
    if (comment.isAnswer) { //unmark comment
      dispatch(markAnswerComment({ index, isAns: false }));
      try {
        const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/answer?isAnswer=false`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          }
        });
        const content = await response.json();
        console.log(content)
        // if (content.success) {
        // } else {
        // }
      } catch (err) {
        console.log("Error:", err);
      }
    } else { //mark comment
      dispatch(markAnswerComment({ index, isAns: true }));
      try {
        const response = await fetch(`http://localhost:8080/comment/${comment?.commentId}/answer?isAnswer=true`, {
          method: "PUT",
          headers: {
            "Content-Type": "application/json",
            "Authorization": `Bearer ${token}`
          }
        });
        const content = await response.json();
        console.log(content)
        // if (content.success) {
        // } else {
        // }
      } catch (err) {
        console.log("Error:", err);
      }
    }
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
              {thread !== null && user.userId === thread.authorId && <Button variant="success" className="" onClick={markCommentAsAnswer} >{comment.isAnswer ? "Unmark as Answer" : "Mark as Answer"}</Button>}
              {user.userId === comment.authorId && <Button variant="primary" className="mx-2" onClick={() => updateCommentNavigation?.(index)} ><EditIcon/></Button>}
              {(user.role === "Admin" || user.userId === comment.authorId) && <Button variant="danger" onClick={handleDeleteComment}><DeleteIcon/></Button>}
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
