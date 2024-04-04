package user

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/pkg/postgres"
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
)

type UserRepositroy struct {
	*postgres.Postgres
}

func NewUserRepositroy(pg *postgres.Postgres) *UserRepositroy {
	return &UserRepositroy{pg}
}

func (u *UserRepositroy) GetUser(ctx context.Context, login string) (*entity.User, error) {
	sql, _, err := u.Builder.Select("id, login, password").From("users").Where(squirrel.Eq{"login": login}).ToSql()
	if err != nil {
		return nil, err
	}
	row := u.Pool.QueryRow(ctx, sql, login)
	usr := &entity.User{}
	err = row.Scan(&usr.ID, &usr.Login, &usr.Password)
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *UserRepositroy) CreateUser(ctx context.Context, login, password string) error {
	sql, _, err := u.Builder.Insert("users").Columns("login", "password").Values(login, password).ToSql()
	if err != nil {
		return err
	}
	_, err = u.Pool.Exec(ctx, sql, login, password)
	if err != nil {
		return err
	}
	return nil
}
