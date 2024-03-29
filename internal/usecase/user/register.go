package user

import (
	"GophKeeper-Server/internal/entity"
	customErr "GophKeeper-Server/internal/errors"
	"GophKeeper-Server/logger"
	"fmt"
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
		uc.l.Errorf(
			fmt.Sprintf(
				"repository GetUser err: %s login: %s",
				err.Error(),
				login))
		return customErr.ErrDatabaseInternal(err)
	}
	if usr != nil {
		err := fmt.Errorf("user %s already exists", login)
		uc.l.Infof(err.Error())
		return customErr.ErrAlreadyExists(err)
	}
	hashPass, err := uc.h(password)
	if err != nil {
		uc.l.Errorf(
			fmt.Sprintf(
				"hash password from login: %s err: %s",
				login,
				err.Error(),
			))
		return customErr.ErrDatabaseInternal(err)
	}
	if err := uc.r.Register(login, hashPass); err != nil {
		uc.l.Errorf(
			fmt.Sprintf(
				"repository Register err: %s login: %s",
				err.Error(),
				login,
			))
		return customErr.ErrDatabaseInternal(err)
	}
	uc.l.Infof("user %s created!", login)
	return nil
}
