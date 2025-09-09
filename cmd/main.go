package main

import (
	"context"
	"log"
	"net"

	delivery "github.com/adhyttungga/go-grpc-test/api/v1"
	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"github.com/adhyttungga/go-grpc-test/pkg/config"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8081",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("error to connect redis: %v", err)
	}

	log.Println("Connected to redis: ", pong)

	db, err := config.DBInit()
	if err != nil {
		log.Fatalf("error to connect db: %v", err)
	}

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("error to listen: %v", err)
	}

	defer listen.Close()

	grpcServer := grpc.NewServer()

	// create new metadata
	md := metadata.New(map[string]string{
		"Headers":       "X-Link-Service",
		"Authorization": "Bearer Token",
	})

	ctxWithMetadata := metadata.NewOutgoingContext(context.Background(), md)

	// register
	auth_repo := repository.NewAuthRepository(db)
	user_repo := repository.NewUserRepository(db)

	auth_delivery := delivery.NewAuthServer(auth_repo, ctxWithMetadata)
	user_delivery := delivery.NewUserServer(user_repo, ctxWithMetadata)

	pb.RegisterAuthServiceServer(grpcServer, auth_delivery)
	pb.RegisterUserServiceServer(grpcServer, user_delivery)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("error to serve: %v", err)
	}
}
