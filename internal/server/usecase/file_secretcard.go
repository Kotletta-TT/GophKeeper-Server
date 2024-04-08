package usecase

import (
	"GophKeeper-Server/internal/entity"
	repository "GophKeeper-Server/internal/repository/postgres/file_secretcard"
	"context"

	"github.com/google/uuid"
)

type FileSecretCardInterface interface {
	Create(ctx context.Context, file *entity.FileSecretCard) error
	Update(ctx context.Context, file *entity.FileSecretCard) error
	Read(ctx context.Context, id uuid.UUID) (*entity.FileSecretCard, error)
	Delete(ctx context.Context, id uuid.UUID) error
	List(ctx context.Context, cardID uuid.UUID) ([]*entity.FileSecretCard, error)
}

// FileSecretCardUseCase представляет юз-кейс для работы с файлами секретных карт.
type FileSecretCardUseCase struct {
	repo *repository.FileSecretCardRepository
}

// NewFileSecretCardUseCase создает новый экземпляр юз-кейса FileSecretCardUseCase.
func NewFileSecretCardUseCase(repo *repository.FileSecretCardRepository) *FileSecretCardUseCase {
	return &FileSecretCardUseCase{
		repo: repo,
	}
}

// Create создает новую запись файла секретной карты.
func (uc *FileSecretCardUseCase) Create(ctx context.Context, file *entity.FileSecretCard) error {
	return uc.repo.Create(ctx, file)
}

// Update обновляет запись файла секретной карты.
func (uc *FileSecretCardUseCase) Update(ctx context.Context, file *entity.FileSecretCard) error {
	return uc.repo.Update(ctx, file)
}

// Read возвращает запись файла секретной карты по её идентификатору.
func (uc *FileSecretCardUseCase) Read(ctx context.Context, id uuid.UUID) (*entity.FileSecretCard, error) {
	return uc.repo.Read(ctx, id)
}

// Delete удаляет запись файла секретной карты по её идентификатору.
func (uc *FileSecretCardUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *FileSecretCardUseCase) List(ctx context.Context, cardID uuid.UUID) ([]*entity.FileSecretCard, error) {
	return uc.repo.List(ctx, cardID)
}

