// **infrastructure:** Implements infrastructure details, such as database access or external service connections
package persistence

import (
	"encoding/json"
	"fmt"

	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
)

// UserRepositoryImpl is an implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	users map[string]*domain.User
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{
		users: make(map[string]*domain.User),
	}
}

func (r *UserRepositoryImpl) SaveUser(user *domain.User) error {
	usersJSON, err := json.Marshal(r.users)
	if err != nil {
		fmt.Println("ERROR MRHASHDA DATA")
	}
	fmt.Println(string(usersJSON))

	if r.users[user.Email] != nil {
		return domain.ErrInvalidOperation
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		fmt.Println("ERROR MRHASHDA DATA")
	}
	fmt.Println(string(userJSON))

	genUid := len(r.users) + 1
	user.ID = genUid

	r.users[user.Email] = user

	return nil
}

func (r *UserRepositoryImpl) GetUserByID(userID int) (*domain.User, error) {
	for _, u := range r.users {
		if u.ID == userID {
			return u, nil
		}
	}

	return nil, domain.ErrUserNotFound
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*domain.User, error) {
	user, found := r.users[email]
	fmt.Println(r.users)
	fmt.Println(user)

	usersJSON, err := json.Marshal(r.users)
	if err != nil {
		fmt.Println("ERROR MRHASHDA DATA")
	}
	fmt.Println(string(usersJSON))

	if !found {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() ([]*domain.User, error) {
	var users []*domain.User
	for _, user := range r.users {
		users = append(users, user)
	}
	return users, nil
}
