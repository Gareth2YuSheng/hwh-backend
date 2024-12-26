package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"userId"`
	Username  string    `json:"username"`
	Password  string    `json:"-"` //does not show in the json
	Role      string    `json:"-"` //does not show in the json
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
