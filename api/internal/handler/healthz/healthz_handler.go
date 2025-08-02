package healthz

import (
	"sorataskapi/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func CheckHealthz(c *gin.Context) {
	base.ToResponseSuccess(c, "healthy")
}
