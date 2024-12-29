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

type ThreadCondensed struct {
	ThreadID  uuid.UUID `json:"threadId"`
	Title     string    `json:"title"`
	AuthorID  uuid.UUID `json:"authorId"`
	TagID     uuid.UUID `json:"tagId"`
	CreatedAt time.Time `json:"createdAt"`
}

type Comment struct {
	CommentID uuid.UUID `json:"commentId"`
	Content   string    `json:"content"`
	VoteCount int       `json:"voteCount"`
	AuthorID  uuid.UUID `json:"authorId"`
	ThreadID  uuid.UUID `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsAnswer  bool      `json:"isAnswer"`
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
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
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
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
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
	if name == "" {
		return nil, fmt.Errorf("tag name cannot be empty")
	}
	return &Tag{
		TagID: uuid.New(),
		Name:  name,
	}, nil
}

func NewThread(title, content string, authorId, tagId uuid.UUID) (*Thread, error) {
	logInfo("Running NewThread")
	if title == "" {
		return nil, fmt.Errorf("thread title cannot be empty")
	}
	if content == "" {
		return nil, fmt.Errorf("thread content cannot be empty")
	}
	if authorId == uuid.Nil {
		return nil, fmt.Errorf("thread authorID cannot be nil")
	}
	if tagId == uuid.Nil {
		return nil, fmt.Errorf("thread tagID cannot be nil")
	}
	return &Thread{
		ThreadID:  uuid.New(),
		Title:     title,
		Content:   content,
		AuthorID:  authorId,
		TagID:     tagId,
		CreatedAt: time.Now().Local().UTC(),
		UpdatedAt: time.Now().Local().UTC(),
	}, nil
}

func (t *Thread) UpdateThread(title, content string) {
	t.Title = title
	t.Content = content
	t.UpdatedAt = time.Now().Local().UTC()
}
