package usecase

import (
	"database/sql"

	"github.com/take-o20/layered-architecture-sample/domain"
	"github.com/take-o20/layered-architecture-sample/domain/repository"
)

type UserUseCase interface {
	GetByUserID(DB *sql.DB, userID string) (*domain.User, error)
	Insert(DB *sql.DB, name, email string) (*domain.User, error)
	List(DB *sql.DB) ([]domain.User, error)
	Update(DB *sql.DB, userID, name, email string) (*domain.User, error)
	Delete(DB *sql.DB, userID string) (*domain.User, error)
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

func (uu userUseCase) Insert(DB *sql.DB, name, email string) (*domain.User, error) {
	user, err := uu.userRepository.Insert(DB, name, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) List(DB *sql.DB) ([]domain.User, error) {
	users, err := uu.userRepository.List(DB)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu userUseCase) Update(DB *sql.DB, userID, name, email string) (*domain.User, error) {
	user, err := uu.userRepository.Update(DB, userID, name, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUseCase) Delete(DB *sql.DB, userID string) (*domain.User, error) {
	user, err := uu.userRepository.Delete(DB, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
