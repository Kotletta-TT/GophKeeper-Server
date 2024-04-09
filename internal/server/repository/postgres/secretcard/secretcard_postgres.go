package secretcard

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
	"github.com/Kotletta-TT/GophKeeper/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

type SecretCardRepository struct {
	*postgres.Postgres
}

func NewSecretCardRepository(pg *postgres.Postgres) *SecretCardRepository {
	return &SecretCardRepository{pg}
}

func (r *SecretCardRepository) CreateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	query, args, err := r.Builder.Insert("secret_cards").
		Columns("id", "user_id", "name", "url", "login", "password", "text", "update_time").
		Values(card.ID, card.UserId, card.Name, card.URL, card.Login, card.Password, card.Text, card.UpdateTime).ToSql()
	if err != nil {
		return fmt.Errorf("failed to generate SQL query: %w", err)
	}
	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to create secret card: %w", err)
	}
	return nil
}

// ReadSecretCard retrieves a secret card from the repository by ID.
func (r *SecretCardRepository) ReadSecretCard(ctx context.Context, id uuid.UUID) (*entity.SecretCard, error) {
	var card entity.SecretCard

	query, args, err := r.Builder.Select("id", "user_id", "name", "url", "login", "password", "text", "update_time").
		From("secret_cards").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to generate SQL query: %w", err)
	}

	row := r.Pool.QueryRow(ctx, query, args...)
	err = row.Scan(&card.ID, &card.UserId, &card.Name, &card.URL, &card.Login, &card.Password, &card.Text, &card.UpdateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("secret card not found")
		}
		return nil, fmt.Errorf("failed to read secret card: %w", err)
	}

	return &card, nil
}

// UpdateSecretCard обновляет существующую секретную карту в репозитории.
func (r *SecretCardRepository) UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	query, args, err := r.Builder.Update("secret_cards").
		Set("user_id", card.UserId).
		Set("name", card.Name).
		Set("url", card.URL).
		Set("login", card.Login).
		Set("password", card.Password).
		Set("text", card.Text).
		Set("update_time", card.UpdateTime).
		Where(squirrel.Eq{"id": card.ID}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to generate SQL query: %w", err)
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}
	return nil
}

// DeleteSecretCard удаляет секретную карту из репозитория по ID.
func (r *SecretCardRepository) DeleteSecretCard(ctx context.Context, id uuid.UUID) error {
	query, args, err := r.Builder.Delete("secret_cards").
		Where(squirrel.Eq{"id": id}).
		ToSql()

	if err != nil {
		return fmt.Errorf("failed to generate SQL query: %w", err)
	}

	_, err = r.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}
	return nil
}

func (r *SecretCardRepository) ListSecretCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.SecretCard, error) {
	var cards []entity.SecretCard

	query, args, err := r.Builder.Select("id", "user_id", "name", "url", "login", "password", "text", "update_time").
		From("secret_cards").
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("failed to generate SQL query: %w", err)
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var card entity.SecretCard
		err := rows.Scan(&card.ID, &card.UserId, &card.Name, &card.URL, &card.Login, &card.Password, &card.Text, &card.UpdateTime)
		if err != nil {
			return nil, fmt.Errorf("failed to scan secret card: %w", err)
		}
		cards = append(cards, card)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during iteration over secret cards rows: %w", err)
	}

	return cards, nil
}
