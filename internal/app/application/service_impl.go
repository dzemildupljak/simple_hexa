// **application:** Contains interfaces and implementations for application services.
package application

import (
	"fmt"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

// UserServiceImpl is an implementation of the UserService interface.
type UserServiceImpl struct {
	UserRepository domain.UserRepository
}

// NewUserService creates a new UserServiceImpl with the given UserRepository.
func NewUserService(userRepository domain.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(username, email string) error {
	newUser := &domain.User{
		Username: username,
		Email:    email,
	}

	err := s.UserRepository.SaveUser(newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *UserServiceImpl) GetUserByID(userID int) (*domain.User, error) {
	return s.UserRepository.GetUserByID(userID)
}

func (s *UserServiceImpl) GetUserByEmail(email string) (*domain.User, error) {
	return s.UserRepository.GetUserByEmail(email)
}

func (s *UserServiceImpl) GetAllUsers() ([]*domain.User, error) {
	return s.UserRepository.GetAllUsers()
}
