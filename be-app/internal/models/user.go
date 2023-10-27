package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int       `json:"id" db:"id"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Role_id    int       `json:"role_id" db:"role_id"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
}

type UserModel interface {
}

type UserModelImpl struct {
	*sql.DB
}

func newUserModel(db *sql.DB) *UserModelImpl {
	return &UserModelImpl{DB: db}
}
