package user

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
)

type GeterUser interface {
	GetUser(ctx context.Context, login, password string) (*entity.User, error)
}

type GetUserRepository interface {
	GetUser(ctx context.Context, login string) (*entity.User, error)
}

type verifyFunc func(hashedPassword, password string) error

type GetUser struct {
	r GetUserRepository
	v verifyFunc
}

func NewGetUserUC(r GetUserRepository, f verifyFunc) *GetUser {
	return &GetUser{r: r, v: f}
}

func (g *GetUser) GetUser(ctx context.Context, login, password string) (*entity.User, error) {
	u, err := g.r.GetUser(ctx, login)
	if err != nil {
		return nil, err
	}
	err = g.v(u.Password, password)
	if err != nil {
		return nil, err
	}
	return u, nil
}
