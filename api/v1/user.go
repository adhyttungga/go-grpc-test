package delivery

import (
	"context"

	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	Repository repository.UserRepository
	Ctx        context.Context
}

func NewUserServer(repo repository.UserRepository, ctx context.Context) *userServer {
	return &userServer{
		Repository: repo,
		Ctx:        ctx,
	}
}

func (s *userServer) Create(ctx context.Context, input *pb.CreateRequest) (*pb.CreateResponse, error) {
	// cek role right

	data := entity.User{
		RoleId:   input.RoleId,
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	_, err := s.Repository.Create(data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp := pb.CreateResponse{
		Status:  true,
		Message: "Successfully",
	}
	return &resp, nil
}

// func (s *userServer) GetAll(ctx context.Context) (*pb.GetAllResponse, error) {
// 	users, err := s.Repository.GetAll()
// 	if err != nil {
// 		return nil, status.Error(codes.InvalidArgument, err.Error())
// 	}
// 	var resp []*pb.UserDBResponse
// 	for _, val := range users {
// 		user := pb.UserDBResponse{
// 			RoleId:     val.RoleId,
// 			RoleName:   "",
// 			Name:       val.Name,
// 			Email:      val.Email,
// 			LastAccess: "",
// 		}

// 		resp = append(resp, &user)
// 	}

// 	result := pb.GetAllResponse{
// 		Status:  true,
// 		Message: "Successfully",
// 		Data: &pb.DataResponse{
// 			Users: resp,
// 		},
// 	}

// 	return &result, nil
// }
