package v1

import (
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	"GophKeeper-Server/internal/entity"
	"context"
	"fmt"

	uc "GophKeeper-Server/internal/usecase"

	"github.com/google/uuid"
)

type MetaSecretCardServer struct {
	pb.UnimplementedMetaSecretCardServiceServer
	UseCase *uc.MetaSecretCardUsecase
}

func NewMetaSecretCardServer(useCase *uc.MetaSecretCardUsecase) *MetaSecretCardServer {
	return &MetaSecretCardServer{
		UseCase: useCase,
	}
}

func (s *MetaSecretCardServer) CreateMetaSecretCard(ctx context.Context, req *pb.MetaSecretCard) (*pb.MetaSecretCard, error) {
	meta := &entity.MetaSecretCard{
		ID:     uuid.New(),
		CardId: uuid.MustParse(req.CardId),
		Key:    req.Key,
		Value:  req.Value,
	}

	if err := s.UseCase.Create(ctx, meta); err != nil {
		return nil, fmt.Errorf("failed to create meta secret card: %w", err)
	}

	return &pb.MetaSecretCard{
		Id:     meta.ID.String(),
		CardId: meta.CardId.String(),
		Key:    meta.Key,
		Value:  meta.Value,
	}, nil
}

func (s *MetaSecretCardServer) UpdateMetaSecretCard(ctx context.Context, req *pb.MetaSecretCard) (*pb.MetaSecretCard, error) {
	meta := &entity.MetaSecretCard{
		ID:     uuid.MustParse(req.Id),
		CardId: uuid.MustParse(req.CardId),
		Key:    req.Key,
		Value:  req.Value,
	}

	if err := s.UseCase.Update(ctx, meta); err != nil {
		return nil, fmt.Errorf("failed to update meta secret card: %w", err)
	}

	return &pb.MetaSecretCard{
		Id:     meta.ID.String(),
		CardId: meta.CardId.String(),
		Key:    meta.Key,
		Value:  meta.Value,
	}, nil
}

func (s *MetaSecretCardServer) ListMetaSecretCard(ctx context.Context, req *pb.ListMetaSecretCardRequest) (*pb.ListMetaSecretCardResponse, error) {
	cardID := uuid.MustParse(req.CardId)

	metas, err := s.UseCase.List(ctx, cardID)
	if err != nil {
		return nil, fmt.Errorf("failed to list meta secret cards: %w", err)
	}

	var resp pb.ListMetaSecretCardResponse
	for _, meta := range metas {
		resp.Cards = append(resp.Cards, &pb.MetaSecretCard{
			Id:     meta.ID.String(),
			CardId: meta.CardId.String(),
			Key:    meta.Key,
			Value:  meta.Value,
		})
	}

	return &resp, nil
}

func (s *MetaSecretCardServer) DeleteMetaSecretCard(ctx context.Context, req *pb.DeleteMetaSecretCardRequest) (*pb.Empty, error) {
	id := uuid.MustParse(req.Id)
	if err := s.UseCase.Delete(ctx, id); err != nil {
		return nil, fmt.Errorf("failed to delete meta secret card: %w", err)
	}
	return &pb.Empty{}, nil
}

func (s *MetaSecretCardServer) ReadMetaSecretCard(ctx context.Context, req *pb.ReadMetaSecretCardRequest) (*pb.MetaSecretCard, error) {
	id := uuid.MustParse(req.Id)
	meta, err := s.UseCase.Read(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to read meta secret card: %w", err)
	}
	if meta == nil {
		return nil, fmt.Errorf("meta secret card not found")
	}
	return &pb.MetaSecretCard{
		Id:     meta.ID.String(),
		CardId: meta.CardId.String(),
		Key:    meta.Key,
		Value:  meta.Value,
	}, nil
}
