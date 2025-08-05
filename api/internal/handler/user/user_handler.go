package userhandler

import (
	userrepository "sorataskapi/internal/repository/user"
	userservice "sorataskapi/internal/service/user"
)

var userSvc userservice.UserService

func InitHandler() {
	userRepo := userrepository.NewUserRepository()
	userSvc = userservice.NewUserService(userRepo)
}
