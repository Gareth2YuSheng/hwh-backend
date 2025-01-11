import { NavigateFunction } from 'react-router-dom';
//Redux
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { selectThread } from '../store/threadSlice';

import { Spinner, Card, Badge } from "react-bootstrap";
import CommentIcon from '@mui/icons-material/Comment';

import formatDateFromUTC from "../helpers/formatDateFromUTC";

interface Props {
  navigate: NavigateFunction;
}

export default function DisplayThreads({ navigate } : Props) {
  const { threads, isLoading } = useSelector((state: RootState) => state.thread);
  const dispatch = useDispatch<AppDispatch>();

  if (isLoading) {
    return <div className="d-flex justify-content-center align-items-center" style={{ height: '100vh' }}>
      <Spinner animation="border" />
    </div>;
  }

  const handleViewThreadDetails = (index: number) => {
    dispatch(selectThread({index}));
    navigate(`/threadDetails`);
  };

  return (
    <>
      {threads.length > 0 ? <div className="mt-4">
        {threads.map((thread, index) => {
          return <Card key={thread.threadId} style={{cursor:"pointer"}} 
            className="mt-2 shadow-sm" onClick={() => handleViewThreadDetails(index)}>
            <Card.Body>
              <Card.Title>{thread.title} <Badge bg="success" className="mx-1">{thread.tagName}</Badge></Card.Title>
              <Card.Text style={{marginBottom:7}}>
                <b>{thread.author}</b> • {formatDateFromUTC(thread.createdAt)}
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
