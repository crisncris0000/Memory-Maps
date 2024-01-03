package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/crisncris0000/Memory-Maps/be-app/internal/utils"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	Password  string    `json:"password" db:"password"`
	RoleID    int       `json:"role_id" db:"role_id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type UserModel interface {
	GetUsers() ([]User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(user User) error
	UserExists(email string) error
	SendUserEmail(from, subject, body string)
}

type UserModelImpl struct {
	DB *sql.DB
}

func NewUserModel(db *sql.DB) *UserModelImpl {
	return &UserModelImpl{DB: db}
}

func (uModel *UserModelImpl) GetUsers() (*[]User, error) {
	query := "SELECT * FROM Users"

	var users []User

	rows, err := uModel.DB.Query(query)
	if err != nil {
		fmt.Println("Error querying table", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

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

func (uModel *UserModelImpl) GetUserByEmail(email string) (*User, error) {
	query := `SELECT * FROM Users WHERE Users.email = ?`

	var user User

	err := uModel.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println("User does not exist", err)
		return nil, err
	}

	return &user, nil
}

func (uModel *UserModelImpl) GetUserByID(id int) (*User, error) {
	query := `SELECT * FROM Users WHERE Users.id = ?`

	var user User

	err := uModel.DB.QueryRow(query, id).Scan(&user.ID, &user.Email, &user.Password, &user.RoleID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		fmt.Println("User not found with the id", id)
		return nil, err
	}

	return &user, err
}

func (uModel *UserModelImpl) CreateUser(user User) error {
	query := `INSERT INTO Users (email, first_name, last_name, password, role_id, created_at, updated_at)
				VALUES(?, ?, ?, ?, ?, ?, ?)`

	createdAt := time.Now()
	updatedAt := time.Now()

	exists, err := uModel.UserExists(user.Email)

	if err != nil {
		fmt.Println("error checking count ", err)
		return err
	}

	if exists {
		fmt.Println("User already exist")
		return errors.New("User exist")
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	_, err = uModel.DB.Exec(query, user.Email, user.FirstName, user.LastName, hashedPassword, user.RoleID, createdAt, updatedAt)

	if err != nil {
		fmt.Println("Error querying a database", err)
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func (uModel *UserModelImpl) UserExists(email string) (bool, error) {
	query := `SELECT COUNT(*) FROM Users WHERE Users.email = ?`

	var count int
	err := uModel.DB.QueryRow(query, email).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
