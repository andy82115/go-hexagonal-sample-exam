package service

import (
	"context"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/port"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/util"
)

/**
 * AuthService implements port.AuthService interface
 * and provides an access to the user repository
 * and token service
 */
type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

// NewAuthService creates a new auth service instance
func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{
		repo,
		ts,
	}
}

// Login gives a registered user an access token if the credentials are valid
func (as *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return "", domain.ErrDataNotFound
		}
		return "", domain.ErrInternal
	}

	// Compare password
	if err := util.ComparePassword(password, user.Password); err != nil {
		return "", domain.ErrInvalidCredentials
	}

	// Create token
	token, err := as.ts.CreateToken(user)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return token, nil
}
