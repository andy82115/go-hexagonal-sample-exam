package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/handler/http"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
)

//	@title						Go Hexagonal Sample API
//	@version					1.0
//	@description				This is the exam for hexagonal architecture in go
//
//	@contact.name				andy82115
//	@contact.url				https://github.com/andy82115/go-hexagonal
//	@contact.email				andy82115@gmail.com
//
//	@license.name				MIT
//	@license.url				https://https://github.com/andy82115/go-hexagonal-sample-exam/LICENSE
//
//	@host						${HTTP_URL}
//	@BasePath					/v1
//	@schemes					http https
//
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and the access token.
func main() {
	// Load environment variables
	// 環境変数のロード
	config, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	secretRepo, err := service.NewSecretRepository(config.AWS)
	if err != nil {
		slog.Error("Error initializing secret repository", "error", err)
		os.Exit(1)
	}

	token, err := service.NewTokenService(config.Token, secretRepo)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	// ! Dependency Injection at Here
	// User Handler
	userRepo, err := service.NewUserRepository(config.DB)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}

	// Create UserService to handle web api
	userService := service.NewUserService(userRepo)
	// Main place to handle web api
	userHandler := http.NewUserHandler(userService)

	// Create Auth Service to handle login
	authService := service.NewAuthService(userRepo, token)
	// Handle login
	authHandler := http.NewAuthHandler(authService)

	// Init router
	router, err := http.NewRouter(
		config.HTTP,
		token,
		*userHandler,
		*authHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.URL, config.HTTP.Port)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
