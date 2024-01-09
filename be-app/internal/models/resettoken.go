package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type ResetToken struct {
	ID    int    `json:"id" db:"id"`
	Email string `json:"email" db:"user_email"`
	Token string `json:"token" db:"token"`
}

type ResetTokenModel interface {
	GetResetToken(token, email string) (string, error)
	CreateResetToken(resetToken ResetToken) error
	DeleteResetTokenByID(id int) error
}

type ResetTokenImpl struct {
	DB *sql.DB
}

func NewResetTokenModel(db *sql.DB) *ResetTokenImpl {
	return &ResetTokenImpl{DB: db}
}

func (rt *ResetTokenImpl) GetResetToken(token, email string) (string, error) {
	query := `SELECT * FROM ResetToken WHERE token = ? AND user_email = ?`

	var resetToken ResetToken

	row := rt.DB.QueryRow(query, token, email)

	err := row.Scan(&resetToken.ID, &resetToken.Email, &resetToken.Token)

	if err == sql.ErrNoRows {
		return "", errors.New("Token not found")
	} else if err != nil {
		fmt.Println("Error scanning into struct", err)
		return "", err
	}

	return resetToken.Token, nil
}

func (rt *ResetTokenImpl) CreateResetToken(resetToken ResetToken) error {
	query := `INSERT INTO ResetToken(user_email, token) VALUES(?, ?)`

	_, err := rt.DB.Exec(query, resetToken.Email, resetToken.Token)

	if err != nil {
		fmt.Println("Error inserting into database", err)
		return err
	}

	return nil
}

func (rt *ResetTokenImpl) DeleteResetTokenByID(id int) error {
	query := `DELETE FROM ResetToken WHERE id = ?`

	_, err := rt.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting from database", err)
		return err
	}

	return nil
}
