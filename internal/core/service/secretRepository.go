package service

import (
	"log/slog"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/port"
	"github.com/go-resty/resty/v2"
)

type SecretRepository struct{
	baseUri string
	client  *resty.Client
}

func NewSecretRepository(config *config.AWS) (port.SecretRepository, error) {
	client := resty.New().SetHeader("Content-Type", "application/json")


	// Registering Request Middleware
	client.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		
    	slog.Info("Outgoing Request",
			"Method", req.Method,
			"URL", req.URL,
			"Headers", req.Header)

    	return nil  // if its success otherwise return error
  	})

	// Registering Response Middleware
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
    	slog.Info("Incoming Response",
			"Status Code", resp.StatusCode(),
			"Body", string(resp.Body()),
			"Headers", resp.Header())

    	return nil  // if its success otherwise return error
  	})


	return &SecretRepository{
		config.AwasUrl,
		client,
	}, nil
}

func (s *SecretRepository) GetTokenSecret(param domain.SecretGetParam) (domain.SecretGetResponse, error) {
	var secretResponse domain.SecretGetResponse
	param.Action = "get"

	url := s.baseUri + "/dev/secret"
	_, err := s.client.R().
		SetBody(param).
		SetResult(&secretResponse).
		Post(url)

	if err != nil {
		return domain.SecretGetResponse{ Message: "fail to get secret" }, domain.ErrInternal
	}

	return secretResponse, nil
}

func (s *SecretRepository) UpdateTokenSecret(param domain.SecretUpdateParam) (domain.SecretUpdateResponse, error) {
	var secretResponse domain.SecretUpdateResponse
	param.Action = "update"
	url := s.baseUri + "/dev/secret"
	_, err := s.client.R().
		SetBody(param).
		SetResult(&secretResponse).
		Post(url)

	if err != nil {
		return domain.SecretUpdateResponse{ Message: "fail to get secret" }, domain.ErrInternal
	}

	return secretResponse, nil
}



