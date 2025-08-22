package routes

import (
	"net/http"
	"sorataskapi/config"
	"sorataskapi/internal/handler/healthz"
	tenanthandler "sorataskapi/internal/handler/tenant"
	userhandler "sorataskapi/internal/handler/user"
	"sorataskapi/internal/middleware"
	tenantrepository "sorataskapi/internal/repository/tenant"
	tenantuserrepository "sorataskapi/internal/repository/tenant_user"
	userrepository "sorataskapi/internal/repository/user"
	tenantservice "sorataskapi/internal/service/tenant"
	userservice "sorataskapi/internal/service/user"
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

	// đăng ký repository
	userRepo := userrepository.NewUserRepository()
	tenantRepo := tenantrepository.NewTenantRepository()
	tenantUserRepo := tenantuserrepository.NewTenantUserRepository()

	// đăng ký service
	userSvc := userservice.NewUserService(userRepo)
	tenantSvc := tenantservice.NewTenantService(tenantRepo, tenantUserRepo)

	// đăng ký handler
	userhandler.UserSvc = userSvc
	tenanthandler.TenantSvc = tenantSvc

	apiV1 := router.Group("/v1")

	apiV1.GET("/healthz", healthz.CheckHealthz)

	userRouter := apiV1.Group("/users")

	userRouter.POST("/login/by-mail", userhandler.LoginByEmail)
	userRouter.POST("/register", userhandler.RegisterUser)

	tenantRouter := apiV1.Group("/tenants")
	tenantRouter.Use(middleware.AuthMiddleware())
	tenantRouter.GET("/me", tenanthandler.GetListByUser)
	tenantRouter.POST("/create", tenanthandler.Create)
	tenantRouter.GET("/current", middleware.TenantMiddleware(), tenanthandler.GetCurrent)
	tenantRouter.POST("/select", tenanthandler.Select)
}
