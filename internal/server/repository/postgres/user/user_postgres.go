package user

import (
	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
	"github.com/Kotletta-TT/GophKeeper/pkg/postgres"
	"context"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type UserRepositroy struct {
	*postgres.Postgres
}

func NewUserRepositroy(pg *postgres.Postgres) *UserRepositroy {
	return &UserRepositroy{pg}
}

func (u *UserRepositroy) GetUser(ctx context.Context, login string) (*entity.User, error) {
	user := entity.User{}

	query, args, err := u.Builder.Select("id", "login", "password").From("users").Where(sq.Eq{"login": login}).ToSql()
	if err != nil {
		fmt.Println(err)
		return &user, err
	}

	row := u.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, pgx.ErrNoRows) {
			return &user, errors.New("user not found")
		}
		return &user, err
	}

	return &user, nil
}

func (u *UserRepositroy) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user := entity.User{}

	query, args, err := u.Builder.Select("id", "login", "password").From("users").Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return &user, err
	}
	row := u.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &user, errors.New("user not found")
		}
		return &user, err
	}
	return &user, nil
}

func (u *UserRepositroy) CreateUser(ctx context.Context, login, password string) error {
	var userExists bool
	err := u.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)", login).Scan(&userExists)
	if err != nil {
		return err
	}
	if userExists {
		return fmt.Errorf("user already exists")
	}

	query := u.Builder.Insert("users").
		Columns("login", "password").
		Values(login, password).Suffix("RETURNING \"id\"")
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	row := u.Pool.QueryRow(ctx, sql, args...)
	var userID string
	err = row.Scan(&userID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepositroy) ChangePassword(ctx context.Context, id, password string) error {
	sql, args, err := u.Builder.Update("users").
		Set("password", password).
		Where(sq.Eq{"id": id}).ToSql()
	if err != nil {
		return err
	}
	_, err = u.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return err
	}
	return nil
}
