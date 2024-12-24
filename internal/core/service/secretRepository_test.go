package service_test

import (
	"testing"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
	"github.com/stretchr/testify/assert"
)

func TestSecretRepositoryWithRealServer(t *testing.T) {
	// real server URL
	realServerURL := "http://enqguo42bd.execute-api.localhost.localstack.cloud:4566"

	// 初始化 SecretRepository
	repo, err := service.NewSecretRepository(&config.AWS{AwasUrl: realServerURL})

	if err != nil {
		t.Fatalf("Failed to initialize SecretRepository: %v", err)
	}

	t.Run("GetTokenSecret - Real Server", func(t *testing.T) {
		t.Skip("This test is for testing real server, skip it for now.")
		param := domain.SecretGetParam{
			Action: "get",
		}

		response, err := repo.GetTokenSecret(param)

		assert.NoError(t, err, "Failed to call GetTokenSecret")
		assert.NotEmpty(t, response.Secret, "Secret should not be empty")
	})

	t.Run("UpdateTokenSecret - Real Server", func(t *testing.T) {
		t.Skip("This test is for testing real server, skip it for now.")
		param := domain.SecretUpdateParam{
			Action: "update",
			Secret: domain.Secret{
				Password: "new-password",
			},
		}

		response, err := repo.UpdateTokenSecret(param)

		assert.NoError(t, err, "Failed to call UpdateTokenSecret")
		assert.Equal(t, "Secret updated successfully", response.Message)
	})
}
