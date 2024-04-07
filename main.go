package main

import (
	"context"
	"fmt"
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
	_, err = c.CreateUser(context.Background(), &pb.UserRequest{
		Login:    "test",
		Password: "test",
	})
	if err != nil {
		log.Println(err)
	}
	u, err := c.GetUser(context.Background(), &pb.UserRequest{
		Login:    "test",
		Password: "test",
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(u)
	}
	u, err = c.GetUser(context.Background(), &pb.UserRequest{
		Login:    "test",
		Password: "tost",
	})
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(u)
	}
	_, err = c.UpdatePassword(context.Background(), &pb.UserUpdatePasswordRequest{
		Login:       "test",
		OldPassword: "test",
		NewPassword: "tost",
	})
	if err != nil {
		fmt.Println("test")
		log.Println(err)
		_, err := c.UpdatePassword(context.Background(), &pb.UserUpdatePasswordRequest{
			Login:       "test",
			OldPassword: "tost",
			NewPassword: "test",
		})
		if err != nil {
			fmt.Println("tost")
			log.Println(err)
		}
	}
	client := pb.NewSecretCardServiceClient(conn)
	ctx := context.Background()
	respCreate, err := client.CreateSecretCard(ctx, &pb.CreateSecretCardRequest{
		UserId:   "711b9288-8aa2-4a4a-a689-a195d4f0b9e4",
		Name:     "MyCard",
		Url:      "https://example.com",
		Login:    "user",
		Password: "password",
		Text:     "Additional text",
	})
	if err != nil {
		log.Fatalf("failed to create secret card: %v", err)
	}
	log.Printf("Created secret card: %v", respCreate)

	// Пример вызова метода ReadSecretCard
	respRead, err := client.ReadSecretCard(ctx, &pb.ReadSecretCardRequest{
		Id: respCreate.Id,
	})
	if err != nil {
		log.Fatalf("failed to read secret card: %v", err)
	}
	log.Printf("Read secret card: %v", respRead)

	// Пример вызова метода UpdateSecretCard
	respUpdate, err := client.UpdateSecretCard(ctx, &pb.UpdateSecretCardRequest{
		Id:       respRead.Id,
		UserId:   respRead.UserId,
		Name:     "UpdatedName",
		Url:      "https://updated.com",
		Login:    "updateduser",
		Password: "updatedpassword",
		Text:     "Updated text",
	})
	if err != nil {
		log.Fatalf("failed to update secret card: %v", err)
	}
	log.Printf("Updated secret card: %v", respUpdate)

	// Пример вызова метода DeleteSecretCard
	// respDelete, err := client.DeleteSecretCard(ctx, &pb.DeleteSecretCardRequest{
	// 	Id: respUpdate.Id,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to delete secret card: %v", err)
	// }
	// log.Printf("Deleted secret card: %v", respDelete)

	// Пример вызова метода ListSecretCard
	respList, err := client.ListSecretCard(ctx, &pb.ListSecretCardRequest{
		UserId: "711b9288-8aa2-4a4a-a689-a195d4f0b9e4",
	})
	if err != nil {
		log.Fatalf("failed to list secret cards: %v", err)
	}
	for _, card := range respList.Cards {
		log.Printf("Listed secret card: %v", card)
	}

	fclient := pb.NewFileSecretCardServiceClient(conn)

	// Пример использования метода CreateFileSecretCard
	createReq := &pb.CreateFileSecretCardRequest{
		CardId: respUpdate.Id,           // Замените на реальные данные
		File:   []byte("file contents"), // Замените на реальные данные
	}
	createdFile, err := fclient.CreateFileSecretCard(context.Background(), createReq)
	if err != nil {
		log.Fatalf("failed to create file secret card: %v", err)
	}
	fmt.Printf("Created file secret card: %v\n", createdFile)

	// Пример использования метода ReadFileSecretCard
	readReq := &pb.ReadFileSecretCardRequest{
		Id: createdFile.Id, // Замените на реальные данные
	}
	readFile, err := fclient.ReadFileSecretCard(context.Background(), readReq)
	if err != nil {
		log.Fatalf("failed to read file secret card: %v", err)
	}
	fmt.Printf("Read file secret card: %v\n", readFile)

	updateReq := &pb.UpdateFileSecretCardRequest{
		Id:     createdFile.Id,                  // Замените на реальные данные
		CardId: respUpdate.Id,                   // Замените на реальные данные
		File:   []byte("updated file contents"), // Замените на реальные данные
	}
	updatedFile, err := fclient.UpdateFileSecretCard(context.Background(), updateReq)
	if err != nil {
		log.Fatalf("failed to update file secret card: %v", err)
	}
	fmt.Printf("Updated file secret card: %v\n", updatedFile)

	// Пример использования метода DeleteFileSecretCard
	// deleteReq := &pb.DeleteFileSecretCardRequest{
	// 	Id: createdFile.Id, // Замените на реальные данные
	// }
	// _, err = fclient.DeleteFileSecretCard(context.Background(), deleteReq)
	// if err != nil {
	// 	log.Fatalf("failed to delete file secret card: %v", err)
	// }
	// fmt.Println("Deleted file secret card")

	// Пример использования метода ListFileSecretCard
	listReq := &pb.ListFileSecretCardRequest{
		CardId: respUpdate.Id, // Замените на реальные данные
	}
	listResponse, err := fclient.ListFileSecretCard(context.Background(), listReq)
	if err != nil {
		log.Fatalf("failed to list file secret cards: %v", err)
	}
	fmt.Println("List of file secret cards:")
	for _, file := range listResponse.Files {
		fmt.Printf("- %v\n", file)
	}
}
