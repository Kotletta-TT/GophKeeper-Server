package entity

import "github.com/google/uuid"

type SecretCard struct {
	ID       uuid.UUID
	UserId   uuid.UUID
	Name     string
	URL      string
	Login    string
	Password string
	Text     string
	Files    map[string]string
	Meta     map[string]string
}

// type JSONSecretCard struct {
// 	Name     string
// 	URL      url.URL
// 	Login    string
// 	Password string
// 	Text     string
// 	File     []byte
// 	Meta     map[string]string
// }
