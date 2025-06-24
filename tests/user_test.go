
package tests

import (
	"context"
	"testing"

	"hexagonal-backend/internal/adapter/repository"
	"hexagonal-backend/internal/adapter/service"
	"hexagonal-backend/internal/app"
	"hexagonal-backend/internal/core/user"
)

func TestUserApp_CreateUser(t *testing.T) {
	// Setup
	userRepo := repository.NewMemoryUserRepository()
	notificationService := service.NewEmailNotificationService()
	userApp := app.NewUserApp(userRepo, notificationService)

	// Test
	ctx := context.Background()
	email := "test@example.com"
	name := "Test User"

	createdUser, err := userApp.CreateUser(ctx, email, name)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if createdUser.Email != email {
		t.Errorf("Expected email %s, got %s", email, createdUser.Email)
	}

	if createdUser.Name != name {
		t.Errorf("Expected name %s, got %s", name, createdUser.Name)
	}

	// Verify user was saved
	savedUser, err := userApp.GetUser(ctx, createdUser.ID)
	if err != nil {
		t.Fatalf("Expected no error retrieving user, got %v", err)
	}

	if savedUser.ID != createdUser.ID {
		t.Errorf("Expected user ID %s, got %s", createdUser.ID, savedUser.ID)
	}
}

func TestUserEntity_UpdateName(t *testing.T) {
	// Setup
	u, err := user.NewUser("test@example.com", "Test User")
	if err != nil {
		t.Fatalf("Expected no error creating user, got %v", err)
	}

	// Test
	newName := "Updated Name"
	err = u.UpdateName(newName)

	// Assertions
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if u.Name != newName {
		t.Errorf("Expected name %s, got %s", newName, u.Name)
	}
}

func TestUserEntity_InvalidEmail(t *testing.T) {
	// Test
	_, err := user.NewUser("invalid-email", "Test User")

	// Assertions
	if err == nil {
		t.Fatal("Expected error for invalid email, got nil")
	}

	if err != user.ErrInvalidEmail {
		t.Errorf("Expected ErrInvalidEmail, got %v", err)
	}
}
