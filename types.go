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
