package base

import (
	"net/http"
	"sorataskapi/internal/model/base"

	"github.com/gin-gonic/gin"
)

func ToResponseSuccess(c *gin.Context, data any) {
	c.AbortWithStatusJSON(http.StatusOK, &base.BaseResponse{
		Status: true,
		Data:   data,
	})
}

func ToResponseError(c *gin.Context, error base.BaseResponseError) {
	c.AbortWithStatusJSON(http.StatusOK, &base.BaseResponse{
		Status: false,
		Error:  error,
	})
}
