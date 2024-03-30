package secretcard

import (
	"GophKeeper-Server/logger"
	"fmt"

	"github.com/google/uuid"
)

type Delete interface {
	DeleteSecret(secretId uuid.UUID) error
}

type DeleteRepository interface {
	DeleteSecretCard(secretId uuid.UUID) error
}

type DeleteUC struct {
	l    logger.Logger
	repo DeleteRepository
}

func NewDeleteUC(l logger.Logger, repo DeleteRepository) *DeleteUC {
	return &DeleteUC{l, repo}
}

func (uc *DeleteUC) DeleteSecret(secretId uuid.UUID) error {
	err := uc.repo.DeleteSecretCard(secretId)
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
