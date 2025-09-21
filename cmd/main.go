package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	delivery "github.com/adhyttungga/go-grpc-test/api/v1"
	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"github.com/adhyttungga/go-grpc-test/internal/usecase"
	"github.com/adhyttungga/go-grpc-test/pkg/config"
	"github.com/adhyttungga/go-grpc-test/pkg/middleware"
	"google.golang.org/grpc"
)

func main() {
	// Initialize redis client
	client, err := config.RedisInit(context.Background())
	if err != nil {
		log.Fatalf("error to initialize redis client: %v", err)
	}

	// Initialize database connection
	db, err := config.DBInit()
	if err != nil {
		log.Fatalf("error to initialize db connection: %v", err)
	}

	// Auto-migrate the models
	db.AutoMigrate(&entity.User{}, &entity.Role{}, &entity.RoleRight{})

	// Initialize repository
	authRepo := repository.NewAuthRepository(db, client)
	userRepo := repository.NewUserRepository(db, client)

	// Initialize usecase
	authUsecase := usecase.NewAuthUsecase(authRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	// Initialize delivery
	authDelivery := delivery.NewAuthServer(authUsecase)
	userDelivery := delivery.NewUserServer(userUsecase)

	// Initialize tcp connection
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", config.Config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer listen.Close()

	// Chain unary interceptor
	chainedUI := grpc.ChainUnaryInterceptor(
		middleware.LoggingInterceptor,
		middleware.AuthInterceptor,
	)

	// Initialize new gRPC server
	grpcServer := grpc.NewServer(
		chainedUI,
	)

	// Register service server
	pb.RegisterAuthServiceServer(grpcServer, authDelivery)
	pb.RegisterUserServiceServer(grpcServer, userDelivery)

	// Gracefully shutdown
	go func() {
		// Service connection
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	log.Printf("Listening and serving.")

	// Wait for interrupt signal to gracefully shutdown the server
	// with timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT (Ctrl + C)
	// kill -9 is syscall.SIGKILL but can't be catch
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server...")

	grpcServer.GracefulStop()

	log.Println("User service shutdown completed gracefully")
}
