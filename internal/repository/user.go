package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// Rename method name
type UserRepository interface {
	CreateUser(user entity.User) (*entity.User, error)
	GetAllUsers() ([]map[string]any, error)
	UpdateUser(id int64, input entity.User) (*entity.User, error)
	DeleteUser(id int64) error
	GetUserCache(ctx context.Context, id int64) (*entity.User, error)
	GetRole(roleId int64, route string) (*map[string]any, error)
}

func NewUserRepository(db *gorm.DB, redisClient *redis.Client) UserRepository {
	return &UserRepositoryImpl{
		DB:          db,
		RedisClient: redisClient,
	}
}

func (r *UserRepositoryImpl) CreateUser(user entity.User) (*entity.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		log.Printf("error create new user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]map[string]any, error) {
	var users []map[string]any

	query := fmt.Sprintln(`
	SELECT 
		user.id AS id,
		user.role_id AS role_id,
		user.name AS name,
		user.email AS email,
		user.last_access AS last_access,
		role.name AS role_name,
	FROM
		user
	LEFT JOIN
		role ON role.id = user.role_id
	`)
	if err := r.DB.Raw(query).Scan(&users).Error; err != nil {
		log.Printf("error retrieve users: %v", err)
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) UpdateUser(id int64, input entity.User) (*entity.User, error) {
	var user entity.User

	err := r.DB.Model(&user).Where("id = ?", id).Clauses(clause.Returning{}).Updates(input).Error
	if err != nil {
		log.Printf("error update user: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) DeleteUser(id int64) error {
	err := r.DB.Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		log.Printf("error delete user: %v", err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserCache(ctx context.Context, id int64) (*entity.User, error) {
	var user entity.User

	strUser, err := r.RedisClient.Get(ctx, fmt.Sprintf("session: %d", id)).Result()
	if err != nil {
		log.Printf("error retrieve user from cache: %v", err)
		return nil, err
	}

	if err := json.Unmarshal([]byte(strUser), &user); err != nil {
		log.Printf("error parse user as JSON: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetRole(roleId int64, route string) (*map[string]any, error) {
	var role map[string]any

	query := fmt.Sprintf(`
	SELECT
		role.id,
		role.name,
		role.role_right_id,
		role_right.section,
		role_right.route,
		role_right.r_create,
		role_right.r_read,
		role_right.r_Update,
		role_right.r_delete,
	FROM 
		role
	LEFT JOIN
		role_right ON role_right.id = role.role_right_id
	WHERE role.id = %d AND role_right.route = '%s'
	`, roleId, route)

	if err := r.DB.Raw(query).Scan(&role).Error; err != nil {
		log.Printf("error retrieve role: %v", err)
		return nil, err
	}

	return &role, nil
}
