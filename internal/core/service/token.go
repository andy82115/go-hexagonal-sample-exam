package service

import (
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/port"
	"github.com/google/uuid"
)

/**
 * TokenService implements port.TokenService interface
 * Provide an access to the paseto library -> go-paseto
 * TokenService は port.TokenService インターフェースを実装する。
 * paseto ライブラリへのアクセスを提供する -> go-paseto
 */
type TokenService struct {
	token    *paseto.Token
	key      *paseto.V4SymmetricKey
	parser   *paseto.Parser
	duration time.Duration
}

// NewTokenService creates a new TokenService instance
func NewTokenService(config *config.Token, repository port.SecretRepository) (port.TokenService, error) {
	duration, err := time.ParseDuration(config.Duration)
	if err != nil {
		return nil, domain.ErrTokenDuration
	}

	key, err := getKey(repository, config.IsSaveSecretAtAws)
	if err != nil {
		return nil, err
	}

	token := paseto.NewToken()
	parser := paseto.NewParser()

	return &TokenService{
		token:    &token,
		key:      key,
		parser:   &parser,
		duration: duration,
	}, nil
}

// getKey retrieves an existing key or generates a new one if not valid
func getKey(repository port.SecretRepository, isSaveToAws bool) (*paseto.V4SymmetricKey, error) {
	// Attempt to fetch the secret if AWS saving is enabled
	// AWS保存が有効な場合、秘密の取得を試みる。
	secretToken := ""
	if isSaveToAws {
		resp, err := repository.GetTokenSecret(domain.SecretGetParam{})
		if err != nil {
			return nil, domain.ErrTokenDuration
		}
		secretToken = resp.Secret.Password
	}

	// Validate or Generate the secret token
	key, err := paseto.V4SymmetricKeyFromBytes([]byte(secretToken))
	if err != nil {
		key = paseto.NewV4SymmetricKey()
		if isSaveToAws {
			_, err = repository.UpdateTokenSecret(domain.SecretUpdateParam{
				Secret: domain.Secret{Password: string(key.ExportBytes())},
			})
			if err != nil {
				return nil, err
			}
		}
	}

	return &key, nil
}

// CreateToken creates a new paseto token
func (pt *TokenService) CreateToken(user *domain.User) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	payload := &domain.TokenPayload{
		ID:     id,
		UserID: user.ID,
		Role:   user.Role,
	}

	err = pt.token.Set("payload", payload)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(pt.duration)

	pt.token.SetIssuedAt(issuedAt)
	pt.token.SetNotBefore(issuedAt)
	pt.token.SetExpiration(expiredAt)

	token := pt.token.V4Encrypt(*pt.key, nil)

	return token, nil
}

// VerifyToken verifies the paseto token
func (pt *TokenService) VerifyToken(token string) (*domain.TokenPayload, error) {
	var payload *domain.TokenPayload

	parsedToken, err := pt.parser.ParseV4Local(*pt.key, token, nil)
	if err != nil {
		if err.Error() == "this token has expired" {
			return nil, domain.ErrExpiredToken
		}
		return nil, domain.ErrInvalidToken
	}

	err = parsedToken.Get("payload", &payload)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}

	return payload, nil
}
