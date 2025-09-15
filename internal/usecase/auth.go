package usecase

import (
	"context"
	"errors"
	"strings"

	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"github.com/adhyttungga/go-grpc-test/pkg/utils"
)

type AuthUsecaseImpl struct {
	Repository repository.AuthRepository
}

type AuthUsecase interface {
	Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error)
	Logout(ctx context.Context) (*pb.LogoutResponse, error)
}

func NewAuthUsecase(repo repository.AuthRepository) AuthUsecase {
	return &AuthUsecaseImpl{
		Repository: repo,
	}
}

func (u *AuthUsecaseImpl) Login(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error) {
	// Retrieve user by email
	user, err := u.Repository.GetUserByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	// Validate password
	// Next: Hashed password
	if !strings.EqualFold(user.Password, input.Password) {
		return nil, errors.New("invalid credential")
	}

	// Next: Update user last access

	// Set user session
	err = u.Repository.StoreUserCache(ctx, user.Id, *user)
	if err != nil {
		return nil, err
	}

	// Next: Generate refresh token with JWT
	accessToken, err := utils.GenerateToken(user.Id)
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		Status:  true,
		Message: "Successfully",
		Data: &pb.TokenResponse{
			AccessToken: accessToken,
		},
	}, nil
}

func (u *AuthUsecaseImpl) Logout(ctx context.Context) (*pb.LogoutResponse, error) {
	// Get user id from ctx
	userId, _ := ctx.Value("user_id").(int64)

	if err := u.Repository.DeleteUserCache(ctx, userId); err != nil {
		return nil, err
	}

	return &pb.LogoutResponse{
		Status:  true,
		Message: "Successfully",
	}, nil
}
