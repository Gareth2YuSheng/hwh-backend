package main

import "github.com/google/uuid"

//REQUESTS
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateThreadRequest struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
	TagID   uuid.UUID `json:"tagId"`
}

type UpdateThreadRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateTagRequest struct {
	Name string `json:"name"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}

type MarkCommentAsAnswerRequest struct {
	IsAnswer bool `json:"isAnswer"`
}

type VoteCommentRequest struct {
	VoteType string `json:"voteType"`
}

//RESPONSES
type LoginResponse struct {
	UserID      uuid.UUID `json:"userId"`
	AccessToken string    `json:"accessToken"`
}

type GetUserDataResponse struct {
	User User `json:"user"`
}

type GetTagsResponse struct {
	Tags []*Tag `json:"tags"`
}

type GetThreadsResponse struct {
	ThreadCount int                `json:"threadCount"`
	Threads     []*ThreadCondensed `json:"threads"`
}

type GetThreadDetailsResponse struct {
	Thread *Thread `json:"thread"`
	// Commends []*Comment `json:"comments"`
}

type GetCommentsResponse struct {
	CommentCount int        `json:"commentCount"`
	Comments     []*Comment `json:"comments"`
}

type GetCommentsWithVoteResponse struct {
	CommentCount int                         `json:"commentCount"`
	Comments     []*CommentWithVoteCondensed `json:"comments"`
}

//OTHERS
type GetThread struct {
	Thread *ThreadCondensed
	Count  int
}

type GetComment struct {
	Comment *Comment
	Count   int
}

type GetCommentWithVoteSQL struct {
	Comment *Comment
	Vote    *VoteSQL
	Count   int
}
