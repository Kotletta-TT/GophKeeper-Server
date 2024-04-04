package user

import (
	"GophKeeper-Server/internal/entity"
	customErr "GophKeeper-Server/internal/errors"
	"GophKeeper-Server/logger"
	"context"
	"fmt"
)

type Register interface {
	Register(ctx context.Context, login, password string) error
}

type RegisterRepository interface {
	GetUser(ctx context.Context, login string) (*entity.User, error)
	CreateUser(ctx context.Context, login, password string) error
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

func (uc *RegisterUC) Register(ctx context.Context, login, password string) error {
	usr, err := uc.r.GetUser(ctx, login)
	if err != nil {
		uc.l.Error("get user from repository", "error", err, "login", login)
		return customErr.ErrDatabaseInternal(err)
	}
	if usr != nil {
		err := fmt.Errorf("user %s already exists", login)
		uc.l.Info("register user", "error", err, "login", login)
		return customErr.ErrAlreadyExists(err)
	}
	hashPass, err := uc.h(password)
	if err != nil {
		uc.l.Error("hash password", "error", err, "login", login)
		return customErr.ErrDatabaseInternal(err)
	}
	if err := uc.r.CreateUser(ctx, login, hashPass); err != nil {
		uc.l.Error("create user", "error", err, "login", login)
		return customErr.ErrDatabaseInternal(err)
	}
	uc.l.Info("user created", "login", login)
	return nil
}
