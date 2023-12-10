// **application:** Contains interfaces and implementations for application services.
package application

import "github.com/dzemildupljak/simple_hexa/internal/app/domain"

// UserService defines the interface for user-related application services.
type UserService interface {
	CreateUser(*domain.User) error
	GetUserByID(userID int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	GetAllUsers() ([]*domain.User, error)
}
