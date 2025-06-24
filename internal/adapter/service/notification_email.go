
package service

import (
	"context"
	"fmt"
	"log"

	"hexagonal-backend/internal/port/service"
)

// EmailNotificationService implements NotificationService using email
type EmailNotificationService struct{}

// NewEmailNotificationService creates a new email notification service
func NewEmailNotificationService() service.NotificationService {
	return &EmailNotificationService{}
}

// SendWelcomeEmail sends a welcome email to new users
func (s *EmailNotificationService) SendWelcomeEmail(ctx context.Context, email, name string) error {
	// In a real implementation, this would integrate with an email service
	log.Printf("Sending welcome email to %s (%s)", email, name)
	return nil
}

// SendNotification sends a notification to a user
func (s *EmailNotificationService) SendNotification(ctx context.Context, userID, message string) error {
	// In a real implementation, this would integrate with an email service
	log.Printf("Sending notification to user %s: %s", userID, message)
	return nil
}
