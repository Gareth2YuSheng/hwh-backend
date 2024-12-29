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

type LoginResponse struct {
	UserID      uuid.UUID `json:"userId"`
	Username    string    `json:"username"`
	Role        string    `json:"role"`
	AccessToken string    `json:"accessToken"`
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

type GetTagsResponse struct {
	Tags []*Tag `json:"tags"`
}

type GetThreadsResponse struct {
	Threads []*Thread `json:"threads"`
}
