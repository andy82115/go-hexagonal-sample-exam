package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/andy82115/go-hexagonal-sample-exam/internal/adapter/config"
	"github.com/andy82115/go-hexagonal-sample-exam/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `gorm:"primaryKey"`         
	Name      string    `gorm:"size:255;not null"` 
	Email     string    `gorm:"size:255;unique;not null"`
	Password  string    `gorm:"size:255;not null"` 
	Role      domain.UserRole  `gorm:"type:user_role;default:'normal'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 * UserRepository は port.UserRepository インターフェースを実装しています。
 * postgres データベースへのアクセスを提供します。
 */
type UserRepository struct {
	db *gorm.DB
}

// The gorm will abstract all the db action. Make this service easy to change db endpoint.
// gormはすべてのdbアクションを抽象化する。dbのエンドポイントを簡単に変更できるようにする。
func NewCustomWayRespository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func NewUserRepository(config *config.DB) (*UserRepository, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Host, config.User, config.Password, config.Name, config.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if (err != nil) {
		return nil, err
	}

	// create enum before auto-migration
	// auto-migrationの前に列挙型を作成する
	db.Exec(`CREATE TYPE user_role AS ENUM ('premium', 'normal')`)
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, domain.ErrConflictingData
	}

	return &UserRepository{
		db,
	}, nil
}

// Database entity trans to local entity
func dbUserToDomainUser(dbUser *User) *domain.User {
	return &domain.User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Role:      domain.UserRole(dbUser.Role), // Convert from string to UserRole
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	} 
}

// Local entity trans to Database entity 
func domainUserToDbUser(user *domain.User) *User {
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      domain.UserRole(user.Role), // Convert from string to UserRole
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	} 
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	dbUser := domainUserToDbUser(user)
	result := ur.db.WithContext(ctx).Create(&dbUser)

	if result.Error != nil {
		return nil, domain.ErrConflictingData
	}

	return dbUserToDomainUser(dbUser), nil
}

func (ur *UserRepository) GetUserByID(ctx context.Context, id uint64) (*domain.User, error) {
	var dbUser User

	// Query the database for the user with the given ID
	if err := ur.db.WithContext(ctx).First(&dbUser, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrConflictingData
	}

	return dbUserToDomainUser(&dbUser), nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	var dbUser User

	// Query the database to find the user by email
	if err := ur.db.WithContext(ctx).Where("email = ?", email).First(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrConflictingData
	}

	return dbUserToDomainUser(&dbUser), nil
}

// ListUsers lists all users from the database, paginated by ID
func (ur *UserRepository) ListUsers(ctx context.Context, lastID uint64, limit uint64) ([]domain.User, error) {
	var dbUsers []User
	var users []domain.User

	if err := ur.db.WithContext(ctx).
		Where("id > ?", lastID).
		Limit(int(limit)).
		Order("id ASC").
		Find(&dbUsers).Error; err != nil {
		return nil, domain.ErrDataNotFound
	}

	for _, dbUser := range dbUsers {
		users = append(users, *dbUserToDomainUser(&dbUser))
	}

	return users, nil
}

func (ur *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	var dbUser User

	if err := ur.db.WithContext(ctx).First(&dbUser, user.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrDataNotFound
		}
		return nil, domain.ErrConflictingData
	}

	if user.Name != "" {
		dbUser.Name = user.Name
	}
	if user.Email != "" {
		dbUser.Email = user.Email
	}
	if user.Password != "" {
		dbUser.Password = user.Password
	}
	if user.Role != "" {
		dbUser.Role = user.Role
	}

	dbUser.UpdatedAt = time.Now()

	if err := ur.db.WithContext(ctx).Save(&dbUser).Error; err != nil {
		return nil, err
	}

	return dbUserToDomainUser(&dbUser), nil
}


func (ur *UserRepository) DeleteUser(ctx context.Context, id uint64) error {
	var dbUser User

	if err := ur.db.WithContext(ctx).First(&dbUser, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.ErrDataNotFound
		}
		return domain.ErrConflictingData
	}

	if err := ur.db.WithContext(ctx).Delete(&dbUser).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) CloseDB() error{
	sqlDB, err := ur.db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	// fmt.Println("Database connection closed successfully.")

	return nil
}
