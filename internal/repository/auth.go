package repository

import (
	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	DB *gorm.DB
}

type AuthRepository interface {
	Login(email string) (*entity.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		DB: db,
	}
}

func (r *AuthRepositoryImpl) Login(email string) (*entity.User, error) {
	var user entity.User

	err := r.DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
