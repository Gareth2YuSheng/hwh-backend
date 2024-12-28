package main

import (
	"database/sql"
	"fmt"
	"strings"

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
	logInfo("Running dbInit")
	//DROP TABLES
	// if err := s.dropUserTable(); err != nil {
	// 	return err
	// }

	//CREATE TABLES
	if err := s.createUserTable(); err != nil {
		return err
	}
	// if err := s.createTagTable(); err != nil {
	// 	return err
	// }
	// if err := s.createThreadTable(); err != nil {
	// 	return err
	// }
	// if err := s.createCommentTable(); err != nil {
	// 	return err
	// }
	// if err := s.createVoteTable(); err != nil {
	// 	return err
	// }
	// if err := s.createVoteTable(); err != nil {
	// 	return err
	// }

	//SEED DATA
	s.seedUserTable()

	return nil
}

//SEED USER DATA

func (s *PGStore) seedUserTable() {
	admin, err := NewAdminUser("Robin Banks", "root")
	if err != nil {
		logError("Error Creating New Admin User Template", err)
	}
	fmt.Printf("New Admin User: %v\n", admin) //remove later
	err = s.CreateUser(admin)
	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		logError("unable to Create Admin User", err)
	}
}

//DROP TABLE FUNCTIONS

func (s *PGStore) dropUserTable() error {
	logInfo("Running dropUserTable")
	query := `DROP TABLE IF EXISTS users;`

	_, err := s.DB.Exec(query)
	return err
}

//CREATE TABLE FUNCTIONS

func (s *PGStore) createUserTable() error {
	logInfo("Running createUserTable")
	query := `CREATE TABLE IF NOT EXISTS users (
		userID UUID PRIMARY KEY,
		username VARCHAR(100) UNIQUE NOT NULL,
		password VARCHAR(150),
		role VARCHAR(20) NOT NULL,
		createdAt TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createTagTable() error {
	logInfo("Running createTagTable")
	query := `CREATE TABLE IF NOT EXISTS tags (
		tagID UUID PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createThreadTable() error {
	logInfo("Running createThreadTable")
	query := `CREATE TABLE IF NOT EXISTS threads (
		threadID UUID PRIMARY KEY,
		title VARCHAR(200) NOT NULL,
		content TEXT NOT NULL,
		authorID UUID NOT NULL REFERENCES users(userID) ON DELETE SET NULL,
		tagID UUID NOT NULL REFERENCES tags(tagID) ON DELETE RESTRICT,
		createdAt TIMESTAMP NOT NULL
		updatedAt TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createCommentTable() error {
	logInfo("Running createCommentTable")
	query := `CREATE TABLE IF NOT EXISTS comments (
		commentID UUID PRIMARY KEY,
		content TEXT NOT NULL,
		voteCount INTEGER NOT NULL,
		authorID UUID NOT NULL REFERENCES user(userID) ON DELETE SET NULL,
		threadID UUID NOT NULL REFERENCES thread(threadID) ON DELETE CASCADE,
		createdAt TIMESTAMP NOT NULL
		updatedAt TIMESTAMP NOT NULL
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createVoteTable() error {
	logInfo("Running createVoteTable")
	query := `CREATE TABLE IF NOT EXISTS votes (
		voteID UUID PRIMARY KEY,
		voteValue INTEGER NOT NULL,
		authorID UUID NOT NULL REFERENCES users(userID) ON DELETE SET NULL,
		commentID UUID NOT NULL REFERENCES threads(threadID) ON DELETE CASCADE,
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createImageTable() error {
	logInfo("Running createImageTable")
	query := `CREATE TABLE IF NOT EXISTS images (
		imageID UUID PRIMARY KEY,
		threadID UUID NOT NULL REFERENCES users(userID) ON DELETE CASCADE,
	);`

	_, err := s.DB.Exec(query)
	return err
}
