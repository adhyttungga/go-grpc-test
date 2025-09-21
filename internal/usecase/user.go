package usecase

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"github.com/adhyttungga/go-grpc-test/internal/repository"
	"golang.org/x/crypto/bcrypt"

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
	userId, _ := ctx.Value("user_id").(string)

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
	rCreate := int((*role)["r_create"].(int64))
	if rCreate != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Generate hash password
	hp, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		log.Printf("error creating hash password")
		return nil, err
	}

	// Create User
	_, err = u.Repository.CreateUser(
		entity.User{
			RoleId:   input.RoleId,
			Name:     input.Name,
			Email:    input.Email,
			Password: string(hp),
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
	userId, _ := ctx.Value("user_id").(string)

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
	rRead := int((*role)["r_read"].(int64))
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
	userId, _ := ctx.Value("user_id").(string)

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
	rUpdate := int((*role)["r_update"].(int64))
	if rUpdate != 1 {
		return nil, errors.New("unauthorized access")
	}

	// Update User
	userIdI64, _ := strconv.ParseInt(userId, 10, 64)
	user.Name = input.Name
	_, err = u.Repository.UpdateUser(userIdI64, *user)
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
	userId, _ := ctx.Value("user_id").(string)

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
	rDelete := int((*role)["r_delete"].(int64))
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
