
package service

import "context"

// NotificationService defines the interface for external notification services
type NotificationService interface {
	SendWelcomeEmail(ctx context.Context, email, name string) error
	SendNotification(ctx context.Context, userID, message string) error
}
