package userhandler

import (
	"errors"
	basehandler "sorataskapi/internal/handler/base"
	basemodel "sorataskapi/internal/model/base"
	usermodel "sorataskapi/internal/model/user"
	userservice "sorataskapi/internal/service/user"

	"github.com/gin-gonic/gin"
)

var UserSvc userservice.UserService

func LoginByEmail(c *gin.Context) {
	var payload basemodel.BaseRequest[usermodel.LoginByEmailRequest]
	if err := c.ShouldBindJSON(&payload); err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "400",
			Message: "Tham số không hợp lệ",
		})
		return
	}

	result, err := UserSvc.LoginByEmail(payload.Data.Attributes)
	if err != nil {
		var loginErr *basemodel.LogicError
		if errors.As(err, &loginErr) {
			basehandler.ToResponseError(c, basemodel.BaseResponseError{
				Code:    loginErr.Code,
				Message: loginErr.Message,
			})
			return
		}
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "999",
			Message: "Thất bại",
		})
		return
	}

	// thiết lập access token vào cookie
	c.SetCookie(basemodel.CookieAccessToken, result.AccessToken, result.ExpiresIn, "", "", false, true)

	basehandler.ToResponseSuccess(c, result)
}

func RegisterUser(c *gin.Context) {
	var payload basemodel.BaseRequest[usermodel.RegisterUserRequest]
	if err := c.ShouldBindJSON(&payload); err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "400",
			Message: "Tham số không hợp lệ",
		})
		return
	}
	if err := UserSvc.RegisterUser(payload.Data.Attributes); err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "999",
			Message: "Thất bại",
		})
		return
	}
	basehandler.ToResponseSuccess(c, nil)
}
