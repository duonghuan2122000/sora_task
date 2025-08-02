package routes

import (
	"net/http"
	"sorataskapi/internal/handler/healthz"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func SecurityMiddleware() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:          []string{"localhost:8080"},
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

func InitRoutes(router *gin.Engine) {
	router.Use(SecurityMiddleware())

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)
}
