// **application:** Contains interfaces and implementations for application services.
package application

import (
	"context"
	"fmt"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	"github.com/dzemildupljak/simple_hexa/internal/app/ports/outbound"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UserServiceImpl is an implementation of the UserService interface.
type UserServiceImpl struct {
	UserRepository outbound.UserRepository
}

// NewUserService creates a new UserServiceImpl with the given UserRepository.
func NewUserService(userRepository outbound.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{
		UserRepository: userRepository,
	}
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, newUser *domain.User) error {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserService-GetAllUsers").End()
	}
	err := s.UserRepository.SaveUser(ctx, newUser)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (s *UserServiceImpl) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserService-GetUserById").End()
	}

	user, err := s.UserRepository.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *UserServiceImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserService-GetUserByEmail").End()
	}
	user, err := s.UserRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *UserServiceImpl) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserService-GetAllUsers").End()
	}

	users, err := s.UserRepository.GetAllUsers(ctx)

	if err != nil {
		return nil, err
	}

	return users, err
}
