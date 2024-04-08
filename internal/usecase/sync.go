package usecase

import (
	"context"

	"GophKeeper-Server/internal/entity"
	filesecretcard "GophKeeper-Server/internal/repository/postgres/file_secretcard"
	metasecretcard "GophKeeper-Server/internal/repository/postgres/meta_secretcard"
	"GophKeeper-Server/internal/repository/postgres/secretcard"

	"github.com/google/uuid"
)

type SyncUsecase struct {
	CardRepository     secretcard.SecretCardRepository
	FileCardRepository filesecretcard.FileSecretCardRepository
	MetaCardRepository metasecretcard.MetaSecretCardRepository
}

func NewSyncUsecase(cardRepo secretcard.SecretCardRepository, fileRepo filesecretcard.FileSecretCardRepository, metaRepo metasecretcard.MetaSecretCardRepository) *SyncUsecase {
	return &SyncUsecase{
		CardRepository:     cardRepo,
		FileCardRepository: fileRepo,
		MetaCardRepository: metaRepo,
	}
}

func (uc *SyncUsecase) UpdateSyncCards(ctx context.Context, userID uuid.UUID, syncCards []*entity.SyncSecretCard) ([]*entity.SyncSecretCard, error) {
	updatedSyncCards := make([]*entity.SyncSecretCard, 0)
	for _, syncCard := range syncCards {
		existingCard, err := uc.CardRepository.ReadSecretCard(ctx, syncCard.Card.ID)
		if err != nil {
			return nil, err
		}

		if existingCard == nil || syncCard.Card.UpdateTime.After(existingCard.UpdateTime) {
			// Дата обновления карточки из среза моложе, чем дата обновления карточки в БД
			// Выполняем обновление карточки и всех связанных файлов и метаданных
			// Ваша логика обновления карточки здесь
			err = uc.CardRepository.UpdateSecretCard(ctx, syncCard.Card)
			if err != nil {
				return nil, err
			}

			// Обновление связанных файлов
			for _, file := range syncCard.Files {
				err := uc.FileCardRepository.Update(ctx, file)
				if err != nil {
					return nil, err
				}
			}

			// Обновление связанных метаданных
			for _, meta := range syncCard.Metas {
				err := uc.MetaCardRepository.Update(ctx, meta)
				if err != nil {
					return nil, err
				}
			}
		} else {

			dbFiles, err := uc.FileCardRepository.List(ctx, syncCard.Card.ID)
			if err != nil {
				return nil, err
			}
			dbMetas, err := uc.MetaCardRepository.List(ctx, syncCard.Card.ID)
			if err != nil {
				return nil, err
			}
			updatedSyncCards = append(updatedSyncCards, &entity.SyncSecretCard{
				Card:  existingCard,
				Files: dbFiles,
				Metas: dbMetas,
			})

		}
	}
	return updatedSyncCards, nil
}
