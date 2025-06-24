
package repository

import (
	"context"
	"errors"
	"sync"

	"hexagonal-backend/internal/core/user"
	"hexagonal-backend/internal/port/repository"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

// MemoryUserRepository implements UserRepository using in-memory storage
type MemoryUserRepository struct {
	users map[string]*user.User
	mutex sync.RWMutex
}

// NewMemoryUserRepository creates a new in-memory user repository
func NewMemoryUserRepository() repository.UserRepository {
	return &MemoryUserRepository{
		users: make(map[string]*user.User),
	}
}

// Create stores a new user
func (r *MemoryUserRepository) Create(ctx context.Context, u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[u.ID]; exists {
		return ErrUserAlreadyExists
	}

	r.users[u.ID] = u
	return nil
}

// GetByID retrieves a user by ID
func (r *MemoryUserRepository) GetByID(ctx context.Context, id string) (*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	u, exists := r.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}

	return u, nil
}

// GetByEmail retrieves a user by email
func (r *MemoryUserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, u := range r.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, ErrUserNotFound
}

// Update updates an existing user
func (r *MemoryUserRepository) Update(ctx context.Context, u *user.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[u.ID]; !exists {
		return ErrUserNotFound
	}

	r.users[u.ID] = u
	return nil
}

// Delete removes a user by ID
func (r *MemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.users[id]; !exists {
		return ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}

// List retrieves a list of users with pagination
func (r *MemoryUserRepository) List(ctx context.Context, limit, offset int) ([]*user.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	users := make([]*user.User, 0, len(r.users))
	for _, u := range r.users {
		users = append(users, u)
	}

	start := offset
	if start > len(users) {
		return []*user.User{}, nil
	}

	end := start + limit
	if end > len(users) {
		end = len(users)
	}

	return users[start:end], nil
}
