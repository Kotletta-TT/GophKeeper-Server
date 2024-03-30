package secretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/pkg/postgres"
	"context"

	"github.com/google/uuid"
)

type SecretCardRepositroy struct {
	*postgres.Postgres
}

func NewSecretCardRepositroy(pg *postgres.Postgres) *SecretCardRepositroy {
	return &SecretCardRepositroy{pg}
}

func (s *SecretCardRepositroy) AddSecretCard(ctx context.Context, card *entity.SecretCard) error {
	sql, _, err := s.Builder.Insert("secret_cards").Columns("user_id", "name", "url", "login", "password", "text", "meta", "file", "update_time").
		Values(card.UserId, card.Name, card.URL, card.Login, card.Password, card.Text, card.Meta, card.Files, card.UpdateTime).
		ToSql()
	if err != nil {
		return err
	}
	_, err = s.Pool.Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *SecretCardRepositroy) UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	sql, _, err := s.Builder.Update("secret_cards").SetMap(map[string]interface{}{
		"name":        card.Name,
		"url":         card.URL,
		"login":       card.Login,
		"password":    card.Password,
		"text":        card.Text,
		"meta":        card.Meta,
		"file":        card.Files,
		"update_time": card.UpdateTime,
	}).Where("id = ?", card.ID).ToSql()
	if err != nil {
		return err
	}
	_, err = s.Pool.Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *SecretCardRepositroy) GetSecretCard(ctx context.Context, id uuid.UUID) (*entity.SecretCard, error) {
	sql, _, err := s.Builder.Select(
		"id",
		"user_id",
		"name",
		"url",
		"login",
		"password",
		"text",
		"meta",
		"file",
		"update_time").From("secret_cards").Where("id = ?", id).ToSql()
	if err != nil {
		return nil, err
	}
	row := s.Pool.QueryRow(ctx, sql)
	secretCard := &entity.SecretCard{}
	err = row.Scan(
		&secretCard.ID,
		&secretCard.UserId,
		&secretCard.Name,
		&secretCard.URL,
		&secretCard.Login,
		&secretCard.Password,
		&secretCard.Text,
		&secretCard.Meta,
		&secretCard.Files,
		&secretCard.UpdateTime,
	)
	if err != nil {
		return nil, err
	}
	return secretCard, nil
}

func (s *SecretCardRepositroy) DeleteSecretCard(ctx context.Context, id uuid.UUID) error {
	sql, _, err := s.Builder.Delete("secret_cards").Where("id = ?", id).ToSql()
	if err != nil {
		return err
	}
	_, err = s.Pool.Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}

func (s *SecretCardRepositroy) GetSecretCards(ctx context.Context, userId uuid.UUID) ([]*entity.SecretCard, error) {
	sql, _, err := s.Builder.Select(
		"id",
		"user_id",
		"name",
		"url",
		"login",
		"password",
		"text",
		"meta",
		"file",
		"update_time").From("secret_cards").Where("user_id = ?", userId).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	var secretCards []*entity.SecretCard
	for rows.Next() {
		secretCard := &entity.SecretCard{}
		err = rows.Scan(
			&secretCard.ID,
			&secretCard.UserId,
			&secretCard.Name,
			&secretCard.URL,
			&secretCard.Login,
			&secretCard.Password,
			&secretCard.Text,
			&secretCard.Meta,
			&secretCard.Files,
			&secretCard.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		secretCards = append(secretCards, secretCard)
	}
	return secretCards, nil
}

func (s *SecretCardRepositroy) GetSecretCardsByName(ctx context.Context, name string) ([]*entity.SecretCard, error) {
	sql, _, err := s.Builder.Select(
		"id",
		"user_id",
		"name",
		"url",
		"login",
		"password",
		"text",
		"meta",
		"file",
		"update_time").From("secret_cards").Where("name = ?", name).ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := s.Pool.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	var secretcards []*entity.SecretCard
	for rows.Next() {
		secretCard := &entity.SecretCard{}
		err = rows.Scan(
			&secretCard.ID,
			&secretCard.UserId,
			&secretCard.Name,
			&secretCard.URL,
			&secretCard.Login,
			&secretCard.Password,
			&secretCard.Text,
			&secretCard.Meta,
			&secretCard.Files,
			&secretCard.UpdateTime,
		)
		if err != nil {
			return nil, err
		}
		secretcards = append(secretcards, secretCard)
	}
	return secretcards, nil
}
