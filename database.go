package main

import (
	"database/sql"
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
	// s.dropAllTables()

	//CREATE TABLES
	s.createAllTables()

	//SEED DATA
	s.seedData()

	return nil
}

//SEED DATA

func (s *PGStore) seedData() {
	s.seedUserTable()
	s.seedTagTable()
}

func (s *PGStore) seedUserTable() {
	admin, err := NewAdminUser("Robin Banks", "root")
	if err != nil {
		logError("Error Creating New Admin User Template", err)
	}
	err = s.CreateUser(admin)
	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		logError("unable to Create Admin User", err)
	}
}

func (s *PGStore) seedTagTable() {
	tagEng, err := NewTag("English")
	if err != nil {
		logError("Error Creating New Tag Template - English", err)
	}
	tagMath, err := NewTag("Math")
	if err != nil {
		logError("Error Creating New Tag Template - Math", err)
	}
	tagScience, err := NewTag("Science")
	if err != nil {
		logError("Error Creating New Tag Template - Science", err)
	}
	err = s.CreateTag(tagEng)
	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		logError("unable to Create Default English Tag", err)
	}
	err = s.CreateTag(tagMath)
	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		logError("unable to Create Default Math Tag", err)
	}
	err = s.CreateTag(tagScience)
	if err != nil && !strings.Contains(err.Error(), "duplicate key") {
		logError("unable to Create Default Science Tag", err)
	}
}

//DROP TABLE FUNCTIONS

func (s *PGStore) dropAllTables() error {
	if err := s.dropImageTable(); err != nil {
		return err
	}
	if err := s.dropVoteTable(); err != nil {
		return err
	}
	if err := s.dropCommentTable(); err != nil {
		return err
	}
	if err := s.dropThreadTable(); err != nil {
		return err
	}
	if err := s.dropTagTable(); err != nil {
		return err
	}
	if err := s.dropUserTable(); err != nil {
		return err
	}
	return nil
}

func (s *PGStore) dropUserTable() error {
	logInfo("Running dropUserTable")
	query := `DROP TABLE IF EXISTS users;`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) dropTagTable() error {
	logInfo("Running dropUserTable")
	query := `DROP TABLE IF EXISTS tags;`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) dropThreadTable() error {
	logInfo("Running dropThreadTable")
	query := `DROP TABLE IF EXISTS threads;`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) dropCommentTable() error {
	logInfo("Running dropCommentTable")
	query := `DROP TABLE IF EXISTS comments;`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) dropVoteTable() error {
	logInfo("Running dropVoteTable")
	query := `DROP TABLE IF EXISTS votes;`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) dropImageTable() error {
	logInfo("Running dropImageTable")
	query := `DROP TABLE IF EXISTS images;`

	_, err := s.DB.Exec(query)
	return err
}

//CREATE TABLE FUNCTIONS

func (s *PGStore) createAllTables() error {
	if err := s.createUserTable(); err != nil {
		return err
	}
	if err := s.createTagTable(); err != nil {
		return err
	}
	if err := s.createThreadTable(); err != nil {
		return err
	}
	if err := s.createCommentTable(); err != nil {
		return err
	}
	if err := s.createVoteTable(); err != nil {
		return err
	}
	if err := s.createImageTable(); err != nil {
		return err
	}
	return nil
}

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
		name VARCHAR(100) UNIQUE NOT NULL
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
		createdAt TIMESTAMP NOT NULL,
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
		authorID UUID NOT NULL REFERENCES users(userID) ON DELETE SET NULL,
		threadID UUID NOT NULL REFERENCES threads(threadID) ON DELETE CASCADE,
		createdAt TIMESTAMP NOT NULL,
		updatedAt TIMESTAMP NOT NULL,
		isAnswer BOOLEAN NOT NULL
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
		commentID UUID NOT NULL REFERENCES threads(threadID) ON DELETE CASCADE
	);`

	_, err := s.DB.Exec(query)
	return err
}

func (s *PGStore) createImageTable() error {
	logInfo("Running createImageTable")
	query := `CREATE TABLE IF NOT EXISTS images (
		imageID UUID PRIMARY KEY,
		threadID UUID NOT NULL REFERENCES users(userID) ON DELETE CASCADE
	);`

	_, err := s.DB.Exec(query)
	return err
}
