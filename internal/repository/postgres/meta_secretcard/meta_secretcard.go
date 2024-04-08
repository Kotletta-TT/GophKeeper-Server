package metasecretcard

import (
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/pkg/postgres"
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type MetaSecretCardRepository struct {
	pg *postgres.Postgres
}

func NewMetaSecretCardRepository(pg *postgres.Postgres) *MetaSecretCardRepository {
	return &MetaSecretCardRepository{
		pg: pg,
	}
}

func (r *MetaSecretCardRepository) Create(ctx context.Context, meta *entity.MetaSecretCard) error {
	query, args, err := r.pg.Builder.
		Insert("meta_secret_card").
		Columns("id", "card_id", "key", "value").
		Values(meta.ID, meta.CardId, meta.Key, meta.Value).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}
	
	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

func (r *MetaSecretCardRepository) Update(ctx context.Context, meta *entity.MetaSecretCard) error {
	query, args, err := r.pg.Builder.
		Update("meta_secret_card").
		Set("card_id", meta.CardId).
		Set("key", meta.Key).
		Set("value", meta.Value).
		Where(sq.Eq{"id": meta.ID}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

func (r *MetaSecretCardRepository) List(ctx context.Context, cardID uuid.UUID) ([]*entity.MetaSecretCard, error) {
	query, args, err := r.pg.Builder.
		Select("id", "card_id", "key", "value").
		From("meta_secret_card").
		Where(sq.Eq{"card_id": cardID}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	rows, err := r.pg.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute SQL query: %w", err)
	}
	defer rows.Close()

	var metas []*entity.MetaSecretCard
	for rows.Next() {
		var meta entity.MetaSecretCard
		err := rows.Scan(&meta.ID, &meta.CardId, &meta.Key, &meta.Value)
		if err != nil {
			return nil, fmt.Errorf("failed to scan meta secret card row: %w", err)
		}
		metas = append(metas, &meta)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to list meta secret cards: %w", err)
	}

	return metas, nil
}

func (r *MetaSecretCardRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query, args, err := r.pg.Builder.
		Delete("meta_secret_card").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = r.pg.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

func (r *MetaSecretCardRepository) Read(ctx context.Context, id uuid.UUID) (*entity.MetaSecretCard, error) {
	query, args, err := r.pg.Builder.
		Select("id", "card_id", "key", "value").
		From("meta_secret_card").
		Where(sq.Eq{"id": id}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	row := r.pg.Pool.QueryRow(ctx, query, args...)
	var meta entity.MetaSecretCard
	err = row.Scan(&meta.ID, &meta.CardId, &meta.Key, &meta.Value)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // Not found
		}
		return nil, fmt.Errorf("failed to read meta secret card: %w", err)
	}

	return &meta, nil
}