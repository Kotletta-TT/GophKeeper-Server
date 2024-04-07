package entity

import (
	"time"

	"github.com/google/uuid"
)

type SecretCard struct {
	ID         uuid.UUID
	UserId     uuid.UUID
	Name       string
	URL        string
	Login      string
	Password   string
	Text       string
	UpdateTime time.Time
}

type FileSecretCard struct {
	ID     uuid.UUID
	CardId uuid.UUID
	File   []byte
}

type MetaSecretCard struct {
	ID     uuid.UUID
	CardId uuid.UUID
	Key    string
	Value  string
}
