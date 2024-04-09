package app

import (
	config "github.com/Kotletta-TT/GophKeeper/config/server"
	v1 "github.com/Kotletta-TT/GophKeeper/internal/server/controller/grpc/v1"
	"github.com/Kotletta-TT/GophKeeper/internal/server/controller/grpc/v1/interceptors"
	filesecretcard "github.com/Kotletta-TT/GophKeeper/internal/server/repository/postgres/file_secretcard"
	metasecretcard "github.com/Kotletta-TT/GophKeeper/internal/server/repository/postgres/meta_secretcard"
	"github.com/Kotletta-TT/GophKeeper/internal/server/repository/postgres/secretcard"
	repouser "github.com/Kotletta-TT/GophKeeper/internal/server/repository/postgres/user"
	"github.com/Kotletta-TT/GophKeeper/internal/server/service"
	"github.com/Kotletta-TT/GophKeeper/internal/server/usecase"
	"github.com/Kotletta-TT/GophKeeper/internal/server/usecase/user"
	"github.com/Kotletta-TT/GophKeeper/internal/server/utils"
	"github.com/Kotletta-TT/GophKeeper/logger"
	"github.com/Kotletta-TT/GophKeeper/pkg/postgres"
	pb "github.com/Kotletta-TT/GophKeeper/proto"
	"google.golang.org/grpc"

	"context"
	"net"
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
			"/v1.SecretCardService/CreateSecretCard":         true,
			"/v1.SecretCardService/ReadSecretCard":           true,
			"/v1.SecretCardService/UpdateSecretCard":         true,
			"/v1.SecretCardService/DeleteSecretCard":         true,
			"/v1.SecretCardService/ListSecretCard":           true,
			"/v1.MetaSecretCardService/CreateMetaSecretCard": true,
			"/v1.MetaSecretCardService/ReadMetaSecretCard":   true,
			"/v1.MetaSecretCardService/UpdateMetaSecretCard": true,
			"/v1.MetaSecretCardService/DeleteMetaSecretCard": true,
			"/v1.MetaSecretCardService/ListMetaSecretCard":   true,
			"/v1.FileSecretCardService/CreateFileSecretCard": true,
			"/v1.FileSecretCardService/ReadFileSecretCard":   true,
			"/v1.FileSecretCardService/UpdateFileSecretCard": true,
			"/v1.FileSecretCardService/DeleteFileSecretCard": true,
			"/v1.FileSecretCardService/ListFileSecretCard":   true,
			"/v1.UserService/GetUser":                        true,
			"/v1.SyncService/Sync":                           true,
		},
	}

	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			Err = err
		}
		s := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary))
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
