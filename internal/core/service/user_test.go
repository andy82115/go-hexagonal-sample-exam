package service_test

import (
	"context"
	"testing"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_Register(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userSvc := service.NewUserService(mockRepo)

	ctx := context.Background()
	user := &domain.User{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Password: "plaintextpassword",
	}

	hashedPassword := "hashedpassword"

	// Mock repository behavior
	mockRepo.On("CreateUser", ctx, mock.MatchedBy(func(u *domain.User) bool {
		return u.Email == user.Email
	})).Return(&domain.User{
		ID:       1,
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}, nil)

	// Call the service
	createdUser, err := userSvc.Register(ctx, user)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, createdUser)
	assert.Equal(t, uint64(1), createdUser.ID)
	assert.Equal(t, hashedPassword, createdUser.Password)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userSvc := service.NewUserService(mockRepo)

	ctx := context.Background()
	expectedUser := &domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}

	// Mock repository behavior
	mockRepo.On("GetUserByID", ctx, uint64(1)).Return(expectedUser, nil)

	// Call the service
	user, err := userSvc.GetUser(ctx, 1)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser, user)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestUserService_ListUsers(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userSvc := service.NewUserService(mockRepo)

	ctx := context.Background()
	users := []domain.User{
		{ID: 1, Name: "John Doe", Email: "johndoe@example.com"},
		{ID: 2, Name: "Jane Doe", Email: "janedoe@example.com"},
	}

	// Mock repository behavior
	mockRepo.On("ListUsers", ctx, uint64(0), uint64(10)).Return(users, nil)

	// Call the service
	result, err := userSvc.ListUsers(ctx, 0, 10)

	// Assertions
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, users, result)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userSvc := service.NewUserService(mockRepo)

	ctx := context.Background()
	existingUser := &domain.User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Role:  "normal",
	}

	updatedUser := &domain.User{
		ID:    1,
		Name:  "John Doe Updated",
		Email: "johndoe@example.com",
		Role:  "normal",
	}

	// Mock repository behavior
	mockRepo.On("GetUserByID", ctx, uint64(1)).Return(existingUser, nil)
	mockRepo.On("UpdateUser", ctx, updatedUser).Return(updatedUser, nil)

	// Call the service
	result, err := userSvc.UpdateUser(ctx, updatedUser)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Name, result.Name)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	userSvc := service.NewUserService(mockRepo)

	ctx := context.Background()

	// Mock repository behavior
	mockRepo.On("GetUserByID", ctx, uint64(1)).Return(&domain.User{ID: 1}, nil)
	mockRepo.On("DeleteUser", ctx, uint64(1)).Return(nil)

	// Call the service
	err := userSvc.DeleteUser(ctx, 1)

	// Assertions
	assert.NoError(t, err)

	// Verify expectations
	mockRepo.AssertExpectations(t)
}
