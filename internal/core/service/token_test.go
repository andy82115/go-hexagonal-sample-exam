package service_test

import (
	"strings"
	"testing"
	"time"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/mocks"
	"github.com/stretchr/testify/assert"
)

func TestTokenService_CreateToken(t *testing.T) {
	mockRepo := new(mocks.SecretRepository)
	mockConfig := &config.Token{
		Duration:          "1h",
		IsSaveSecretAtAws: false,
	}

	// Create the TokenService
	tokenService, err := service.NewTokenService(mockConfig, mockRepo)
	assert.NoError(t, err)

	// Test data
	user := &domain.User{
		ID:   1,
		Role: "normal",
	}

	// Call CreateToken
	token, err := tokenService.CreateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	containsV4 := strings.Contains(token, "v4")

	assert.True(t, containsV4, "The string does not contain 'v4'")

	mockRepo.AssertExpectations(t)
}

func TestTokenService_VerifyToken(t *testing.T) {
	mockRepo := new(mocks.SecretRepository)
	mockConfig := &config.Token{
		Duration:          "1h",
		IsSaveSecretAtAws: false,
	}

	// Create the TokenService
	tokenService, err := service.NewTokenService(mockConfig, mockRepo)
	assert.NoError(t, err)

	// Test data
	user := &domain.User{
		ID:   1,
		Role: "normal",
	}

	// Create a token
	token, err := tokenService.CreateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	// Verify the token
	payload, err := tokenService.VerifyToken(token)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, payload.UserID)
	assert.Equal(t, user.Role, payload.Role)

	mockRepo.AssertExpectations(t)
}

func TestTokenService_VerifyTokenExpired(t *testing.T) {
	mockRepo := new(mocks.SecretRepository)
	mockConfig := &config.Token{
		Duration:          "1s",
		IsSaveSecretAtAws: false,
	}

	// Create the TokenService
	tokenService, err := service.NewTokenService(mockConfig, mockRepo)
	assert.NoError(t, err)

	// Test data
	user := &domain.User{
		ID:   1,
		Role: "normal",
	}

	// Create a token
	token, err := tokenService.CreateToken(user)
	assert.NoError(t, err)

	// Wait for the token to expire
	time.Sleep(2 * time.Second)

	// Verify the token
	_, err = tokenService.VerifyToken(token)
	assert.ErrorIs(t, err, domain.ErrExpiredToken)

	mockRepo.AssertExpectations(t)
}

func TestTokenService_VerifyToken_Invalid(t *testing.T) {
	mockRepo := new(mocks.SecretRepository)
	mockConfig := &config.Token{
		Duration:          "1h",
		IsSaveSecretAtAws: false,
	}

	// Create the TokenService
	tokenService, err := service.NewTokenService(mockConfig, mockRepo)
	assert.NoError(t, err)

	// Use an invalid token
	invalidToken := "invalid.token"

	// Verify the token
	_, err = tokenService.VerifyToken(invalidToken)
	assert.ErrorIs(t, err, domain.ErrInvalidToken)

	mockRepo.AssertExpectations(t)
}
