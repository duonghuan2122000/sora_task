// File chính của app
package main

import (
	"fmt"
	"sorataskapi/config"
	"sorataskapi/internal/database"
	"sorataskapi/internal/entity"
	"sorataskapi/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig(".")
	database.InitMysql(config.AppConfig.MysqlConnectionString)

	err := database.MysqlConnect.AutoMigrate(&entity.UserEntity{})
	if err != nil {
		fmt.Printf("Error migrate: %v", err)
		return
	}

	router := gin.Default()
	router.SetTrustedProxies(config.AppConfig.TrustProxies)
	routes.InitRoutes(router, config.AppConfig)
	router.Run(config.AppConfig.HostName + ":" + config.AppConfig.AppPort)
}
