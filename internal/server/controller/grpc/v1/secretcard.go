package v1

import (
	pb "github.com/Kotletta-TT/GophKeeper/proto"
	"github.com/Kotletta-TT/GophKeeper/internal/server/entity"
	"github.com/Kotletta-TT/GophKeeper/internal/server/usecase"
	"context"
	"fmt"

	"github.com/google/uuid"
)

// SecretCardServiceServer представляет реализацию gRPC-сервера SecretCardService.
type SecretCardServiceServer struct {
	pb.UnimplementedSecretCardServiceServer
	UseCase usecase.SecretCardInterface
}

// NewSecretCardServiceServer создает новый экземпляр SecretCardServiceServer.
func NewSecretCardServiceServer(uc usecase.SecretCardInterface) *SecretCardServiceServer {
	return &SecretCardServiceServer{UseCase: uc}
}

// CreateSecretCard реализует метод CreateSecretCard интерфейса SecretCardService.
func (s *SecretCardServiceServer) CreateSecretCard(ctx context.Context, req *pb.CreateSecretCardRequest) (*pb.SecretCard, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user_id: %w", err)
	}

	card := &entity.SecretCard{
		ID:       uuid.New(),
		UserId:   userId,
		Name:     req.Name,
		URL:      req.Url,
		Login:    req.Login,
		Password: req.Password,
		Text:     req.Text,
	}

	if err := s.UseCase.CreateSecretCard(ctx, card); err != nil {
		return nil, err
	}

	return &pb.SecretCard{Id: card.ID.String(), UserId: card.UserId.String(), Name: card.Name, Url: card.URL, Login: card.Login, Password: card.Password, Text: card.Text}, nil
}

// ReadSecretCard реализует метод ReadSecretCard интерфейса SecretCardService.
func (s *SecretCardServiceServer) ReadSecretCard(ctx context.Context, req *pb.ReadSecretCardRequest) (*pb.SecretCard, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	card, err := s.UseCase.ReadSecretCard(ctx, id)
	if err != nil {
		return nil, err
	}

	return &pb.SecretCard{
		Id:       card.ID.String(),
		UserId:   card.UserId.String(),
		Name:     card.Name,
		Url:      card.URL,
		Login:    card.Login,
		Password: card.Password,
		Text:     card.Text,
	}, nil
}

// UpdateSecretCard реализует метод UpdateSecretCard интерфейса SecretCardService.
func (s *SecretCardServiceServer) UpdateSecretCard(ctx context.Context, req *pb.UpdateSecretCardRequest) (*pb.SecretCard, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user_id: %w", err)
	}

	card := &entity.SecretCard{
		ID:       id,
		UserId:   userID,
		Name:     req.Name,
		URL:      req.Url,
		Login:    req.Login,
		Password: req.Password,
		Text:     req.Text,
	}

	if err := s.UseCase.UpdateSecretCard(ctx, card); err != nil {
		return nil, err
	}

	return &pb.SecretCard{
		Id:       card.ID.String(),
		UserId:   card.UserId.String(),
		Name:     card.Name,
		Url:      card.URL,
		Login:    card.Login,
		Password: card.Password,
		Text:     card.Text,
	}, nil
}

// DeleteSecretCard реализует метод DeleteSecretCard интерфейса SecretCardService.
func (s *SecretCardServiceServer) DeleteSecretCard(ctx context.Context, req *pb.DeleteSecretCardRequest) (*pb.Empty, error) {
	id, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to parse id: %w", err)
	}

	if err := s.UseCase.DeleteSecretCard(ctx, id); err != nil {
		return nil, err
	}

	return &pb.Empty{}, nil
}

// ListSecretCard реализует метод ListSecretCard интерфейса SecretCardService.
func (s *SecretCardServiceServer) ListSecretCard(ctx context.Context, req *pb.ListSecretCardRequest) (*pb.ListSecretCardResponse, error) {
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user_id: %w", err)
	}

	cards, err := s.UseCase.GetSecretCardsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var responseCards []*pb.SecretCard
	for _, card := range cards {
		responseCards = append(responseCards, &pb.SecretCard{
			Id:       card.ID.String(),
			UserId:   card.UserId.String(),
			Name:     card.Name,
			Url:      card.URL,
			Login:    card.Login,
			Password: card.Password,
			Text:     card.Text,
		})
	}

	return &pb.ListSecretCardResponse{Cards: responseCards}, nil
}
