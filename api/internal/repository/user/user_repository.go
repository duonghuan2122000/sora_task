package userrepository

import (
	"sorataskapi/internal/database"
	"sorataskapi/internal/entity"
	"time"
)

type UserRepository interface {
	GetByEmail(email string) (*entity.UserEntity, error)
	Insert(userEntity entity.UserEntity) error
}

type userRepo struct {
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (repo *userRepo) GetByEmail(email string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	if err := database.MysqlConnect.Where(&entity.UserEntity{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *userRepo) Insert(userEntity entity.UserEntity) error {
	userEntity.CreatedDate = time.Now().UTC()
	return database.MysqlConnect.Create(&userEntity).Error
}
