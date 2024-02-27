package usecase

import (
	"database/sql"

	"github.com/take-o20/layered-architecture-sample/domain"
	"github.com/take-o20/layered-architecture-sample/domain/repository"
)

type UserUseCase interface {
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
	Insert(DB *sql.DB, name, email string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

// to confirtm UserUseCase interface
var _ UserUseCase = userUseCase{}

func NewUserUseCase(ur repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: ur,
	}
}

func (uu userUseCase) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	user, err := uu.userRepository.GetByUserID(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Insert(DB *sql.DB, name, email string) error {
	err := uu.userRepository.Insert(DB, name, email)
	if err != nil {
		return err
	}
	return nil
}
