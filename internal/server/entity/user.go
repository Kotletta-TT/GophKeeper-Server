package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}
