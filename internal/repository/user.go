package repository

import (
	"github.com/adhyttungga/go-grpc-test/internal/models/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

type UserRepository interface {
	Create(user entity.User) (*entity.User, error)
	GetAll() ([]entity.User, error)
	Update(id int64, input entity.User) (*entity.User, error)
	Delete(id int64) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

func (r *UserRepositoryImpl) Create(user entity.User) (*entity.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetAll() ([]entity.User, error) {
	var users []entity.User

	// Update query to include role name
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepositoryImpl) Update(id int64, input entity.User) (*entity.User, error) {
	var user entity.User

	err := r.DB.Model(&user).Where("id = ?", id).Clauses(clause.Returning{}).Updates(input).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) Delete(id int64) error {
	err := r.DB.Where("id = ?", id).Delete(&entity.User{}).Error

	if err != nil {
		return err
	}

	return nil
}
