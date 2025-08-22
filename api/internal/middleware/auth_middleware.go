package middleware

import (
	"errors"
	"sorataskapi/config"
	basehandler "sorataskapi/internal/handler/base"
	basemodel "sorataskapi/internal/model/base"
	tenantmodel "sorataskapi/internal/model/tenant"
	usermodel "sorataskapi/internal/model/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie(basemodel.CookieAccessToken)
		if err != nil {
			basehandler.ToResponseError(c, basemodel.BaseResponseError{
				Code:    "401",
				Message: "Unauthorized",
			})
			return
		}
		token, err := jwt.ParseWithClaims(accessToken, &usermodel.CurrentUserDto{}, func(t *jwt.Token) (any, error) {
			// Bảo đảm thuật toán đúng
			if t.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return []byte(config.AppConfig.JwtSecret), nil
		}, // Cho phép lệch giờ nhỏ khi check exp/nbf/iAT
			jwt.WithLeeway(30*time.Second))

		if err != nil || !token.Valid {
			basehandler.ToResponseError(c, basemodel.BaseResponseError{
				Code:    "401",
				Message: "Unauthorized",
			})
			return
		}
		// Lưu claims vào context để handler sau dùng
		if claims, ok := token.Claims.(*usermodel.CurrentUserDto); ok {
			c.Set(basemodel.CurrentUserInfoKey, claims)
			ctx := usermodel.IntoContext(c.Request.Context(), claims)
			c.Request = c.Request.WithContext(ctx)
		}
		c.Next()
	}
}

func TenantMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tenantId, err := c.Cookie(basemodel.CookieTenantId)
		if err != nil {
			basehandler.ToResponseError(c, basemodel.BaseResponseError{
				Code:    "204",
				Message: "Chưa chọn tenant",
			})
			return
		}
		c.Set(basemodel.CurrentTenantIdKey, tenantId)
		ctx := tenantmodel.IntoContext(c.Request.Context(), tenantId)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
