package http

import (
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/go-playground/validator/v10"
)

// userRoleValidator is a custom validator for validating user roles
// userRoleValidator はユーザーのバリデータです。userRoleのパラメータが正しいをチェックする。
var userRoleValidator validator.Func = func(fl validator.FieldLevel) bool {
	userRole := fl.Field().Interface().(domain.UserRole)

	switch userRole {
	case "premium", "normal":
		return true
	default:
		return false
	}
}