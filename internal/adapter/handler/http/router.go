package http

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/port"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	docs "github.com/andy82115/go-hexagonal-sample-exam/docs"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	*gin.Engine
}

func NewRouter(
	config *config.HTTP,
	token port.TokenService,
	userHandler UserHandler,
	authHandler AuthHandler,
) (*Router, error) {
	// Disable debug mode in production
	if config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Handle CORS. Will define which url can pass through.
	// CORSを処理する。　どのURLを通過させるかを設定する。
	ginConfig := cors.DefaultConfig()
	allowedOrigins := config.AllowedOrigins
	originsList := strings.Split(allowedOrigins, ",")
	ginConfig.AllowOrigins = originsList

	router := gin.New()
	router.Use(sloggin.New(slog.Default()), gin.Recovery(), cors.New(ginConfig))

	// Custom validators
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		// !hardcode: user_role is the UserRole in json key for validator
		if err := v.RegisterValidation("user_role", userRoleValidator); err != nil {
			return nil, err
		}
	}

	// set up swagger dynamically from config (.env)
	// 設定ファイルかを利用して、swaggerのURLを調整する。
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", config.URL, config.Port)
	// Swagger Bind
	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/", userHandler.Register)
			user.POST("/login", authHandler.Login)

			authUser := user.Group("/").Use(authMiddleware(token))
			{
				authUser.GET("/", userHandler.ListUsers)
				authUser.GET("/:id", userHandler.GetUser)

				admin := authUser.Use(adminMiddleware())
				{
					admin.PUT("/:id", userHandler.UpdateUser)
					admin.DELETE("/:id", userHandler.DeleteUser)
				}
			}
		}
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
