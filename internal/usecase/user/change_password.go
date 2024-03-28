package user

import "GophKeeper-Server/logger"

type ChangePassword interface {
	ChangePassword(login, password string) error
}

type ChangePasswordRepository interface {
	ChangePassword(login, password string) error
}

type ChangePasswordUC struct {
	r ChangePasswordRepository
	l logger.Logger
}

func NewChangePasswordUC(l logger.Logger, r ChangePasswordRepository) *ChangePasswordUC {
	return &ChangePasswordUC{l: l, r: r}
}