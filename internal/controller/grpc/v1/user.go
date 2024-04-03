package v1

import (
	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"
	"GophKeeper-Server/internal/usecase/user"
	"context"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Uc user.Register
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var resp pb.CreateUserResponse
	resp.Error = s.Uc.Register(ctx, in.Login, in.Password).Error()
	return &resp, nil
}
