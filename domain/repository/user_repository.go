package repository

import (
	"database/sql"

	"github.com/take-o20/layered-architecture-sample/domain"
)

type UserRepository interface {
	Insert(DB *sql.DB, name, email string) (*domain.User, error)
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
	List(DB *sql.DB) ([]domain.User, error)
	Update(DB *sql.DB, userID, name, email string) (*domain.User, error)
	Delete(DB *sql.DB, userID string) (*domain.User, error)
}
