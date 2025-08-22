package usermodel

import (
	"context"
	basemodel "sorataskapi/internal/model/base"

	"github.com/golang-jwt/jwt/v5"
)

type LoginByEmailRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginByEmailResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CurrentUserDto struct {
	Exp    int64  `json:"exp"`
	UserId string `json:"sub"`
	jwt.RegisteredClaims
}

func IntoContext(ctx context.Context, u *CurrentUserDto) context.Context {
	return context.WithValue(ctx, basemodel.CurrentUserInfoKey, u)
}

func FromContext(ctx context.Context) (*CurrentUserDto, bool) {
	v := ctx.Value(basemodel.CurrentUserInfoKey)
	u, ok := v.(*CurrentUserDto)
	return u, ok
}
