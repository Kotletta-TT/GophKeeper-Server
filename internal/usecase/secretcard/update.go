package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
	"fmt"
)

type Update interface {
	UpdateSecret(card *entity.SecretCard) error
}

type UpdateRepo interface {
	UpdateSecretCard(card *entity.SecretCard) error
}

type UpdateUC struct {
	l logger.Logger
	r UpdateRepo
}

func NewUpdateUC(l logger.Logger, r UpdateRepo) *UpdateUC {
	return &UpdateUC{l: l, r: r}
}

func (uc *UpdateUC) UpdateSecret(card *entity.SecretCard) error {
	err := uc.r.UpdateSecretCard(card)
	if err != nil {
		uc.l.Errorf(
			fmt.Sprintf(
				"update secret card %s for user %s err: %s",
				card.Name,
				card.UserId,
				err.Error(),
			))
		return err
	}
	return nil
}
