package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/mocks"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupTest(t *testing.T) (*mocks.UserRepository, *mocks.TokenService, *service.AuthService) {
	mockUserRepo := mocks.NewUserRepository(t)
	mockTokenService := mocks.NewTokenService(t)
	authService := service.NewAuthService(mockUserRepo, mockTokenService)
	return mockUserRepo, mockTokenService, authService
}

func TestAuthService_Login_Success(t *testing.T) {
	// Setup
	mockUserRepo, mockTokenService, authService := setupTest(t)

	hashedPassword, _ := util.HashPassword("123456789")
	
	mockUser := &domain.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}
	expectedToken := "valid.jwt.token"

	// Set expectations
	mockUserRepo.On("GetUserByEmail", mock.Anything, "test@example.com").
		Return(mockUser, nil)
	mockTokenService.On("CreateToken", mockUser).
		Return(expectedToken, nil)

	// Execute
	token, err := authService.Login(context.Background(), "test@example.com", "123456789")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, token)

	mockUserRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuthService_Login_UserNotFound(t *testing.T) {
	// Setup
	mockUserRepo, mockTokenService, authService := setupTest(t)

	// Set expectations
	mockUserRepo.On("GetUserByEmail", mock.Anything, "nonexistent@example.com").
		Return(nil, domain.ErrDataNotFound)

	// Execute
	token, err := authService.Login(context.Background(), "nonexistent@example.com", "anypassword")

	// Assert
	assert.Equal(t, "", token)
	assert.Equal(t, domain.ErrDataNotFound, err)
	mockUserRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuthService_Login_InternalError(t *testing.T) {
	// Setup
	mockUserRepo, mockTokenService, authService := setupTest(t)

	// Set expectations
	mockUserRepo.On("GetUserByEmail", mock.Anything, "test@example.com").
		Return(nil, errors.New("database error"))

	// Execute
	token, err := authService.Login(context.Background(), "test@example.com", "password")

	// Assert
	assert.Equal(t, "", token)
	assert.Equal(t, domain.ErrInternal, err)
	mockUserRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuthService_Login_InvalidCredentials(t *testing.T) {
	// Setup
	mockUserRepo, mockTokenService, authService := setupTest(t)

	mockUser := &domain.User{
		Email:    "test@example.com",
		Password: "$2a$10$h.dl5J86rGH7I8bD9bZeZe",
	}

	// Set expectations
	mockUserRepo.On("GetUserByEmail", mock.Anything, "test@example.com").
		Return(mockUser, nil)

	// Execute
	token, err := authService.Login(context.Background(), "test@example.com", "wrongpassword")

	// Assert
	assert.Equal(t, "", token)
	assert.Equal(t, domain.ErrInvalidCredentials, err)
	mockUserRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}

func TestAuthService_Login_TokenCreationError(t *testing.T) {
	// Setup
	mockUserRepo, mockTokenService, authService := setupTest(t)

	hashedPassword, _ := util.HashPassword("123456789")

	mockUser := &domain.User{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	// Set expectations
	mockUserRepo.On("GetUserByEmail", mock.Anything, "test@example.com").
		Return(mockUser, nil)
	mockTokenService.On("CreateToken", mockUser).
		Return("", errors.New("token creation failed"))

	// Execute
	token, err := authService.Login(context.Background(), "test@example.com", "123456789")

	// Assert
	assert.Equal(t, "", token)
	assert.Equal(t, domain.ErrTokenCreation, err)
	mockUserRepo.AssertExpectations(t)
	mockTokenService.AssertExpectations(t)
}