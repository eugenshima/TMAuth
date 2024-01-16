package model

import "github.com/google/uuid"

type FullAuthModel struct {
	ID           uuid.UUID `json:"id"`
	Login        string    `json:"login"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
}

type ProfileLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
