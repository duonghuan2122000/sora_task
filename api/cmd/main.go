// File chính của app
package main

import (
	"sorataskapi/config"
	"sorataskapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig(".")

	router := gin.Default()

	router.SetTrustedProxies(config.AppConfig.TrustProxies)

	routes.InitRoutes(router, config.AppConfig)

	router.Run(config.AppConfig.HostName + ":" + config.AppConfig.AppPort)
}
