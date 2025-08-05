package userrepository

import (
	"sorataskapi/internal/database"
	"sorataskapi/internal/entity"
)

type UserRepository interface {
	GetByEmail(email string) (*entity.UserEntity, error)
}

type userRepo struct {
}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (repo *userRepo) GetByEmail(email string) (*entity.UserEntity, error) {
	var user entity.UserEntity
	if err := database.MysqlConnect.Model(entity.UserEntity{Email: email}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
