package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"fmt"

	"github.com/google/uuid"
)

type List interface {
	ListSecretByName(secretName string) ([]*entity.SecretCard, error)
	ListSecretByUserId(userId uuid.UUID) ([]*entity.SecretCard, error)
}

type ListRepository interface {
	ListSecretCardByName(secretName string) ([]*entity.SecretCard, error)
	ListSecretCardByUserId(userId uuid.UUID) ([]*entity.SecretCard, error)
}

type ListUC struct {
	l logger.Logger
	r ListRepository
}

func NewListUC(l logger.Logger, r ListRepository) *ListUC {
	return &ListUC{l: l, r: r}
}

func (uc *ListUC) ListSecretByName(secretName string) ([]*entity.SecretCard, error) {
	cards, err := uc.r.ListSecretCardByName(secretName)
	if err != nil {
		uc.l.Error(
			fmt.Sprintf("get list secret card by name %s err: %s",
				secretName,
				err.Error()),
		)
		return nil, err
	}
	return cards, err
}

func (uc *ListUC) ListSecretByUserId(userId uuid.UUID) ([]*entity.SecretCard, error) {
	cards, err := uc.r.ListSecretCardByUserId(userId)
	if err != nil {
		uc.l.Error(
			fmt.Sprintf("get list secret card by user id %s err: %s",
				userId.String(),
				err.Error()),
		)
		return nil, err
	}
	return cards, err
}
