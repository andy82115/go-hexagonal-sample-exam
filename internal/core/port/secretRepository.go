package port

import "github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"

type SecretRepository interface {
	GetTokenSecret(param domain.SecretGetParam) (domain.SecretGetResponse, error)

	UpdateTokenSecret(param domain.SecretUpdateParam) (domain.SecretUpdateResponse, error)
}
