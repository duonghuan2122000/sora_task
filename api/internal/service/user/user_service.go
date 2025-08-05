package userservice

import (
	usermodel "sorataskapi/internal/model/user"
	userrepository "sorataskapi/internal/repository/user"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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
	userEntity, err := userSvc.userRepo.GetByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userEntity.PasswordHashed), []byte(payload.Password))
	if err != nil {
		return nil, err
	}

	claims := jwt.MapClaims{
		"exp": time.Now().UTC().Add(time.Hour * 1).Unix(),
		"sub": userEntity.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &usermodel.LoginByEmailResponse{
		AccessToken: accessToken,
	}, nil
}
