package healthz

import (
	basehandler "sorataskapi/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func CheckHealthz(c *gin.Context) {
	basehandler.ToResponseSuccess(c, "healthy")
}
