package basehandler

import (
	"net/http"
	basemodel "sorataskapi/internal/model/base"

	"github.com/gin-gonic/gin"
)

func ToResponseSuccess(c *gin.Context, data any) {
	c.AbortWithStatusJSON(http.StatusOK, &basemodel.BaseResponse{
		Status: true,
		Data:   data,
	})
}

func ToResponseError(c *gin.Context, error basemodel.BaseResponseError) {
	c.AbortWithStatusJSON(http.StatusOK, &basemodel.BaseResponse{
		Status: false,
		Error:  error,
	})
}
