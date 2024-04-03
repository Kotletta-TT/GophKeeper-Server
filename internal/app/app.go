package app

import (
	"GophKeeper-Server/config"
	v1 "GophKeeper-Server/internal/controller/grpc/v1"
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	repouser "GophKeeper-Server/internal/repository/postgres/user"
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
	pg.Pool.Exec(ctx, "CREATE TABLE IF NOT EXISTS users (id uuid PRIMARY KEY, login text, password text)")
	ur := repouser.NewUserRepositroy(pg)
	uc := user.NewRegisterUC(l, ur, utils.HashPassword)
	go func() {
		listen, err := net.Listen("tcp", ":3200")
		if err != nil {
			Err = err
		}
		s := grpc.NewServer()
		pb.RegisterUserServiceServer(s, &v1.UserServer{Uc: uc})
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
