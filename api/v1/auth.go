package delivery

import (
	"context"
	"fmt"

	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
	Usecase usecase.AuthUsecase
}

func NewAuthServer(uc usecase.AuthUsecase) *authServer {
	return &authServer{
		Usecase: uc,
	}
}

func (s *authServer) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := s.Usecase.Login(ctx, input)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Set authorization header
	header := metadata.Pairs(
		"authorization", fmt.Sprintf("Bearer %s", result.Data.AccessToken),
	)

	_ = grpc.SetHeader(ctx, header)

	return result, nil
}

func (s *authServer) Logout(ctx context.Context, input *emptypb.Empty) (*pb.LogoutResponse, error) {
	result, err := s.Usecase.Logout(ctx)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Empty authorization header
	header := metadata.Pairs(
		"authorization", "",
	)

	_ = grpc.SetHeader(ctx, header)

	return result, nil
}
