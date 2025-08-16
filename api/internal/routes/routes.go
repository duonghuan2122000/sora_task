package routes

import (
	"net/http"
	"sorataskapi/config"
	"sorataskapi/internal/handler/healthz"
	userhandler "sorataskapi/internal/handler/user"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SecurityMiddleware(appConfig config.Config) gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; script-src 'self'",
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		IsDevelopment:         false,
	})

	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Next()
	}
}

func InitRoutes(router *gin.Engine, appConfig config.Config) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     appConfig.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "Origin", "Host"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))

	router.Use(SecurityMiddleware(appConfig))

	userhandler.InitHandler()

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)

	apiV1.POST("/users/login/by-mail", userhandler.LoginByEmail)
	apiV1.POST("/users/register", userhandler.RegisterUser)
}
