package v1

import (
	"GophKeeper-Server/internal/controller/grpc/v1/proto"
	"GophKeeper-Server/internal/entity"
	"context"
	"fmt"

	uc "GophKeeper-Server/internal/usecase"

	"github.com/google/uuid"
)

type FileSecretCardServer struct {
	proto.UnimplementedFileSecretCardServiceServer
	fileSecretCardUseCase uc.FileSecretCardInterface
}

// NewFileSecretCardServer создает новый экземпляр сервера FileSecretCardServer.
func NewFileSecretCardServer(useCase uc.FileSecretCardInterface) *FileSecretCardServer {
	return &FileSecretCardServer{
		fileSecretCardUseCase: useCase,
	}
}

// CreateFileSecretCard реализует метод CreateFileSecretCard сервиса FileSecretCardService.
func (s *FileSecretCardServer) CreateFileSecretCard(ctx context.Context, req *proto.CreateFileSecretCardRequest) (*proto.FileSecretCard, error) {
	cardID, err := uuid.Parse(req.GetCardId())
	if err != nil {
		return nil, fmt.Errorf("invalid card ID: %w", err)
	}

	file := &entity.FileSecretCard{
		CardId: cardID,
		File:   req.GetFile(),
	}

	if err := s.fileSecretCardUseCase.Create(ctx, file); err != nil {
		return nil, fmt.Errorf("failed to create file secret card: %w", err)
	}

	return &proto.FileSecretCard{
		Id:     file.ID.String(),
		CardId: file.CardId.String(),
		File:   file.File,
	}, nil
}

// ReadFileSecretCard реализует метод ReadFileSecretCard сервиса FileSecretCardService.
func (s *FileSecretCardServer) ReadFileSecretCard(ctx context.Context, req *proto.ReadFileSecretCardRequest) (*proto.FileSecretCard, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	file, err := s.fileSecretCardUseCase.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to read file secret card: %w", err)
	}

	return &proto.FileSecretCard{
		Id:     file.ID.String(),
		CardId: file.CardId.String(),
		File:   file.File,
	}, nil
}

// UpdateFileSecretCard реализует метод UpdateFileSecretCard сервиса FileSecretCardService.
func (s *FileSecretCardServer) UpdateFileSecretCard(ctx context.Context, req *proto.UpdateFileSecretCardRequest) (*proto.FileSecretCard, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	cardID, err := uuid.Parse(req.GetCardId())
	if err != nil {
		return nil, fmt.Errorf("invalid card ID: %w", err)
	}

	file := &entity.FileSecretCard{
		ID:     id,
		CardId: cardID,
		File:   req.GetFile(),
	}

	if err := s.fileSecretCardUseCase.Update(ctx, file); err != nil {
		return nil, fmt.Errorf("failed to update file secret card: %w", err)
	}

	return &proto.FileSecretCard{
		Id:     file.ID.String(),
		CardId: file.CardId.String(),
		File:   file.File,
	}, nil
}

// DeleteFileSecretCard реализует метод DeleteFileSecretCard сервиса FileSecretCardService.
func (s *FileSecretCardServer) DeleteFileSecretCard(ctx context.Context, req *proto.DeleteFileSecretCardRequest) (*proto.Empty, error) {
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %w", err)
	}

	if err := s.fileSecretCardUseCase.Delete(ctx, id); err != nil {
		return nil, fmt.Errorf("failed to delete file secret card: %w", err)
	}

	return &proto.Empty{}, nil
}

// ListFileSecretCard реализует метод ListFileSecretCard сервиса FileSecretCardService.
func (s *FileSecretCardServer) ListFileSecretCard(ctx context.Context, req *proto.ListFileSecretCardRequest) (*proto.ListFileSecretCardResponse, error) {
	cardID, err := uuid.Parse(req.GetCardId())
	if err != nil {
		return nil, fmt.Errorf("invalid card ID: %w", err)
	}

	files, err := s.fileSecretCardUseCase.List(ctx, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to list file secret cards: %w", err)
	}

	var protoFiles []*proto.FileSecretCard
	for _, file := range files {
		protoFiles = append(protoFiles, &proto.FileSecretCard{
			Id:     file.ID.String(),
			CardId: file.CardId.String(),
			File:   file.File,
		})
	}

	return &proto.ListFileSecretCardResponse{
		Files: protoFiles,
	}, nil
}
