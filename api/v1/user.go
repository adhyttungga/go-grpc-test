package delivery

import (
	"context"
	"errors"

	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	Usecase usecase.UserUsecase
}

func NewUserServer(uc usecase.UserUsecase) *userServer {
	return &userServer{Usecase: uc}
}

func (s *userServer) Create(ctx context.Context, input *pb.CreateRequest) (*pb.CreateResponse, error) {
	result, err := s.Usecase.Create(ctx, input)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized access")) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return result, nil
}

func (s *userServer) GetAll(ctx context.Context, input *emptypb.Empty) (*pb.GetAllResponse, error) {
	result, err := s.Usecase.GetAll(ctx)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized access")) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return result, nil
}

func (s *userServer) Update(ctx context.Context, input *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	result, err := s.Usecase.Update(ctx, input)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized access")) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return result, nil
}

func (s *userServer) Delete(ctx context.Context, input *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	result, err := s.Usecase.Delete(ctx, input)
	if err != nil {
		if errors.Is(err, errors.New("unauthorized access")) {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return result, nil
}
