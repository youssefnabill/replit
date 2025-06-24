
package repository

import (
	"context"
	"hexagonal-backend/internal/core/user"
)

// UserRepository defines the interface for user data persistence
type UserRepository interface {
	Create(ctx context.Context, user *user.User) error
	GetByID(ctx context.Context, id string) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
	Update(ctx context.Context, user *user.User) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*user.User, error)
}
