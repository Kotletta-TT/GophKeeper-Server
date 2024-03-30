package secretcard

import (
	"GophKeeper-Server/logger"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Delete interface {
	DeleteSecret(ctx context.Context, secretId uuid.UUID) error
}

type DeleteRepository interface {
	DeleteSecretCard(ctx context.Context, secretId uuid.UUID) error
}

type DeleteUC struct {
	l    logger.Logger
	repo DeleteRepository
}

func NewDeleteUC(l logger.Logger, repo DeleteRepository) *DeleteUC {
	return &DeleteUC{l, repo}
}

func (uc *DeleteUC) DeleteSecret(ctx context.Context, secretId uuid.UUID) error {
	err := uc.repo.DeleteSecretCard(ctx, secretId)
	if err != nil {
		uc.l.Error(
			fmt.Sprintf("failed to delete secret %s err: %s",
				secretId.String(),
				err.Error(),
			))
		return err
	}
	uc.l.Info(fmt.Sprintf("secret %s deleted", secretId.String()))
	return nil
}
