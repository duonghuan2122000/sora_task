// File chính của app
package main

import (
	"sorataskapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.SetTrustedProxies([]string{"127.0.0.1"})

	routes.InitRoutes(router)

	router.Run("127.0.0.1:8080")
}
