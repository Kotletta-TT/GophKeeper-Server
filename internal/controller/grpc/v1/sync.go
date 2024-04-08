package v1

import (
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	"GophKeeper-Server/internal/entity"
	"GophKeeper-Server/internal/usecase"
	"context"

	"github.com/google/uuid"
)

type SyncSecretCardServer struct {
	pb.UnimplementedSyncServiceServer
	SyncUseCase *usecase.SyncUsecase
}

func NewSyncSecretCardServer(syncUseCase *usecase.SyncUsecase) *SyncSecretCardServer {
	return &SyncSecretCardServer{
		SyncUseCase: syncUseCase,
	}
}

func (s *SyncSecretCardServer) Sync(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
	userID := uuid.MustParse(req.UserId)
	cards := s.ProtoToSync(req.SecretCards)
	result, err := s.SyncUseCase.UpdateSyncCards(ctx, userID, cards)
	if err != nil {
		return nil, err
	}

	return &pb.SyncResponse{SecretCards: s.SyncToProto(result)}, nil
}

func (s *SyncSecretCardServer) ProtoToSync(protoSync []*pb.SyncSecretCard) []*entity.SyncSecretCard {
	syncs := make([]*entity.SyncSecretCard, 0)
	for _, sync := range protoSync {
		syncs = append(syncs, &entity.SyncSecretCard{
			Card: &entity.SecretCard{
				ID:       uuid.MustParse(sync.SecretCard.Id),
				UserId:   uuid.MustParse(sync.SecretCard.UserId),
				Name:     sync.SecretCard.Name,
				URL:      sync.SecretCard.Url,
				Login:    sync.SecretCard.Login,
				Password: sync.SecretCard.Password,
				Text:     sync.SecretCard.Text,
			},
			Files: s.FileProtoToSync(sync.FileSecretCard),
			Metas: s.MetaProtoToSync(sync.MetaSecretCard),
		})
	}
	return syncs
}

func (s *SyncSecretCardServer) FileProtoToSync(protoFiles []*pb.FileSecretCard) []*entity.FileSecretCard {
	files := make([]*entity.FileSecretCard, 0)
	for _, file := range protoFiles {
		files = append(files, &entity.FileSecretCard{
			ID:     uuid.MustParse(file.Id),
			CardId: uuid.MustParse(file.CardId),
			File:   file.File,
		})
	}
	return files
}

func (s *SyncSecretCardServer) MetaProtoToSync(protoMetas []*pb.MetaSecretCard) []*entity.MetaSecretCard {
	metas := make([]*entity.MetaSecretCard, 0)
	for _, meta := range protoMetas {
		metas = append(metas, &entity.MetaSecretCard{
			ID:     uuid.MustParse(meta.Id),
			CardId: uuid.MustParse(meta.CardId),
			Key:    meta.Key,
			Value:  meta.Value,
		})
	}
	return metas
}

func (s *SyncSecretCardServer) SyncToProto(syncs []*entity.SyncSecretCard) []*pb.SyncSecretCard {
	protoSyncs := make([]*pb.SyncSecretCard, 0)
	for _, sync := range syncs {
		protoSyncs = append(protoSyncs, &pb.SyncSecretCard{
			SecretCard: &pb.SecretCard{
				Id:       sync.Card.ID.String(),
				UserId:   sync.Card.UserId.String(),
				Name:     sync.Card.Name,
				Url:      sync.Card.URL,
				Login:    sync.Card.Login,
				Password: sync.Card.Password,
				Text:     sync.Card.Text,
			},
			FileSecretCard: s.FileSyncToProto(sync.Files),
			MetaSecretCard: s.MetaSyncToProto(sync.Metas),
		})
	}
	return protoSyncs
}

func (s *SyncSecretCardServer) FileSyncToProto(files []*entity.FileSecretCard) []*pb.FileSecretCard {
	protoFiles := make([]*pb.FileSecretCard, 0)
	for _, file := range files {
		protoFiles = append(protoFiles, &pb.FileSecretCard{
			Id:     file.ID.String(),
			CardId: file.CardId.String(),
			File:   file.File,
		})
	}
	return protoFiles
}

func (s *SyncSecretCardServer) MetaSyncToProto(metas []*entity.MetaSecretCard) []*pb.MetaSecretCard {
	protoMetas := make([]*pb.MetaSecretCard, 0)
	for _, meta := range metas {
		protoMetas = append(protoMetas, &pb.MetaSecretCard{
			Id:     meta.ID.String(),
			CardId: meta.CardId.String(),
			Key:    meta.Key,
			Value:  meta.Value,
		})
	}
	return protoMetas
}
