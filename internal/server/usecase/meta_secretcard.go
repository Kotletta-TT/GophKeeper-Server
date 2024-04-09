package usecase

import (
	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
	metasecretcard "github.com/Kotletta-TT/GophKeeper/internal/server/repository/postgres/meta_secretcard"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type MetaSecretCardUsecase struct {
	Repository *metasecretcard.MetaSecretCardRepository
}

func NewMetaSecretCardUsecase(repo *metasecretcard.MetaSecretCardRepository) *MetaSecretCardUsecase {
	return &MetaSecretCardUsecase{
		Repository: repo,
	}
}

func (uc *MetaSecretCardUsecase) Create(ctx context.Context, meta *entity.MetaSecretCard) error {
	if err := uc.Repository.Create(ctx, meta); err != nil {
		return fmt.Errorf("failed to create meta secret card: %w", err)
	}
	return nil
}

func (uc *MetaSecretCardUsecase) Update(ctx context.Context, meta *entity.MetaSecretCard) error {
	if err := uc.Repository.Update(ctx, meta); err != nil {
		return fmt.Errorf("failed to update meta secret card: %w", err)
	}
	return nil
}

func (uc *MetaSecretCardUsecase) List(ctx context.Context, cardID uuid.UUID) ([]*entity.MetaSecretCard, error) {
	metas, err := uc.Repository.List(ctx, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to list meta secret cards: %w", err)
	}
	return metas, nil
}

func (uc *MetaSecretCardUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	if err := uc.Repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete meta secret card: %w", err)
	}
	return nil
}

func (uc *MetaSecretCardUsecase) Read(ctx context.Context, id uuid.UUID) (*entity.MetaSecretCard, error) {
	meta, err := uc.Repository.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to read meta secret card: %w", err)
	}
	return meta, nil
}
