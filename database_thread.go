package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateThread(thread *Thread) error {
	logInfo("Running CreateThread")
	query := `INSERT INTO threads 
	(threadID, title, content, commentCount, authorID, tagID, createdAt, updatedAt) 
	values ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.DB.Query(query,
		thread.ThreadID,
		thread.Title,
		thread.Content,
		thread.CommentCount,
		thread.AuthorID,
		thread.TagID,
		thread.CreatedAt,
		thread.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetAllThreads(count, page int, search string, tagID uuid.UUID) ([]*ThreadCondensed, int, error) {
	logInfo("Running GetAllThreads")
	query := `SELECT threadID, title, commentCount, authorID, tagID, createdAt, count(*) OVER() AS totalCount FROM threads`
	if search != "" || tagID != uuid.Nil {
		query += ` WHERE`
	}
	if search != "" {
		query += ` title ILIKE '%` + search + `%'`
	}
	if search != "" && tagID != uuid.Nil {
		query += ` AND`
	}
	if tagID != uuid.Nil {
		query += fmt.Sprintf(` tagID = '%v'`, tagID)
	}
	query += fmt.Sprintf(` ORDER BY createdAt DESC OFFSET %d LIMIT %d;`, (page-1)*count, count)

	fmt.Println(query) //remove later
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, 0, err
	}
	threads := []*ThreadCondensed{}
	threadCount := 0
	for rows.Next() {
		getThread, err := scanIntoGetThreads(rows)
		if err != nil {
			return nil, 0, err
		}
		threads = append(threads, getThread.Thread)
		threadCount = getThread.Count
	}
	return threads, threadCount, nil
}

func (s *PGStore) GetThreadByThreadID(threadId uuid.UUID) (*Thread, error) {
	logInfo("Running GetThreadByThreadID")
	query := `SELECT * FROM threads WHERE threadID = $1;`
	rows, err := s.DB.Query(query, threadId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoThread(rows)
	}
	return nil, fmt.Errorf("thread [%v] not found", threadId)
}

func (s *PGStore) UpdateThread(thread *Thread) error {
	logInfo("Running UpdateThread")
	query := `UPDATE threads
	SET title = $1, content = $2, updatedAt = $3
	WHERE threadID = $4`
	_, err := s.DB.Query(query,
		thread.Title,
		thread.Content,
		thread.UpdatedAt,
		thread.ThreadID)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) UpdateThreadCommentCountByThreadID(threadId uuid.UUID, amt int) error {
	logInfo("Running UpdateThreadCommentCountByThreadID")
	query := `UPDATE threads
	SET commentCount = commentCount + $1
	WHERE threadID = $2;`
	_, err := s.DB.Query(query, amt, threadId)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) DeleteThreadByThreadID(threadId uuid.UUID) error {
	logInfo("Running DeleteThreadByThreadID")
	query := `DELETE FROM threads
	WHERE threadID = $1;`
	_, err := s.DB.Query(query, threadId)
	if err != nil {
		return err
	}
	return nil
}

func scanIntoThread(rows *sql.Rows) (*Thread, error) {
	thread := new(Thread)
	err := rows.Scan(
		&thread.ThreadID,
		&thread.Title,
		&thread.Content,
		&thread.CommentCount,
		&thread.AuthorID,
		&thread.TagID,
		&thread.CreatedAt,
		&thread.UpdatedAt)
	return thread, err
}

func scanIntoGetThreads(rows *sql.Rows) (*GetThread, error) {
	getThread := new(GetThread)
	thread := new(ThreadCondensed)
	getThread.Thread = thread
	err := rows.Scan(
		&thread.ThreadID,
		&thread.Title,
		&thread.CommentCount,
		&thread.AuthorID,
		&thread.TagID,
		&thread.CreatedAt,
		&getThread.Count)
	return getThread, err
}
