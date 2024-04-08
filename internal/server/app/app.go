package app

import (
	"GophKeeper-Server/config"
	v1 "GophKeeper-Server/internal/controller/grpc/v1"
	"GophKeeper-Server/internal/controller/grpc/v1/interceptors"
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	filesecretcard "GophKeeper-Server/internal/repository/postgres/file_secretcard"
	metasecretcard "GophKeeper-Server/internal/repository/postgres/meta_secretcard"
	"GophKeeper-Server/internal/repository/postgres/secretcard"
	repouser "GophKeeper-Server/internal/repository/postgres/user"
	"GophKeeper-Server/internal/service"
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
	aServ := service.AuthService{}
	usrSrv := v1.NewUserServiceServer(ruc, guc, cuc, aServ)

	repo := secretcard.NewSecretCardRepository(pg)
	frepo := filesecretcard.NewFileSecretCardRepository(pg)
	mrepo := metasecretcard.NewMetaSecretCardRepository(pg)

	useCase := usecase.NewSecretCardUseCase(repo)
	fuseCase := usecase.NewFileSecretCardUseCase(frepo)
	museCase := usecase.NewMetaSecretCardUsecase(mrepo)
	suseCase := usecase.NewSyncUsecase(repo, frepo, mrepo)

	srv := v1.NewSecretCardServiceServer(useCase)
	fsrv := v1.NewFileSecretCardServer(fuseCase)
	msrv := v1.NewMetaSecretCardServer(museCase)
	ssrv := v1.NewSyncSecretCardServer(suseCase)
	interceptor := interceptors.AuthInterceptor{
		AuthService: aServ,
		Methods: map[string]bool{
			"/SecretCardService.CreateSecretCard":         true,
			"/SecretCardService.ReadSecretCard":           true,
			"/SecretCardService.UpdateSecretCard":         true,
			"/SecretCardService.DeleteSecretCard":         true,
			"/SecretCardService.ListSecretCard":           true,
			"/MetaSecretCardService.CreateMetaSecretCard": true,
			"/MetaSecretCardService.ReadMetaSecretCard":   true,
			"/MetaSecretCardService.UpdateMetaSecretCard": true,
			"/MetaSecretCardService.DeleteMetaSecretCard": true,
			"/MetaSecretCardService.ListMetaSecretCard":   true,
			"/FileSecretCardService.CreateFileSecretCard": true,
			"/FileSecretCardService.ReadFileSecretCard":   true,
			"/FileSecretCardService.UpdateFileSecretCard": true,
			"/FileSecretCardService.DeleteFileSecretCard": true,
			"/FileSecretCardService.ListFileSecretCard":   true,
			"/UserService.CreateUser":                     true,
			"/UserService.GetUser":                        true,
			"/SyncService.Sync":                           true,
		},
	}

	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			Err = err
		}
		s := grpc.NewServer()
		grpc.UnaryInterceptor(interceptor.Unary)
		pb.RegisterUserServiceServer(s, usrSrv)
		pb.RegisterSecretCardServiceServer(s, srv)
		pb.RegisterFileSecretCardServiceServer(s, fsrv)
		pb.RegisterMetaSecretCardServiceServer(s, msrv)
		pb.RegisterSyncServiceServer(s, ssrv)

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
