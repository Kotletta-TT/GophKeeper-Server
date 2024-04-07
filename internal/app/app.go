package app

import (
	"GophKeeper-Server/config"
	v1 "GophKeeper-Server/internal/controller/grpc/v1"
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	filesecretcard "GophKeeper-Server/internal/repository/postgres/file_secretcard"
	"GophKeeper-Server/internal/repository/postgres/secretcard"
	repouser "GophKeeper-Server/internal/repository/postgres/user"
	"GophKeeper-Server/internal/usecase"
	"GophKeeper-Server/internal/usecase/user"
	"GophKeeper-Server/internal/utils"
	"GophKeeper-Server/logger"
	"GophKeeper-Server/pkg/postgres"
	"context"
	"net"

	"google.golang.org/grpc"
)

func Run(ctx context.Context, cfg *config.Config, l logger.Logger) error {
	var Err error
	pg, err := postgres.New(ctx, l, "postgresql://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable")
	if err != nil {
		return err
	}
	_, err = pg.Pool.Exec(ctx, TableUsers)
	if err != nil {
		return err
	}
	_, err = pg.Pool.Exec(ctx, TableSecretCards)
	if err != nil {
		return err
	}
	_, err = pg.Pool.Exec(ctx, TableFileSecretCards)
	if err != nil {
		return err
	}
	_, err = pg.Pool.Exec(ctx, TableMetaSecretCards)
	if err != nil {
		return err
	}
	ur := repouser.NewUserRepositroy(pg)
	ruc := user.NewRegisterUC(ur, utils.HashPassword)
	guc := user.NewGetUserUC(ur, utils.VerifyPassword)
	cuc := user.NewChangePasswordUC(ur, utils.HashPassword, utils.VerifyPassword)

	repo := secretcard.NewSecretCardRepository(pg)
	frepo := filesecretcard.NewFileSecretCardRepository(pg)

	// Инициализация use case
	useCase := usecase.NewSecretCardUseCase(repo)
	fuseCase := usecase.NewFileSecretCardUseCase(frepo)

	// Инициализация gRPC-сервера
	srv := v1.NewSecretCardServiceServer(useCase)
	fsrv := v1.NewFileSecretCardServer(fuseCase)

	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			Err = err
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, &v1.UserServer{Ruc: ruc, Guc: guc, Cuc: cuc})
		pb.RegisterSecretCardServiceServer(s, srv)
		pb.RegisterFileSecretCardServiceServer(s, fsrv)

		if err := s.Serve(listen); err != nil {
			Err = err
		}
	}()
	<-ctx.Done()
	if Err != nil {
		return Err
	} else {
		return ctx.Err()
	}
}
