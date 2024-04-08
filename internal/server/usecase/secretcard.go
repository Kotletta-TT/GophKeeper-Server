package usecase

import (
	"GophKeeper-Server/internal/entity"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type SecretCardInterface interface {
	GetSecretCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.SecretCard, error)
	CreateSecretCard(ctx context.Context, card *entity.SecretCard) error
	UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error
	DeleteSecretCard(ctx context.Context, id uuid.UUID) error
	ReadSecretCard(ctx context.Context, id uuid.UUID) (*entity.SecretCard, error)
}

type SecretCardRepository interface {
	ListSecretCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.SecretCard, error)
	CreateSecretCard(ctx context.Context, card *entity.SecretCard) error
	UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error
	DeleteSecretCard(ctx context.Context, id uuid.UUID) error
	ReadSecretCard(ctx context.Context, id uuid.UUID) (*entity.SecretCard, error)
}

type SecretCardUseCase struct {
	secretCardRepo SecretCardRepository
}

// NewSecretCardUseCase создает новый экземпляр SecretCardUseCase.
func NewSecretCardUseCase(repo SecretCardRepository) *SecretCardUseCase {
	return &SecretCardUseCase{secretCardRepo: repo}
}

// GetSecretCardsByUserID получает все секретные карты для указанного пользователя из репозитория.
func (uc *SecretCardUseCase) GetSecretCardsByUserID(ctx context.Context, userID uuid.UUID) ([]entity.SecretCard, error) {
	cards, err := uc.secretCardRepo.ListSecretCardsByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret cards for user: %w", err)
	}
	return cards, nil
}

// CreateSecretCard создает новую секретную карту и сохраняет ее в репозитории.
func (uc *SecretCardUseCase) CreateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	err := uc.secretCardRepo.CreateSecretCard(ctx, card)
	if err != nil {
		return fmt.Errorf("failed to create secret card: %w", err)
	}
	return nil
}

// UpdateSecretCard обновляет существующую секретную карту в репозитории.
func (uc *SecretCardUseCase) UpdateSecretCard(ctx context.Context, card *entity.SecretCard) error {
	err := uc.secretCardRepo.UpdateSecretCard(ctx, card)
	if err != nil {
		return fmt.Errorf("failed to update secret card: %w", err)
	}
	return nil
}

// DeleteSecretCard удаляет секретную карту из репозитории по ID.
func (uc *SecretCardUseCase) DeleteSecretCard(ctx context.Context, id uuid.UUID) error {
	err := uc.secretCardRepo.DeleteSecretCard(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete secret card: %w", err)
	}
	return nil
}

func (uc *SecretCardUseCase) ReadSecretCard(ctx context.Context, id uuid.UUID) (*entity.SecretCard, error) {
	card, err := uc.secretCardRepo.ReadSecretCard(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get secret card: %w", err)
	}
	return card, nil
}
