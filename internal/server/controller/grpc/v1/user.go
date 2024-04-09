package v1

import (
	"context"

	"github.com/Kotletta-TT/GophKeeper/internal/server/service"
	"github.com/Kotletta-TT/GophKeeper/internal/server/usecase/user"
	pb "github.com/Kotletta-TT/GophKeeper/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	Ruc         user.Register
	Guc         user.GeterUser
	Cuc         user.ChangePassword
	AuthService service.AuthService
}

func NewUserServiceServer(ruc user.Register, guc user.GeterUser, cuc user.ChangePassword, authService service.AuthService) *UserServer {
	return &UserServer{
		Ruc:         ruc,
		Guc:         guc,
		Cuc:         cuc,
		AuthService: authService,
	}
}

func (s *UserServer) CreateUser(
	ctx context.Context,
	in *pb.UserRequest) (*pb.CreateUserResponse, error) {
	var resp pb.CreateUserResponse
	if err := s.Ruc.Register(ctx, in.Login, in.Password); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserServer) GetUser(
	ctx context.Context,
	in *pb.UserRequest) (*pb.GetUserResponse, error) {
	usr, err := s.Guc.GetUser(ctx, in.Login, in.Password)
	if err != nil {
		return nil, err
	}
	u := pb.UserResponse{
		Login: usr.Login,
		Id:    usr.ID.String(),
	}
	return &pb.GetUserResponse{
		User: &u,
	}, nil
}

func (s *UserServer) UpdatePassword(
	ctx context.Context,
	in *pb.UserUpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	var resp pb.UpdatePasswordResponse
	if err := s.Cuc.ChangePassword(ctx, in.Login, in.OldPassword, in.NewPassword); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (s *UserServer) Login(
	ctx context.Context,
	in *pb.UserRequest) (*pb.LoginResponse, error) {
	var resp pb.LoginResponse
	usr, err := s.Guc.GetUser(ctx, in.Login, in.Password)
	if err != nil {
		return nil, err
	}
	token, err := s.AuthService.GenerateToken(usr.ID.String())
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return &resp, nil
}
