package userservice

import (
	usermodel "sorataskapi/internal/model/user"
	userrepository "sorataskapi/internal/repository/user"
)

type UserService interface {
	LoginByEmail(payload usermodel.LoginByEmailRequest) (*usermodel.LoginByEmailResponse, error)
}

type userService struct {
	userRepo userrepository.UserRepository
}

func NewUserService(userRepo userrepository.UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (userSvc *userService) LoginByEmail(payload usermodel.LoginByEmailRequest) (*usermodel.LoginByEmailResponse, error) {
	return nil, nil
}
