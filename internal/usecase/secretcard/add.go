package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"context"
	"fmt"
)

type Add interface {
	AddSecret(ctx context.Context, card *entity.SecretCard) error
}

type AddRepository interface {
	CreateSecretCard(ctx context.Context, card *entity.SecretCard) error
}

type AddUC struct {
	l logger.Logger
	r AddRepository
}

func NewAddUC(l logger.Logger, r AddRepository) *AddUC {
	return &AddUC{l: l, r: r}
}

func (uc *AddUC) AddSecret(ctx context.Context, card *entity.SecretCard) error {
	err := uc.r.CreateSecretCard(ctx, card)
	if err != nil {
		uc.l.Errorf(
			fmt.Sprintf(
				"store secret card %s for user %s err: %s",
				card.Name,
				card.UserId,
				err.Error()))
		return err
	}
	uc.l.Infof(
		fmt.Sprintf(
			"secret card %s for user %s stored",
			card.ID,
			card.UserId))
	return nil
}
