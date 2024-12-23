package service

import (
	"context"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/port"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/util"
)

/**
 * UserService implements port.UserService interface
 */
type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Register creates a new user
// !Password will save to db
// !パスワードがDBに保存される
func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return nil, domain.ErrInternal
	}

	user.Password = hashedPassword

	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, domain.ErrConflictingData
	}

	return user, nil
}

// GetUser gets a user by ID
func (us *UserService) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	var user *domain.User

	user, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// Get all users
func (us *UserService) ListUsers(ctx context.Context, skip, limit uint64) ([]domain.User, error) {
	var users []domain.User

	users, err := us.repo.ListUsers(ctx, skip, limit)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return users, nil
}

// UpdateUser updates a user's name, email, and password
func (us *UserService) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	existingUser, err := us.repo.GetUserByID(ctx, user.ID)
	if err != nil {
		return nil, domain.ErrInternal
	}

	emptyData := user.Name == "" &&
		user.Email == "" &&
		user.Password == "" &&
		user.Role == ""

	sameData := existingUser.Name == user.Name &&
		existingUser.Email == user.Email &&
		existingUser.Role == user.Role
		
	if emptyData || sameData {
		return nil, domain.ErrNoUpdatedData
	}

	var hashedPassword string

	if user.Password != "" {
		hashedPassword, err = util.HashPassword(user.Password)
		if err != nil {
			return nil, domain.ErrInternal
		}
	}

	user.Password = hashedPassword

	_, err = us.repo.UpdateUser(ctx, user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (us *UserService) DeleteUser(ctx context.Context, id uint64) error {
	_, err := us.repo.GetUserByID(ctx, id)
	if err != nil {
		return domain.ErrInternal
	}

	return us.repo.DeleteUser(ctx, id)
}
