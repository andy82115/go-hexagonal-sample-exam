package domain

import (
	"errors"
)

var (
	// ErrInternal is an error for internal service fails to process the request
	// ErrInternalは、internalサービスがリクエストの処理に失敗した場合のエラーです。
	ErrInternal = errors.New("internal error")
	// ErrDataNotFound is an error for requested data is not found at both cache and database
	// ErrDataNotFoundは、要求されたデータがキャッシュとデータベースの両方で見つからない場合のエラーです。
	ErrDataNotFound = errors.New("data not found")
	// ErrNoUpdatedData is an error when no data is provided for updating
	// ErrNoUpdatedDataは、更新のためのデータが提供されていない場合のエラーである。
	ErrNoUpdatedData = errors.New("no data to update")
	// ErrConflictingData is an error when data conflicts with existing data
	// ErrConflictingDataは、データが既存のデータと衝突した場合のエラー。
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	// ErrTokenDuration is an error when the token duration format is invalid
	// ErrTokenDurationは、トークンの継続時間形式が無効な場合のエラーです。
	ErrTokenDuration = errors.New("invalid token duration format")
	// ErrTokenCreation is an error when the token creation fails
	// ErrTokenCreationは、トークンの作成に失敗したときのエラーです。
	ErrTokenCreation = errors.New("error creating token")
	// ErrExpiredToken is an error when the access token is expired
	// ErrExpiredTokenは、アクセストークンが期限切れの場合のエラーです。
	ErrExpiredToken = errors.New("access token has expired")
	// ErrInvalidToken is an error when the access token is invalid
	// ErrInvalidTokenは、アクセストークンが無効な場合のエラーです。
	ErrInvalidToken = errors.New("access token is invalid")
	// ErrInvalidCredentials is an error when the credentials are invalid
	// ErrInvalidCredentials は、クレデンシャルが無効な場合のエラーです。
	ErrInvalidCredentials = errors.New("invalid email or password")
	// ErrEmptyAuthorizationHeader is an error when the authorization header is empty
	// ErrEmptyAuthorizationHeaderは、認可ヘッダーが空の場合のエラーです。
	ErrEmptyAuthorizationHeader = errors.New("authorization header is not provided")
	// ErrInvalidAuthorizationHeader is an error when the authorization header is invalid
	// ErrInvalidAuthorizationHeaderは、認証ヘッダーが無効な場合のエラーです。
	ErrInvalidAuthorizationHeader = errors.New("authorization header format is invalid")
	// ErrInvalidAuthorizationType is an error when the authorization type is invalid
	// ErrInvalidAuthorizationTypeは、認証タイプが無効な場合のエラーである。
	ErrInvalidAuthorizationType = errors.New("authorization type is not supported")
	// ErrUnauthorized is an error when the user is unauthorized
	// ErrUnauthorizedは、ユーザーが認証されていない場合のエラーです。
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")
	// ErrForbidden is an error when the user is forbidden to access the resource
	// ErrForbiddenは、ユーザーがリソースにアクセスすることを禁じられている場合のエラーです。
	ErrForbidden = errors.New("user is forbidden to access the resource")
)
