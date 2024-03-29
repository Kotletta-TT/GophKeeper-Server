package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/logger"
)

type Add interface {
	AddSecret(card *entity.SecretCard) error 
}

type AddRepository interface {
	StoreSecretCard(card *entity.SecretCard)
}

type AddUC struct {
	l logger.Logger
	r 
}