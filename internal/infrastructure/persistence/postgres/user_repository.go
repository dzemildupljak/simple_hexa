// **infrastructure:** Implements infrastructure details, such as database access or external service connections
package persistence

import (
	"context"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	"github.com/dzemildupljak/simple_hexa/internal/app/ports/outbound"
	"github.com/newrelic/go-agent/v3/newrelic"
)

// UserRepositoryImpl is an implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	users map[string]*domain.User
}

func NewUserRepository() *UserRepositoryImpl {
	users := make(map[string]*domain.User)
	users["999"] = &domain.User{
		ID:       999,
		Username: "defaultUser",
		Email:    "defaultuser@mail.com",
	}
	return &UserRepositoryImpl{
		users: users,
	}
}

func (r *UserRepositoryImpl) SaveUser(ctx context.Context, user *domain.User) error {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserRepository-SaveUser").End()
	}
	if r.users[user.Email] != nil {
		return outbound.ErrInvalidOperation
	}
	genUid := len(r.users) + 1
	user.ID = genUid

	r.users[user.Email] = user

	return nil
}

func (r *UserRepositoryImpl) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserRepository-GetUserById").End()
	}

	for _, u := range r.users {
		if u.ID == userId {
			return u, nil
		}
	}

	return nil, outbound.ErrUserNotFound
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserRepository-GetUserByEmail").End()
	}
	user, found := r.users[email]
	if !found {
		return nil, outbound.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	if txn := newrelic.FromContext(ctx); txn != nil {
		defer txn.StartSegment("UserRepository-GetUserByEmail").End()
	}

	var users []*domain.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
