package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateThread(thread *Thread) error {
	logInfo("Running CreateThread")
	query := `INSERT INTO threads 
	(threadID, title, content, authorID, tagID, createdAt, updatedAt) 
	values ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.DB.Query(query,
		thread.ThreadID,
		thread.Title,
		thread.Content,
		thread.AuthorID,
		thread.TagID,
		thread.CreatedAt,
		thread.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetAllThreads(count, page int, search string, tagID uuid.UUID) ([]*Thread, error) {
	logInfo("Running GetAllThreads")
	query := `SELECT * FROM threads`
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
	query += fmt.Sprintf(` ORDER BY createdAt DESC OFFSET %d LIMIT %d`, (page-1)*count, count)
	fmt.Println(query) //remove later
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	threads := []*Thread{}
	for rows.Next() {
		thread, err := scanIntoThread(rows)
		if err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func (s *PGStore) GetThreadByThreadID(threadID uuid.UUID) (*Thread, error) {
	logInfo("Running GetThreadByThreadID")
	query := `SELECT * FROM threads WHERE threadID = $1 LIMIT 1`
	rows, err := s.DB.Query(query, threadID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoThread(rows)
	}
	return nil, fmt.Errorf("thread [%v] not found", threadID)
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

func (s *PGStore) DeleteThreadByThreadID(threadID uuid.UUID) error {
	logInfo("Running DeleteThreadByThreadID")
	query := `DELETE FROM threads
	WHERE threadID = $1`
	_, err := s.DB.Query(query, threadID)
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
		&thread.AuthorID,
		&thread.TagID,
		&thread.CreatedAt,
		&thread.UpdatedAt)
	return thread, err
}
