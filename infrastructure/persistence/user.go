package persistence

import (
	"database/sql"
	"fmt"

	"github.com/take-o20/layered-architecture-sample/domain"
	"github.com/take-o20/layered-architecture-sample/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) Insert(DB *sql.DB, name, email string) error {
	stmt, err := DB.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(name, email)
	return err
}

func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM users WHERE user_id = ?", userID)
	//row型をgolangで利用できる形にキャストする。
	return convertToUser(row)
}

func (up userPersistence) List(DB *sql.DB) ([]domain.User, error) {
	rows, err := DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := convertToUserList(rows)
	if err != nil {
		return nil, fmt.Errorf("failed to List users: %w", err)
	}
	return users, nil
}
func (up userPersistence) Update(DB *sql.DB, userID, name, email string) (*domain.User, error)
func (up userPersistence) Delete(DB *sql.DB, userID string) (*domain.User, error)

func convertToUserList(rows *sql.Rows) ([]domain.User, error) {
	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.UserID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("failed to List users: %w", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func convertToUser(row *sql.Row) (*domain.User, error) {
	user := domain.User{}
	err := row.Scan(&user.UserID, &user.Name, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
