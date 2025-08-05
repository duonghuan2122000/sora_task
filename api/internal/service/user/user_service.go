package userservice

import (
	"sorataskapi/internal/entity"
	usermodel "sorataskapi/internal/model/user"
	userrepository "sorataskapi/internal/repository/user"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	LoginByEmail(payload usermodel.LoginByEmailRequest) (*usermodel.LoginByEmailResponse, error)
	RegisterUser(payload usermodel.RegisterUserRequest) error
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

func (userSvc *userService) RegisterUser(payload usermodel.RegisterUserRequest) error {
	passwordHased, _ := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	userEntity := entity.UserEntity{
		Id:             strings.ReplaceAll(uuid.New().String(), "-", ""),
		Email:          payload.Email,
		PasswordHashed: string(passwordHased),
	}
	userSvc.userRepo.Insert(userEntity)
	return nil
}
