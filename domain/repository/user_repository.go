package repository

import (
	"database/sql"

	"github.com/take-o20/layered-architecture-sample/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, userID, name, email string) error
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
}
