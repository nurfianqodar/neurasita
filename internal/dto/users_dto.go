package dto

import (
	"time"

	"github.com/google/uuid"
)

type SingleUserData struct {
	User any `json:"user"`
}

type ManyUserData struct {
	Users any `json:"users"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserReponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}
