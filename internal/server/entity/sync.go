package entity

import (
	"github.com/google/uuid"
)

type SyncSecretCard struct {
	Card  *SecretCard
	Files []*FileSecretCard
	Metas []*MetaSecretCard
}

type Sync struct {
	UserId uuid.UUID
	Cards  []*SyncSecretCard
}
