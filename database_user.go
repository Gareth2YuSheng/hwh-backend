package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func (s *PGStore) CreateUser(user *User) error {
	logInfo("Running CreateUser")
	query := `INSERT INTO users 
	(userID, username, password, role, createdAt) 
	values ($1, $2, $3, $4, $5)`
	_, err := s.DB.Query(query,
		user.UserID,
		user.Username,
		user.Password,
		user.Role,
		user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *PGStore) GetUserByUserID(userID uuid.UUID) (*User, error) {
	logInfo("Running GetUserByUserID")
	query := `SELECT * FROM users WHERE userID = $1 LIMIT 1`
	rows, err := s.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, fmt.Errorf("user [%s] not found", userID.String())
}

func (s *PGStore) GetUserByUsername(username string) (*User, error) {
	logInfo("Running GetUserByUsername")
	query := `SELECT * FROM users WHERE username = $1 LIMIT 1`
	rows, err := s.DB.Query(query, username)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		return scanIntoUser(rows)
	}
	return nil, fmt.Errorf("user [%s] not found", username)
}

func scanIntoUser(rows *sql.Rows) (*User, error) {
	user := new(User)
	err := rows.Scan(
		&user.UserID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.CreatedAt)
	return user, err
}
