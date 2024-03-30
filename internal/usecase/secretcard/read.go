package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Read interface {
	ReadSecret(ctx context.Context, secretId uuid.UUID) (*entity.SecretCard, error)
}

type ReadRepository interface {
	GetSecretCardByID(ctx context.Context, secretId uuid.UUID) (*entity.SecretCard, error)
}

type ReadUC struct {
	l    logger.Logger
	repo ReadRepository
}

func NewReadUC(l logger.Logger, repo ReadRepository) *ReadUC {
	return &ReadUC{l: l, repo: repo}
}

func (uc *ReadUC) ReadSecret(ctx context.Context, secretId uuid.UUID) (*entity.SecretCard, error) {
	secret, err := uc.repo.GetSecretCardByID(ctx, secretId)
	if err != nil {
		uc.l.Error(
			fmt.Sprintf(
				"error while reading secret with id %s err: %s",
				secretId.String(),
				err.Error(),
			))
		return nil, err
	}
	return secret, nil
}
