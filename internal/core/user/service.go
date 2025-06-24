
package user

// DomainService provides domain-specific business logic
type DomainService struct{}

// NewDomainService creates a new user domain service
func NewDomainService() *DomainService {
	return &DomainService{}
}

// ValidateUserCreation validates user creation business rules
func (s *DomainService) ValidateUserCreation(email, name string) error {
	if _, err := NewEmail(email); err != nil {
		return err
	}
	if name == "" {
		return ErrInvalidName
	}
	return nil
}
