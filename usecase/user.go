package usecase

import (
	"database/sql"

	"github.com/google/uuid"
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
	//本来ならemailのバリデーションをする

	//一意でランダムな文字列を生成する
	userID, err := uuid.NewRandom() //返り値はuuid型
	if err != nil {
		return err
	}

	//domainを介してinfrastructureで実装した関数を呼び出す。
	// Persistence（Repository）を呼出
	err = uu.userRepository.Insert(DB, userID.String(), name, email)
	if err != nil {
		return err
	}
	return nil
}
