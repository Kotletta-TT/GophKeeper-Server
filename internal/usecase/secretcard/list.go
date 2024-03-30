package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type List interface {
	ListSecretByName(ctx context.Context, secretName string) ([]*entity.SecretCard, error)
	ListSecretByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.SecretCard, error)
}

type ListRepository interface {
	ListSecretCardByName(ctx context.Context, secretName string) ([]*entity.SecretCard, error)
	ListSecretCardByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.SecretCard, error)
}

type ListUC struct {
	l logger.Logger
	r ListRepository
}

func NewListUC(l logger.Logger, r ListRepository) *ListUC {
	return &ListUC{l: l, r: r}
}

func (uc *ListUC) ListSecretByName(ctx context.Context, secretName string) ([]*entity.SecretCard, error) {
	cards, err := uc.r.ListSecretCardByName(ctx, secretName)
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

func (uc *ListUC) ListSecretByUserId(ctx context.Context, userId uuid.UUID) ([]*entity.SecretCard, error) {
	cards, err := uc.r.ListSecretCardByUserId(ctx, userId)
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
