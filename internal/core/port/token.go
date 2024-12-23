package port

import (
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
)

type TokenService interface {

	CreateToken(user *domain.User) (string, error)

	VerifyToken(token string) (*domain.TokenPayload, error)
}