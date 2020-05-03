package main

import (
	"context"
	"log"
	"os"

	pb "github.com/fusidic/user-service/proto/user"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
)

func main() {
	srv := micro.NewService(
		micro.Name("user-cli"),
		micro.Version("latest"),
	)

	// 初始化并解析命令行参数
	srv.Init()

	client := pb.NewUserServiceClient("user", microclient.DefaultClient)

	name := "Ewan Valentine"
	email := "ewan.valentine89@gmail.com"
	password := "test123"
	company := "BBC"

	r, err := client.Create(context.TODO(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.TODO(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v \n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// 退出
	os.Exit(0)
}
