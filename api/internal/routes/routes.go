package routes

import (
	"net/http"
	"sorataskapi/config"
	"sorataskapi/internal/handler/healthz"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SecurityMiddleware(appConfig config.Config) gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          appConfig.AllowedOrigins,
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
	router.Use(SecurityMiddleware(appConfig))

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)
}
