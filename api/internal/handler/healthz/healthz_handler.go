package healthz

import (
	base_handler "sorataskapi/internal/handler/base"

	"github.com/gin-gonic/gin"
)

func CheckHealthz(c *gin.Context) {
	base_handler.ToResponseSuccess(c, "healthy")
}
