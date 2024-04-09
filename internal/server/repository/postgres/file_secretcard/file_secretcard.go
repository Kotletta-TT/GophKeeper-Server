package filesecretcard

import (
	"context"
	"fmt"

	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
	"github.com/Kotletta-TT/GophKeeper/pkg/postgres"

	sq "github.com/Masterminds/squirrel"

	"github.com/google/uuid"
)

// FileSecretCardRepository представляет репозиторий для работы с файлами секретных карт.
type FileSecretCardRepository struct {
	pg *postgres.Postgres
}

// NewFileSecretCardRepository создает новый экземпляр репозитория FileSecretCardRepository.
func NewFileSecretCardRepository(pg *postgres.Postgres) *FileSecretCardRepository {
	return &FileSecretCardRepository{
		pg: pg,
	}
}

// Create создает новую запись файла секретной карты.
func (r *FileSecretCardRepository) Create(ctx context.Context, file *entity.FileSecretCard) error {
	query, args, err := r.pg.Builder.Insert("file_secret_cards").
		Columns("id", "card_id", "file").
		Values(file.ID, file.CardId, file.File).ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	if _, err := r.pg.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

// Update обновляет запись файла секретной карты.
func (r *FileSecretCardRepository) Update(ctx context.Context, file *entity.FileSecretCard) error {
	query, args, err := r.pg.Builder.Update("file_secret_cards").
		Set("file", file.File).
		Where(sq.Eq{"id": file.ID.String()}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	if _, err := r.pg.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

// Read возвращает запись файла секретной карты по её идентификатору.
func (r *FileSecretCardRepository) Read(ctx context.Context, id uuid.UUID) (*entity.FileSecretCard, error) {
	query, args, err := r.pg.Builder.Select("id", "card_id", "file").
		From("file_secret_cards").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL query: %w", err)
	}

	row := r.pg.Pool.QueryRow(ctx, query, args...)

	var file entity.FileSecretCard
	if err := row.Scan(&file.ID, &file.CardId, &file.File); err != nil {
		return nil, fmt.Errorf("failed to scan row: %w", err)
	}

	return &file, nil
}

// Delete удаляет запись файла секретной карты по её идентификатору.
func (r *FileSecretCardRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query, args, err := r.pg.Builder.Delete("file_secret_cards").
		Where(sq.Eq{"id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL query: %w", err)
	}

	if _, err := r.pg.Pool.Exec(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute SQL query: %w", err)
	}

	return nil
}

func (r *FileSecretCardRepository) List(ctx context.Context, cardID uuid.UUID) ([]*entity.FileSecretCard, error) {
	query, args, err := r.pg.Builder.Select("id", "card_id", "file").
		From("file_secret_cards").
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

	var files []*entity.FileSecretCard
	for rows.Next() {
		var file entity.FileSecretCard
		if err := rows.Scan(&file.ID, &file.CardId, &file.File); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		files = append(files, &file)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after scanning rows: %w", err)
	}

	return files, nil
}
