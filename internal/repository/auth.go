package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB          *gorm.DB
	RedisClient *redis.Client
}

// rename login to GetUserByEmail
type AuthRepository interface {
	GetUserByEmail(email string) (*entity.User, error)
	StoreUserCache(ctx context.Context, userId int64, user entity.User) error
	DeleteUserCache(ctx context.Context, userId int64) error
}

func NewAuthRepository(db *gorm.DB, redisClient *redis.Client) AuthRepository {
	return &AuthRepositoryImpl{
		DB:          db,
		RedisClient: redisClient,
	}
}

func (r *AuthRepositoryImpl) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		log.Printf("error retrieve user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepositoryImpl) StoreUserCache(ctx context.Context, userId int64, user entity.User) error {
	byteUser, err := json.Marshal(user)
	if err != nil {
		log.Printf("error convert user: %v", err)
		return err
	}

	if err := r.RedisClient.Set(ctx, fmt.Sprintf("session: %d", userId), byteUser, 0).Err(); err != nil {
		log.Printf("error store user to cache: %v", err)
		return err
	}

	return nil
}

func (r *AuthRepositoryImpl) DeleteUserCache(ctx context.Context, userId int64) error {
	if err := r.RedisClient.Del(ctx, fmt.Sprintf("session: %d", userId)).Err(); err != nil {
		log.Printf("error delete user from cache: %v", err)
		return err
	}

	return nil
}
