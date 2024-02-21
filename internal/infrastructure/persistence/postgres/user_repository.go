// Package postgres **infrastructure:** Implements infrastructure details,
// such as database access or external service connections
package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/dzemildupljak/simple_hexa/config"
	"github.com/dzemildupljak/simple_hexa/internal/app/domain"
	"github.com/dzemildupljak/simple_hexa/internal/app/ports/outbound"
	"log"
)

var publicUsers = "public.users"

// UserRepositoryImpl is an implementation of the UserRepository interface.
type UserRepositoryImpl struct {
	//users map[string]*domain.User
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {

	return &UserRepositoryImpl{
		db: db,
	}
}

func (r *UserRepositoryImpl) SaveUser(ctx context.Context, user *domain.User) error {
	if endSegment, segErr := config.NewRelicSegment(ctx); segErr == nil {
		defer endSegment()
	}
	// Define the SQL query to insert a new user
	query := `INSERT INTO public.users (username, email) VALUES ($1, $2) RETURNING id`
	if endDtSegment, segErr := config.StartDatastoreNewRelicSegment(
		ctx, query, publicUsers, config.INSERT); segErr == nil {
		defer endDtSegment()
	}
	// Execute the query
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return outbound.ErrInvalidOperation
		}
		log.Printf("Error saving user: %v", err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetUserById(ctx context.Context, userId int) (*domain.User, error) {
	if endSegment, segErr := config.NewRelicSegment(ctx); segErr == nil {
		defer endSegment()
	}
	// Define the SQL query
	query := `SELECT id, username, email FROM public.users WHERE id = $1`
	if endDtSegment, segErr := config.StartDatastoreNewRelicSegment(
		ctx, query, publicUsers, config.SELECT); segErr == nil {
		defer endDtSegment()
	}

	// Execute the query
	var user domain.User
	err := r.db.QueryRowContext(ctx, query, userId).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found
			return nil, outbound.ErrUserNotFound
		}
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	if endSegment, segErr := config.NewRelicSegment(ctx); segErr == nil {
		defer endSegment()
	}

	// Define the SQL query
	query := `SELECT id, username, email FROM public.users WHERE email = $1`

	if endDtSegment, segErr := config.StartDatastoreNewRelicSegment(
		ctx, query, publicUsers, config.SELECT); segErr == nil {
		defer endDtSegment()
	}

	// Execute the query
	var user domain.User
	err := r.db.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// User not found
			return nil, outbound.ErrUserNotFound
		}
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	if endSegment, segErr := config.NewRelicSegment(ctx); segErr == nil {
		defer endSegment()
	}

	// Define the paginated SQL query
	query := `SELECT id, username, email FROM public.users LIMIT $1 OFFSET $2`

	if endDtSegment, segErr := config.StartDatastoreNewRelicSegment(
		ctx, query, publicUsers, config.SELECT); segErr == nil {
		defer endDtSegment()
	}

	// Eventually add it like parameters
	pageNumber, pageSize := 1, 10
	// Calculate Offset
	offset := (pageNumber - 1) * pageSize

	// Execute the query
	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		log.Printf("Error fetching users: %v", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error closing rows: %v", err)
		}
	}(rows)

	var users []*domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Printf("Error scanning user: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating users: %v", err)
		return nil, err
	}

	return users, nil
}
