package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateComment(comment *Comment) error {
	logInfo("Running: Database - CreateComment")
	query := `INSERT INTO comments 
	(commentID, content, voteCount, authorID, threadID, createdAt, updatedAt, isAnswer) 
	values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.DB.Query(query,
		comment.CommentID,
		comment.Content,
		comment.VoteCount,
		comment.AuthorID,
		comment.ThreadID,
		comment.CreatedAt,
		comment.UpdatedAt,
		comment.IsAnswer)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetAllCommentsByThreadIDWithVotesByUserID(count, page int, threadId, userId uuid.UUID) ([]*CommentWithVoteCondensed, int, error) {
	logInfo("Running: Database - GetAllCommentsByThreadIDWithVotesByUserID")
	query := `SELECT c.commentID, c.content, c.voteCount, 
	c.authorID, u.username, c.threadID, c.createdAt, c.updatedAt, 
	c.isAnswer, v.voteID, v.voteValue, 
	count(*) OVER() AS totalCount FROM comments AS c JOIN users AS u ON c.authorId = u.userId
	LEFT JOIN votes AS v ON c.commentID = v.commentID AND v.authorId = $1
	WHERE u.userId = c.authorId AND c.threadID = $2
	ORDER BY c.isAnswer DESC, c.voteCount DESC, c.createdAt ASC 
	OFFSET $3 LIMIT $4`

	rows, err := s.DB.Query(query, userId, threadId, (page-1)*count, count)
	if err != nil {
		return nil, 0, err
	}

	comments := []*CommentWithVoteCondensed{}
	commentCount := 0
	for rows.Next() {
		getCommentWithVoteSQL, err := scanIntoGetCommentsWithVotes(rows)
		if err != nil {
			return nil, 0, err
		}
		getComment := getCommentWithVoteSQL.Comment
		getVote := getCommentWithVoteSQL.Vote
		commentWithVote := &CommentWithVoteCondensed{
			CommentID: getComment.CommentID,
			Content:   getComment.Content,
			VoteCount: getComment.VoteCount,
			AuthorID:  getComment.AuthorID,
			Author:    getComment.Author,
			ThreadID:  getComment.ThreadID,
			CreatedAt: getComment.CreatedAt,
			UpdatedAt: getComment.UpdatedAt,
			IsAnswer:  getComment.IsAnswer,
		}
		if getVote.VoteID != uuid.Nil {
			voteVal := 0
			if getVote.VoteValue.Valid {
				voteVal = int(getVote.VoteValue.Int32)
			}
			commentWithVote.Vote = &VoteCondensed{
				VoteID:    getVote.VoteID,
				VoteValue: voteVal,
			}
		}
		comments = append(comments, commentWithVote)
		commentCount = getCommentWithVoteSQL.Count
	}
	return comments, commentCount, nil
}

func (s *PGStore) GetAllCommentsByThreadID(count, page int, threadId uuid.UUID) ([]*Comment, int, error) {
	logInfo("Running: Database - GetAllCommentsByThreadID")
	query := `SELECT *, count(*) OVER() AS totalCount FROM comments 
	WHERE threadID = $1
	ORDER BY createdAt DESC OFFSET $2 LIMIT $3;`

	rows, err := s.DB.Query(query, threadId, (page-1)*count, count)
	if err != nil {
		return nil, 0, err
	}
	comments := []*Comment{}
	commentCount := 0
	for rows.Next() {
		getComment, err := scanIntoGetComments(rows)
		if err != nil {
			return nil, 0, err
		}
		comments = append(comments, getComment.Comment)
		commentCount = getComment.Count
	}
	return comments, commentCount, nil
}

func (s *PGStore) GetCommentByCommentID(commentId uuid.UUID) (*Comment, error) {
	logInfo("Running: Database - GetCommentByCommentID")
	query := `SELECT * FROM comments WHERE commentID = $1;`

	rows, err := s.DB.Query(query, commentId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoComment(rows)
	}
	return nil, fmt.Errorf("comment [%s] not found", commentId.String())
}

func (s *PGStore) UpdateCommentContent(comment *Comment) error {
	logInfo("Running: Database - UpdateCommentContent")
	query := `UPDATE comments
	SET content = $1, updatedAt = $2
	WHERE commentID = $3`
	_, err := s.DB.Query(query,
		comment.Content,
		comment.UpdatedAt,
		comment.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) UpdateCommentVoteCountByCommentID(commentId uuid.UUID, amt int) error {
	logInfo("Running: Database - UpdateCommentVoteCountByCommentID")
	query := `UPDATE comments
	SET voteCount = voteCount + $1
	WHERE commentID = $2`
	_, err := s.DB.Query(query, amt, commentId)
	if err != nil {
		return err
	}
	return nil
}

// func (s *PGStore) UpdateCommentVoteCount(comment *Comment) error {
// 	logInfo("Running UpdateCommentVoteCount")
// 	query := `UPDATE comments
// 	SET voteCount = voteCount + $1
// 	WHERE commentID = $2`
// 	_, err := s.DB.Query(query, amt, commentId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *PGStore) UpdateCommentIsAnswer(comment *Comment) error {
	logInfo("Running: Database - UpdateCommentIsAnswer")
	query := `UPDATE comments
	SET isAnswer = $1
	WHERE commentID = $2`
	_, err := s.DB.Query(query, comment.IsAnswer, comment.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) DeleteCommentByCommentID(commentId uuid.UUID) error {
	logInfo("Running: Database - DeleteCommentByCommentID")
	query := `DELETE FROM comments
	WHERE commentID = $1`
	_, err := s.DB.Query(query, commentId)
	if err != nil {
		return err
	}
	return nil
}

func scanIntoComment(rows *sql.Rows) (*Comment, error) {
	comment := new(Comment)
	err := rows.Scan(
		&comment.CommentID,
		&comment.Content,
		&comment.VoteCount,
		&comment.AuthorID,
		&comment.ThreadID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.IsAnswer)
	return comment, err
}

func scanIntoGetComments(rows *sql.Rows) (*GetComment, error) {
	getComment := new(GetComment)
	comment := new(Comment)
	getComment.Comment = comment
	err := rows.Scan(
		&comment.CommentID,
		&comment.Content,
		&comment.VoteCount,
		&comment.AuthorID,
		&comment.ThreadID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.IsAnswer,
		&getComment.Count)
	return getComment, err
}

func scanIntoGetCommentsWithVotes(rows *sql.Rows) (*GetCommentWithVoteSQL, error) {
	getComment := new(GetCommentWithVoteSQL)
	comment := new(CommentWithAuthor)
	vote := new(VoteSQL)
	getComment.Comment = comment
	getComment.Vote = vote
	err := rows.Scan(
		&comment.CommentID,
		&comment.Content,
		&comment.VoteCount,
		&comment.AuthorID,
		&comment.Author,
		&comment.ThreadID,
		&comment.CreatedAt,
		&comment.UpdatedAt,
		&comment.IsAnswer,
		&vote.VoteID,
		&vote.VoteValue,
		&getComment.Count)
	return getComment, err
}
