// File chính của app
package main

import (
	"sorataskapi/internal/handler/healthz"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)

	router.Run(":8080")
}
