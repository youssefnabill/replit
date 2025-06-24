
package app

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	"hexagonal-backend/internal/core/user"
	"hexagonal-backend/internal/port/repository"
	"hexagonal-backend/internal/port/service"
)

// UserApp represents the user application service
type UserApp struct {
	userRepo            repository.UserRepository
	notificationService service.NotificationService
	domainService       *user.DomainService
}

// NewUserApp creates a new user application service
func NewUserApp(userRepo repository.UserRepository, notificationService service.NotificationService) *UserApp {
	return &UserApp{
		userRepo:            userRepo,
		notificationService: notificationService,
		domainService:       user.NewDomainService(),
	}
}

// CreateUser creates a new user
func (app *UserApp) CreateUser(ctx context.Context, email, name string) (*user.User, error) {
	// Validate using domain service
	if err := app.domainService.ValidateUserCreation(email, name); err != nil {
		return nil, err
	}

	// Create user entity
	u, err := user.NewUser(email, name)
	if err != nil {
		return nil, err
	}

	// Generate ID
	u.ID = generateID()

	// Save to repository
	if err := app.userRepo.Create(ctx, u); err != nil {
		return nil, err
	}

	// Send welcome notification
	if err := app.notificationService.SendWelcomeEmail(ctx, u.Email, u.Name); err != nil {
		// Log error but don't fail the operation
		// In a real implementation, you might want to use a message queue for reliability
	}

	return u, nil
}

// GetUser retrieves a user by ID
func (app *UserApp) GetUser(ctx context.Context, id string) (*user.User, error) {
	return app.userRepo.GetByID(ctx, id)
}

// ListUsers retrieves a list of users
func (app *UserApp) ListUsers(ctx context.Context, limit, offset int) ([]*user.User, error) {
	return app.userRepo.List(ctx, limit, offset)
}

// generateID generates a random ID
func generateID() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
