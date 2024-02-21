package postgres

import (
	"fmt"

	"database/sql"
	_ "github.com/lib/pq"
)

var PostgresConnectionConfig DatabaseConnectionConfig

type DatabaseConnectionConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Dbname   string
	SslMode  string
}

// DatabaseConnectionString builds and returns the database connection string
func DatabaseConnectionString(dbConnConfig DatabaseConnectionConfig) string {

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConnConfig.Host,
		dbConnConfig.Port,
		dbConnConfig.User,
		dbConnConfig.Password,
		dbConnConfig.Dbname,
		dbConnConfig.SslMode)
}

func NewDatabaseConnection(connectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
