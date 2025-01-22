package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateThread(thread *Thread) error {
	logInfo("Running: Database - CreateThread")
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
	logInfo("Running: Database - GetAllThreads")
	query := `SELECT threadID, title, commentCount, users.username, threads.tagID, tags.name, threads.createdAt, count(*) OVER() AS totalCount 
	FROM threads, users, tags WHERE threads.authorId = users.userId AND tags.tagId = threads.tagId`
	if search != "" {
		query += ` AND title ILIKE '%` + search + `%'`
	}
	if tagID != uuid.Nil {
		query += fmt.Sprintf(` AND threads.tagID = '%v'`, tagID)
	}
	query += fmt.Sprintf(` ORDER BY createdAt DESC OFFSET %d LIMIT %d;`, (page-1)*count, count)

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
	logInfo("Running: Database - GetThreadByThreadID")
	query := `SELECT * FROM threads WHERE threadID = $1;`
	rows, err := s.DB.Query(query, threadId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoThread(rows)
	}
	return nil, fmt.Errorf("thread [%s] not found", threadId.String())
}

func (s *PGStore) GetThreadDetailsByThreadID(threadId uuid.UUID) (*ThreadDetails, error) {
	logInfo("Running: Database - GetThreadDetailsByThreadID")
	query := `SELECT threads.threadId, title, content, commentCount, authorId, username, threads.tagId, 
	 name, threads.createdAt, threads.updatedAt, images.cloudinaryURL
	 FROM threads 
	 JOIN tags ON threads.tagId = tags.tagId 
	 JOIN users ON threads.authorId = users.userId 
	 LEFT JOIN images ON images.threadId = threads.threadId 
	 WHERE threads.threadID = $1;`
	rows, err := s.DB.Query(query, threadId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoThreadDetails(rows)
	}
	return nil, fmt.Errorf("thread [%s] not found", threadId.String())
}

func (s *PGStore) UpdateThread(thread *Thread) error {
	logInfo("Running: Database - UpdateThread")
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
	logInfo("Running: Database - UpdateThreadCommentCountByThreadID")
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
	logInfo("Running: Database - DeleteThreadByThreadID")
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

func scanIntoThreadDetails(rows *sql.Rows) (*ThreadDetails, error) {
	thread := new(ThreadDetails)
	err := rows.Scan(
		&thread.ThreadID,
		&thread.Title,
		&thread.Content,
		&thread.CommentCount,
		&thread.AuthorID,
		&thread.Author,
		&thread.TagID,
		&thread.TagName,
		&thread.CreatedAt,
		&thread.UpdatedAt,
		&thread.ImageURLNullable)
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
		&thread.Author,
		&thread.TagID,
		&thread.TagName,
		&thread.CreatedAt,
		&getThread.Count)
	return getThread, err
}
