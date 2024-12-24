package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	db.AutoMigrate(&service.User{})

	return db
}

// This test is make sure we can change db anytime we want. Ex: postgresql -> mysql...
// The only thing we have to make sure is the we need to init any table or enum first
func TestUserRepository(t *testing.T) {
	db := setupTestDB()
	repo := service.NewCustomWayRespository(db)

	ctx := context.Background()

	t.Run("CreateUser", func(t *testing.T) {
		user := &domain.User{
			Name:     "Test User",
			Email:    "test@example.com",
			Password: "hashed-password",
			Role:     domain.Normal,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createdUser, err := repo.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.Equal(t, user.Name, createdUser.Name)
		assert.Equal(t, user.Email, createdUser.Email)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		// Create a user to search
		user := &domain.User{
			Name:     "Fetch User",
			Email:    "fetch@example.com",
			Password: "hashed-password",
			Role:     domain.Normal,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		createdUser, _ := repo.CreateUser(ctx, user)

		fetchedUser, err := repo.GetUserByID(ctx, createdUser.ID)
		assert.NoError(t, err)
		assert.Equal(t, createdUser.ID, fetchedUser.ID)
	})

	t.Run("GetUserByEmail", func(t *testing.T) {
		email := "fetch@example.com"
		fetchedUser, err := repo.GetUserByEmail(ctx, email)
		assert.NoError(t, err)
		assert.Equal(t, email, fetchedUser.Email)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		// Create a user to update
		user := &domain.User{
			Name:     "Update User",
			Email:    "update@example.com",
			Password: "hashed-password",
			Role:     domain.Normal,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		createdUser, _ := repo.CreateUser(ctx, user)

		// Update the user's details
		createdUser.Name = "Updated Name"
		updatedUser, err := repo.UpdateUser(ctx, createdUser)
		assert.NoError(t, err)
		assert.Equal(t, "Updated Name", updatedUser.Name)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		// Create a user to delete
		user := &domain.User{
			Name:     "Delete User",
			Email:    "delete@example.com",
			Password: "hashed-password",
			Role:     domain.Normal,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		createdUser, _ := repo.CreateUser(ctx, user)

		// Delete the user
		err := repo.DeleteUser(ctx, createdUser.ID)
		assert.NoError(t, err)

		// Verify the user no longer exists
		_, err = repo.GetUserByID(ctx, createdUser.ID)
		assert.Error(t, err)
	})
}
