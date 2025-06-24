
package user

import (
	"time"
)

// User represents the user entity in the domain
type User struct {
	ID        string
	Email     string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new user with validation
func NewUser(email, name string) (*User, error) {
	if email == "" {
		return nil, ErrInvalidEmail
	}
	if name == "" {
		return nil, ErrInvalidName
	}

	return &User{
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

// UpdateName updates the user's name
func (u *User) UpdateName(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	u.Name = name
	u.UpdatedAt = time.Now()
	return nil
}
