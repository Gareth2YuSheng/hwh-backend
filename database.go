package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type PGStore struct {
	DB *sql.DB
}

func NewPGStore(dbURL string) (*PGStore, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PGStore{
		DB: db,
	}, nil
}

// INIT FUNCTIONS
func (s *PGStore) dbInit() error {
	log.Println("Running dbInit")
	if err := s.createUserTable(); err != nil {
		return err
	}

	return nil
}

func (s *PGStore) createUserTable() error {
	log.Println("Running createUserTable")
	query := `CREATE TABLE IF NOT EXISTS user (
		id UUID PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100),
		role VARCHAR(20) NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createTagTable() error {
	log.Println("Running createTagTable")
	query := `CREATE TABLE IF NOT EXISTS tag (
		id UUID PRIMARY KEY,
		username VARCHAR(100) NOT NULL,
		password VARCHAR(100),
		role VARCHAR(20) NOT NULL,
		created_at TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}
