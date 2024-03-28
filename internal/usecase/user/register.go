package user

import (
	"GophKeeper-Server/internal/entity"
	customErr "GophKeeper-Server/internal/errors"
	"GophKeeper-Server/logger"
)

type Register interface {
	Register(login, password string) error
}

type RegisterRepository interface {
	GetUser(login string) (*entity.User, error)
	Register(login, password string) error
}

type HashFunc func(password string) (string, error)

type RegisterUC struct {
	l logger.Logger
	r RegisterRepository
	h HashFunc
}

func NewRegisterUC(l logger.Logger, r RegisterRepository, h HashFunc) *RegisterUC {
	return &RegisterUC{l: l, r: r, h: h}
}

func (uc *RegisterUC) Register(login, password string) error {
	usr, err := uc.r.GetUser(login)
	if err != nil {
		return customErr.ErrDatabaseInternal(err)
	}
	if usr != nil {
		return customErr.ErrAlreadyExists(err)
	}
	hashPass, err := uc.h(password)
	if err != nil {
		return customErr.ErrDatabaseInternal(err)
	}
	return uc.r.Register(login, hashPass)
}
