package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateVote(vote *Vote) error {
	logInfo("Running: Database - CreateVote")
	query := `INSERT INTO votes 
	(voteID, voteValue, authorID, commentID) 
	values ($1, $2, $3, $4);`
	_, err := s.DB.Query(query,
		vote.VoteID,
		vote.VoteValue,
		vote.AuthorID,
		vote.CommentID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetVotesForCommentByUser(commentId, authorId uuid.UUID) (*Vote, error) {
	logInfo("Running: Database - GetVotesForCommentByUser")
	query := `SELECT * FROM votes
	WHERE commentID = $1 AND authorID = $2;`
	rows, err := s.DB.Query(query, commentId, authorId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoVote(rows)
	}
	return nil, fmt.Errorf("vote for comment [%s] by user [%s] not found", commentId.String(), authorId.String())
}

func (s *PGStore) UpdateVoteVoteValue(vote *Vote) error {
	logInfo("Running: Database - UpdateVoteVoteValue")
	query := `UPDATE votes 
	SET voteValue = $1
	WHERE voteID = $2;`
	_, err := s.DB.Query(query, vote.VoteValue, vote.VoteID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) DeleteVote(vote *Vote) error {
	logInfo("Running: Database - DeleteVoteByVoteID")
	query := `DELETE FROM votes
	WHERE voteID = $1;`
	_, err := s.DB.Query(query, vote.VoteID)
	if err != nil {
		return err
	}
	return nil
}

func scanIntoVote(rows *sql.Rows) (*Vote, error) {
	vote := new(Vote)
	err := rows.Scan(
		&vote.VoteID,
		&vote.VoteValue,
		&vote.AuthorID,
		&vote.CommentID)
	return vote, err
}
