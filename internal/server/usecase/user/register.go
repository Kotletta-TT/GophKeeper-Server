package user

import (
	"context"
)

type Register interface {
	Register(ctx context.Context, login, password string) error
}

type RegisterRepository interface {
	CreateUser(ctx context.Context, login, password string) error
}

type HashFunc func(password string) (string, error)

type RegisterUC struct {
	r RegisterRepository
	h HashFunc
}

func NewRegisterUC(r RegisterRepository, h HashFunc) *RegisterUC {
	return &RegisterUC{r: r, h: h}
}

func (uc *RegisterUC) Register(ctx context.Context, login, password string) error {
	hashedPassword, err := uc.h(password)
	if err != nil {
		return err
	}
	err = uc.r.CreateUser(ctx, login, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}
