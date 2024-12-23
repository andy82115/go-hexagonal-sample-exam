package http

import (
	"errors"
	"net/http"
	"time"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// response represents a response body format
// response 基本的返回格式
type response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

func newResponse(success bool, message string, data any) response {
	return response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// meta represents metadata for pager design
// metaはページャーデザインのメタデータを表す。
type meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

func newMeta(total, limit, skip uint64) meta {
	return meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

type authResponse struct {
	AccessToken string `json:"token" example:"v2.local.Gdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

func newAuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

type userResponse struct {
	ID        uint64    `json:"id" example:"1"`
	Name      string    `json:"name" example:"John Doe"`
	Email     string    `json:"email" example:"test@example.com"`
	CreatedAt time.Time `json:"created_at" example:"1970-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" example:"1970-01-01T00:00:00Z"`
}

func newUserResponse(user *domain.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

var errorStatusMap = map[error]int{
	domain.ErrInternal:                   http.StatusInternalServerError,
	domain.ErrDataNotFound:               http.StatusNotFound,
	domain.ErrConflictingData:            http.StatusConflict,
	domain.ErrInvalidCredentials:         http.StatusUnauthorized,
	domain.ErrUnauthorized:               http.StatusUnauthorized,
	domain.ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	domain.ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	domain.ErrInvalidToken:               http.StatusUnauthorized,
	domain.ErrExpiredToken:               http.StatusUnauthorized,
	domain.ErrForbidden:                  http.StatusForbidden,
	domain.ErrNoUpdatedData:              http.StatusBadRequest,
}

// validationError for error response when validation error happened
// validation error発生した用
func validationError(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

// determines the status code of an error and returns a JSON response
// エラーのステータスコードを決定し、JSONレスポンスを返す。
func handleError(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}

// aborts request when special status code is given
// 特定のステータスコードが与えられたらリクエストを中断する。
func handleAbort(ctx *gin.Context, err error) {
	statusCode, ok := errorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.AbortWithStatusJSON(statusCode, errRsp)
}

// parses error messages and returns a slice of error messages
// エラーメッセージを解析し、エラーメッセージのスライスを返す
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

type errorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

func newErrorResponse(errMsgs []string) errorResponse {
	return errorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

func handleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(true, "Success", data)
	ctx.JSON(http.StatusOK, rsp)
}
