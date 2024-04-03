package main

import (
	"context"
	"log"

	pb "GophKeeper-Server/internal/controller/grpc/v1/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial(":3200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	c.CreateUser(context.Background(), &pb.CreateUserRequest{
		Login:    "test",
		Password: "test",
	})
}
