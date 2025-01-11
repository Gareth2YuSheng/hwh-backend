package main

import (
	"database/sql"
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
	ThreadID     uuid.UUID `json:"threadId"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CommentCount int       `json:"commentCount"`
	AuthorID     uuid.UUID `json:"authorId"`
	TagID        uuid.UUID `json:"tagId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ThreadDetails struct {
	ThreadID     uuid.UUID `json:"threadId"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CommentCount int       `json:"commentCount"`
	AuthorID     uuid.UUID `json:"authorId"`
	Author       string    `json:"author"`
	TagID        uuid.UUID `json:"tagId"`
	TagName      string    `json:"tagName"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type ThreadCondensed struct {
	ThreadID     uuid.UUID `json:"threadId"`
	Title        string    `json:"title"`
	CommentCount int       `json:"commentCount"`
	Author       string    `json:"author"`
	TagID        uuid.UUID `json:"tagId"`
	TagName      string    `json:"tagName"`
	CreatedAt    time.Time `json:"createdAt"`
}

// type ThreadTally struct {
// 	TallyID int       `json:"tallyId"`
// 	TagID   uuid.UUID `json:"tagId"`
// 	Count   int       `json:"count"`
// }

// type TotalThreadTally struct {
// 	Count int `json:"count"`
// }

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

type CommentWithAuthor struct {
	CommentID uuid.UUID `json:"commentId"`
	Content   string    `json:"content"`
	VoteCount int       `json:"voteCount"`
	AuthorID  uuid.UUID `json:"authorId"`
	Author    string    `json:"author"`
	ThreadID  uuid.UUID `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsAnswer  bool      `json:"isAnswer"`
}

type CommentWithVoteCondensed struct {
	CommentID uuid.UUID `json:"commentId"`
	Content   string    `json:"content"`
	VoteCount int       `json:"voteCount"`
	AuthorID  uuid.UUID `json:"authorId"`
	Author    string    `json:"author"`
	ThreadID  uuid.UUID `json:"threadId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	IsAnswer  bool      `json:"isAnswer"`
	Vote      *VoteCondensed
}

type Vote struct {
	VoteID    uuid.UUID `json:"voteId"`
	CommentID uuid.UUID `json:"commentId"`
	VoteValue int       `json:"voteValue"`
	AuthorID  uuid.UUID `json:"authorId"`
}

type VoteCondensed struct {
	VoteID    uuid.UUID `json:"voteId"`
	VoteValue int       `json:"voteValue"`
}

type VoteSQL struct {
	VoteID    uuid.UUID
	CommentID uuid.UUID
	VoteValue sql.NullInt32
	AuthorID  uuid.UUID
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
			CreatedAt: getTimeNow(),
		}, nil
	}
	return &User{
		UserID:    uuid.New(),
		Username:  username,
		Role:      "Standard",
		CreatedAt: getTimeNow(),
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
		CreatedAt: getTimeNow(),
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
		ThreadID:     uuid.New(),
		Title:        title,
		Content:      content,
		CommentCount: 0,
		AuthorID:     authorId,
		TagID:        tagId,
		CreatedAt:    getTimeNow(),
		UpdatedAt:    getTimeNow(),
	}, nil
}

func (t *Thread) UpdateThread(title, content string) error {
	logInfo("Running UpdateThread")
	if title == "" {
		return fmt.Errorf("title cannot be empty")
	}
	if content == "" {
		return fmt.Errorf("content cannot be empty")
	}
	t.Title = title
	t.Content = content
	t.UpdatedAt = getTimeNow()
	return nil
}

func NewComment(content string, threadID, authorId uuid.UUID) (*Comment, error) {
	logInfo("Running NewComment")
	if content == "" {
		return nil, fmt.Errorf("comment content cannot be empty")
	}
	if threadID == uuid.Nil {
		return nil, fmt.Errorf("thread threadID cannot be nil")
	}
	if authorId == uuid.Nil {
		return nil, fmt.Errorf("thread authorID cannot be nil")
	}
	return &Comment{
		CommentID: uuid.New(),
		Content:   content,
		VoteCount: 0,
		AuthorID:  authorId,
		ThreadID:  threadID,
		CreatedAt: getTimeNow(),
		UpdatedAt: getTimeNow(),
		IsAnswer:  false,
	}, nil
}

func (c *Comment) UpdateCommentContent(content string) error {
	logInfo("Running UpdateCommentContent")
	if content == "" {
		fmt.Println("Lig")
		return fmt.Errorf("content cannot be empty")
	}
	c.Content = content
	c.UpdatedAt = getTimeNow()
	return nil
}

func (c *Comment) UpdateCommentIsAnswer(isAnswer bool) error {
	logInfo("Running UpdateCommentIsAnswer")
	c.IsAnswer = isAnswer
	return nil
}

func (c *Comment) UpdateCommentVoteCount(amt int) error {
	logInfo("Running UpdateCommentContent")
	if amt == 0 {
		return fmt.Errorf("amt cannot be 0")
	}
	c.VoteCount += amt
	return nil
}

func NewVote(commentId, authorId uuid.UUID, voteVal int) (*Vote, error) {
	logInfo("Running NewVote")
	if voteVal != 1 && voteVal != -1 {
		return nil, fmt.Errorf("invalid vote value")
	}
	if commentId == uuid.Nil {
		return nil, fmt.Errorf("vote commentID cannot be nil")
	}
	if authorId == uuid.Nil {
		return nil, fmt.Errorf("vote authorID cannot be nil")
	}
	return &Vote{
		VoteID:    uuid.New(),
		VoteValue: voteVal,
		CommentID: commentId,
		AuthorID:  authorId,
	}, nil
}

func (v *Vote) UpdateVoteValue(voteVal int) error {
	logInfo("Running UpdateVoteValue")
	if voteVal != 1 && voteVal != -1 {
		return fmt.Errorf("invalid vote value: %d", voteVal)
	}
	v.VoteValue = voteVal
	return nil
}

// func NewThreadTally(tagId uuid.UUID) (*ThreadTally, error) {
// 	logInfo("Running NewThreadTally")
// 	if tagId == uuid.Nil {
// 		return nil, fmt.Errorf("tagID cannot be empty")
// 	}
// 	return &ThreadTally{
// 		TagID: tagId,
// 		Count: 0,
// 	}, nil
// }

// func (tt *ThreadTally) UpdateThreadTally(amt int) error {
// 	logInfo("Running UpdateThreadTally")
// 	if amt == 0 {
// 		return fmt.Errorf("amt cannot be 0")
// 	}
// 	tt.Count += amt
// 	return nil
// }
