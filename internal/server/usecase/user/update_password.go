package user

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
)

type ChangePassword interface {
	ChangePassword(ctx context.Context, login, oldPassword, newPassword string) error
}

type ChangePasswordRepository interface {
	GetUser(ctx context.Context, login string) (*entity.User, error)
	ChangePassword(ctx context.Context, login, password string) error
}

type ChangePasswordUC struct {
	r ChangePasswordRepository
	v verifyFunc
	h HashFunc
}

func NewChangePasswordUC(r ChangePasswordRepository, h HashFunc, v verifyFunc) *ChangePasswordUC {
	return &ChangePasswordUC{h: h, v: v, r: r}
}

func (u *ChangePasswordUC) ChangePassword(ctx context.Context, login, oldPassword, newPassword string) error {
	user, err := u.r.GetUser(ctx, login)
	if err != nil {
		return err
	}
	err = u.v(user.Password, oldPassword)
	if err != nil {
		return err
	}
	hashedPassword, err := u.h(newPassword)
	if err != nil {
		return err
	}
	return u.r.ChangePassword(ctx, user.ID.String(), hashedPassword)
}
