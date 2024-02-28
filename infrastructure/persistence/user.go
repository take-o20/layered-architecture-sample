package persistence

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/take-o20/layered-architecture-sample/domain"
	"github.com/take-o20/layered-architecture-sample/domain/repository"
)

type userPersistence struct{}

func NewUserPersistence() repository.UserRepository {
	return &userPersistence{}
}

func (up userPersistence) Insert(DB *sql.DB, name, email string) (*domain.User, error) {
	stmt, err := DB.Prepare("INSERT INTO users(name, email) VALUES(?, ?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	user, err := getUserByUserId(DB, strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (up userPersistence) GetByUserID(DB *sql.DB, userID string) (*domain.User, error) {
	return getUserByUserId(DB, userID)
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

func (up userPersistence) Update(DB *sql.DB, userID, name, email string) (*domain.User, error) {
	stmt, err := DB.Prepare("UPDATE users SET name=?, email=? where user_id=?")
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(name, email, userID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	updatedUser, err := getUserByUserId(DB, userID)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (up userPersistence) Delete(DB *sql.DB, userID string) (*domain.User, error) {
	user, err := getUserByUserId(DB, userID)
	if err != nil {
		return nil, err
	}

	stmt, err := DB.Prepare("DELETE from users where user_id=?")
	if err != nil {
		return nil, err
	}

	result, err := stmt.Exec(userID)
	if err != nil {
		return nil, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("expected single row affected, got %d rows affected", rows)
	}

	return user, nil
}

func getUserByUserId(DB *sql.DB, userID string) (*domain.User, error) {
	row := DB.QueryRow("SELECT * FROM users WHERE user_id = ?", userID)
	return convertToUser(row)
}

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
