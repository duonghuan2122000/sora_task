// File chính của app
package main

import (
	"sorataskapi/internal/handler/healthz"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)

	router.Run("127.0.0.1:8080")
}
