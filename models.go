package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` //does not show in the json
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type Thread struct {
	ThreadID  uuid.UUID `json:"threadId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  uuid.UUID `json:"authorId"`
	TagID     uuid.UUID `json:"tagId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Comment struct {
	CommentID uuid.UUID `json:"commentId"`
	Content   string    `json:"content"`
	VoteCount int       `json:"voteCount"`
	AuthorID  uuid.UUID `json:"authorId"`
	ThreadID  uuid.UUID `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Vote struct {
	VoteID    uuid.UUID `json:"voteId"`
	CommentID uuid.UUID `json:"commentId"`
	VoteValue int       `json:"voteCount"`
	AuthorID  uuid.UUID `json:"authorId"`
}

type Tag struct {
	TagID uuid.UUID `json:"tagId"`
	Name  string    `json:"name"`
}

type Image struct {
	ImageID  uuid.UUID `json:"imageId"`
	ThreadID uuid.UUID `json:"threaId"`
}

func NewStandardUser(username, password string) (*User, error) {
	logInfo("Running NewStandardUser")
	encryptedPwd, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	if password != "" {
		return &User{
			UserID:    uuid.New(),
			Username:  username,
			Password:  string(encryptedPwd),
			Role:      "Standard",
			CreatedAt: time.Now().Local().UTC(),
		}, nil
	}
	return &User{
		UserID:    uuid.New(),
		Username:  username,
		Role:      "Standard",
		CreatedAt: time.Now().Local().UTC(),
	}, nil

}

func NewAdminUser(username, password string) (*User, error) {
	logInfo("Running NewAdminUser")
	if password == "" {
		return nil, fmt.Errorf("password cannot be empty for admin user")
	}
	encryptedPwd, err := GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	return &User{
		UserID:    uuid.New(),
		Username:  username,
		Password:  string(encryptedPwd),
		Role:      "Admin",
		CreatedAt: time.Now().Local().UTC(),
	}, nil
}

func NewTag(name string) (*Tag, error) {
	logInfo("Running NewTag")
	return &Tag{
		TagID: uuid.New(),
		Name:  name,
	}, nil
}
