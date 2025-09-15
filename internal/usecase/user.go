package usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"github.com/adhyttungga/go-grpc-test/internal/repository"

	pb "github.com/adhyttungga/go-grpc-test/internal/pb"
)

type UserUsecaseImpl struct {
	Repository repository.UserRepository
}

type UserUsecase interface {
	Create(ctx context.Context, input *pb.CreateRequest) (*pb.CreateResponse, error)
	GetAll(ctx context.Context) (*pb.GetAllResponse, error)
	Update(ctx context.Context, input *pb.UpdateRequest) (*pb.UpdateResponse, error)
	Delete(ctx context.Context, input *pb.DeleteRequest) (*pb.DeleteResponse, error)
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{
		Repository: repo,
	}
}

func (u *UserUsecaseImpl) Create(ctx context.Context, input *pb.CreateRequest) (*pb.CreateResponse, error) {
	route := "/users/user"

	// Get user id from ctx
	userId, _ := ctx.Value("user_id").(int64)

	// Get section from ctx
	section, _ := ctx.Value("section").(string)

	// Get User from redis by session id
	user, err := u.Repository.GetUserCache(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Get role left join role right by role id and route
	role, err := u.Repository.GetRole(user.RoleId, route)
	if err != nil {
		return nil, err
	}

	// Return err if x-link-service header != role rights section
	roleSection, _ := (*role)["section"].(string)
	if !strings.EqualFold(section, roleSection) {
		return nil, errors.New("unauthorized access")
	}

	// Return err if r_created != 1
	rCreate, _ := (*role)["r_create"].(int)
	if rCreate != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Next: Hashed password

	// Create User
	_, err = u.Repository.CreateUser(
		entity.User{
			RoleId:   input.RoleId,
			Name:     input.Name,
			Email:    input.Email,
			Password: input.Password,
		},
	)
	if err != nil {
		return nil, err
	}

	result := pb.CreateResponse{
		Status:  true,
		Message: "Successfully",
	}
	return &result, nil
}

func (u *UserUsecaseImpl) GetAll(ctx context.Context) (*pb.GetAllResponse, error) {
	route := "/users/user"

	// Get user id from ctx
	userId, _ := ctx.Value("user_id").(int64)

	// Get section from ctx
	section, _ := ctx.Value("section").(string)

	// Get User from redis by session id
	user, err := u.Repository.GetUserCache(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Get role left join role right by role id and route
	role, err := u.Repository.GetRole(user.RoleId, route)
	if err != nil {
		return nil, err
	}

	// Return err if x-link-service header != role right section
	roleSection, _ := (*role)["section"].(string)
	if !strings.EqualFold(section, roleSection) {
		return nil, errors.New("unauthorized access")
	}

	// Return err if r_read != 1
	rRead, _ := (*role)["r_read"].(int)
	if rRead != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Get Users
	users, err := u.Repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var resp []*pb.UserDBResponse
	for _, val := range users {
		roleId, _ := val["role_id"].(int64)
		name, _ := val["name"].(string)
		email, _ := val["email"].(string)
		roleName, _ := val["role_name"].(string)
		lastAccess, _ := val["last_access"].(int64)

		t := time.UnixMilli(lastAccess)
		formattedTime := t.Format("2006-01-02 15:04:05")

		resp = append(resp, &pb.UserDBResponse{
			RoleId:     roleId,
			Name:       name,
			Email:      email,
			LastAccess: formattedTime,
			RoleName:   roleName,
		})
	}

	result := pb.GetAllResponse{
		Status:  true,
		Message: "Successfully",
		Data: &pb.DataResponse{
			Users: resp,
		},
	}

	return &result, nil
}

func (u *UserUsecaseImpl) Update(ctx context.Context, input *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	route := "/users/user"

	// Get user id from ctx
	userId, _ := ctx.Value("user_id").(int64)

	// Get section from ctx
	section, _ := ctx.Value("section").(string)

	// Get user fro redis by session id
	user, err := u.Repository.GetUserCache(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Get role left join role right by role id and route
	role, err := u.Repository.GetRole(user.RoleId, route)
	if err != nil {
		return nil, err
	}

	// Return err if x-link-service header != role right section
	roleSection, _ := (*role)["section"].(string)
	if !strings.EqualFold(section, roleSection) {
		return nil, errors.New("unauthorized access")
	}

	// Return err if r_update != 1
	rUpdate, _ := (*role)["r_update"].(int)
	if rUpdate != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Update User
	user.Name = input.Name
	_, err = u.Repository.UpdateUser(userId, *user)
	if err != nil {
		return nil, err
	}

	result := pb.UpdateResponse{
		Status:  true,
		Message: "Successfully",
	}

	return &result, nil
}

func (u *UserUsecaseImpl) Delete(ctx context.Context, input *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	route := "/users/user"

	// Get user id from ctx
	userId, _ := ctx.Value("user_id").(int64)

	// Get section from ctx
	section, _ := ctx.Value("section").(string)

	// Get User from redis by session id
	user, err := u.Repository.GetUserCache(ctx, userId)
	if err != nil {
		return nil, err
	}

	// Get role left join role right by role id and route
	role, err := u.Repository.GetRole(user.RoleId, route)
	if err != nil {
		return nil, err
	}

	// Return err if x-link-service header != role rights section
	roleSection, _ := (*role)["section"].(string)
	if !strings.EqualFold(section, roleSection) {
		return nil, errors.New("unauthorized access")
	}

	// Return err if r_delete != 1
	rDelete, _ := (*role)["r_delete"].(int)
	if rDelete != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Delete user
	if err := u.Repository.DeleteUser(input.Id); err != nil {
		return nil, err
	}

	result := pb.DeleteResponse{
		Status:  true,
		Message: "Successfully",
	}

	return &result, nil
}
