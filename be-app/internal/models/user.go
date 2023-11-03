package models

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	RoleID    int       `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type UserModel interface {
	GetUsers() ([]User, error)
	CreateUser(user User) error
}

type UserModelImpl struct {
	Database *sql.DB
}

func NewUserModel(db *sql.DB) *UserModelImpl {
	return &UserModelImpl{Database: db}
}

func (uModel *UserModelImpl) GetUsers() (*[]User, error) {
	query := "SELECT * FROM Users"

	var users []User

	rows, err := uModel.Database.Query(query)
	if err != nil {
		fmt.Println("Error querying table", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			fmt.Println("Error retrieving users", err)
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows", err)
		return nil, err
	}

	return &users, nil
}

func (uModel *UserModelImpl) CreateUser(user User) error {
	query := `INSERT INTO Users (email, password, role_id, created_at, updated_at)
				VALUES(?, ?, ?, ?, ?)`

	result, err := uModel.Database.Exec(query, user)

	if err != nil {
		fmt.Println("Error querying a database", err)
		return err
	}

	id, err := result.LastInsertId()

	fmt.Println("Last inserted id", id)

	if err != nil {
		return err
	}

	fmt.Println(id)

	return nil
}
