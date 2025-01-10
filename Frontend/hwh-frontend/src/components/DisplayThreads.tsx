import { NavigateFunction } from 'react-router-dom';

//Redux
import { useSelector } from 'react-redux';
import { RootState } from '../store/store';

import { Spinner, Card, Badge } from "react-bootstrap";
import CommentIcon from '@mui/icons-material/Comment';

import formatDateFromUTC from "../helpers/formatDateFromUTC";

interface Props {
  navigate: NavigateFunction;
}

export default function DisplayThreads({ navigate } : Props) {
  const { threads, isLoading } = useSelector((state: RootState) => state.thread);

  if (isLoading) {
    return <Spinner animation="border" />;
  }

  return (
    <>
      {threads.length > 0 ? <div className="mt-4">
        {threads.map((thread, index) => {
          return <Card key={thread.threadId} style={{cursor:"pointer"}} 
            className="mt-2 shadow-sm" onClick={() => navigate(`/threadDetails/${thread.threadId}`)}>
            <Card.Body>
              <Card.Title>{thread.title} <Badge bg="success" className="mx-1">{thread.tagName}</Badge></Card.Title>
              <Card.Text style={{marginBottom:7}}>
                <b>{thread.author}</b> - {formatDateFromUTC(thread.createdAt)}
              </Card.Text>
              <Card.Text>
                <CommentIcon style={{marginRight:"5px", marginTop:"-3px"}}/>{thread.commentCount}
              </Card.Text>
            </Card.Body>
          </Card>;
        })}
      </div> : 
      <div>
        <Card className="mt-2 shadow-sm">
          <Card.Body>
            <Card.Title>No Threads Found</Card.Title>
          </Card.Body>
        </Card>
      </div>}
    </>
  );
};
