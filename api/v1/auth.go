package delivery

import (
	"context"
	"fmt"
	"strings"

	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	Repository repository.AuthRepository
	Ctx        context.Context
}

func NewAuthServer(repo repository.AuthRepository, ctx context.Context) *authServer {
	return &authServer{
		Repository: repo,
		Ctx:        ctx,
	}
}

func (s *authServer) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {

	user, err := s.Repository.Login(input.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if !strings.EqualFold(user.Password, input.Password) {
		return nil, status.Error(codes.Unauthenticated, fmt.Sprintln("invalid credential"))
	}

	// set user to redis here

	// create access token here

	result := pb.LoginResponse{
		Status:  true,
		Message: "Successfully",
		Data: &pb.TokenResponse{
			AccessToken: "",
		},
	}

	return &result, nil
}
