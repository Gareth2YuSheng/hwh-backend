//Redux
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

import { Spinner, Card, Badge } from "react-bootstrap";

import formatDateFromUTC from "../helpers/formatDateFromUTC";

export default function DisplayComments() {
  const { comments, isLoading } = useSelector((state: RootState) => state.comment);

  // if (isLoading) {
  //   return <Spinner animation="border" />;
  // }

  return (
    <>
      {comments.length > 0 ? <div className="mt-4">
        {comments.map((comment, index) => {
          return <Card key={comment.commentId} style={{cursor:"pointer"}} 
            className="mt-2 shadow-sm" onClick={() => {}}>
            <Card.Body>
              <Card.Title>{comment.content} 
                {/* <Badge bg="success" className="mx-1">{thread.tagName}</Badge> */}
              </Card.Title>
              <Card.Text style={{marginBottom:7}}>
                <b>{comment.author}</b> - {comment.createdAt === comment.updatedAt 
                  ? formatDateFromUTC(comment.createdAt) : 
                  "Updated " + formatDateFromUTC(comment.updatedAt)}
              </Card.Text>
              {/* <Card.Text>
                <CommentIcon style={{marginRight:"5px", marginTop:"-3px"}}/>{thread.commentCount}
              </Card.Text> */}
            </Card.Body>
          </Card>;
        })}
      </div> : 
      <div>
        <Card className="mt-2 shadow-sm">
          <Card.Body>
            <Card.Title>No Comments Found</Card.Title>
          </Card.Body>
        </Card>
      </div>}
    </>
  );
};
