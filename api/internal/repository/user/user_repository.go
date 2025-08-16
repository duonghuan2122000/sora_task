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
	result := database.MysqlConnect.Raw("SELECT id, email, passwordHashed, createdDate, updatedDate FROM soraUsers u WHERE u.email = ? LIMIT 1", email).Scan(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}
	return &user, nil
}

func (repo *userRepo) Insert(userEntity entity.UserEntity) error {
	userEntity.CreatedDate = time.Now().UTC()
	return database.MysqlConnect.Create(&userEntity).Error
}
