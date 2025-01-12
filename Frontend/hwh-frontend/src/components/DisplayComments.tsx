import { NavigateFunction } from 'react-router-dom';
//Redux
import { useDispatch, useSelector } from 'react-redux';
import { AppDispatch, RootState } from '../store/store';
import { selectComment } from '../store/commentsSlice';

import { Spinner, Card } from "react-bootstrap";
import CommentCard from './CommentCard';

interface Props {
  navigate: NavigateFunction;
  handleShowDeleteModal?: () => void;
  token: string | undefined;
}

export default function DisplayComments({ navigate, handleShowDeleteModal, token } : Props) {
  const { comments, isLoading } = useSelector((state: RootState) => state.comment);
  const dispatch = useDispatch<AppDispatch>();

  if (isLoading) {
    return <div className="d-flex justify-content-center align-items-center" style={{ height: '100vh' }}>
      <Spinner animation="border" />
    </div>;
  }

  const handleUpdateComment = (index: number) => {
    dispatch(selectComment({index}));
    navigate(`/updateComment`);
  };

  const handleDeleteComment = (index: number) => {
    dispatch(selectComment({index}));
    handleShowDeleteModal?.();
  };

  return (
    <>
      {comments.length > 0 ? <div className="mt-4">
        {comments.map((comment, index) => {
          return <CommentCard key={index} 
            comment={comment} 
            index={index}
            updateCommentNavigation={handleUpdateComment}
            handleDeleteComment={() => handleDeleteComment?.(index)}
            token={token}
            />;
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
