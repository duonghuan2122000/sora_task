package tenanthandler

import (
	"errors"
	basehandler "sorataskapi/internal/handler/base"
	basemodel "sorataskapi/internal/model/base"
	tenantmodel "sorataskapi/internal/model/tenant"
	usermodel "sorataskapi/internal/model/user"
	tenantservice "sorataskapi/internal/service/tenant"

	"github.com/gin-gonic/gin"
)

var TenantSvc tenantservice.TenantService

func GetListByUser(c *gin.Context) {
	currentUserDto, ok := usermodel.FromContext(c.Request.Context())
	if !ok {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "401",
			Message: "Unauthorized",
		})
		return
	}
	tenants, err := TenantSvc.GetListByUser(currentUserDto.UserId)
	if err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "999",
			Message: "Thất bại",
		})
		return
	}
	basehandler.ToResponseSuccess(c, tenants)
}

func Create(c *gin.Context) {
	var payload basemodel.BaseRequest[tenantmodel.CreateTenantReqDto]
	if err := c.ShouldBindJSON(&payload); err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "400",
			Message: "Tham số không hợp lệ",
		})
		return
	}

	result, err := TenantSvc.Create(c.Request.Context(), payload.Data.Attributes)
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
	basehandler.ToResponseSuccess(c, result)
}

func GetCurrent(c *gin.Context) {
	result, err := TenantSvc.GetCurrent(c.Request.Context())
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
	basehandler.ToResponseSuccess(c, result)
}

func Select(c *gin.Context) {
	var payload basemodel.BaseRequest[tenantmodel.SelectTenantReqDto]
	if err := c.ShouldBindJSON(&payload); err != nil {
		basehandler.ToResponseError(c, basemodel.BaseResponseError{
			Code:    "400",
			Message: "Tham số không hợp lệ",
		})
		return
	}

	err := TenantSvc.Select(c.Request.Context(), payload.Data.Attributes)
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
	c.SetCookie(basemodel.CookieTenantId, payload.Data.Attributes.Id, 10*24*60*60, "", "", false, true)
	basehandler.ToResponseSuccess(c, true)
}
