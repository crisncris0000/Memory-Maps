package models

import (
	"database/sql"
	"fmt"
)

type ResetToken struct {
	ID     int    `json:"id" db:"id"`
	UserID int    `json:"userID" db:"user_id"`
	Token  string `json:"token" db:"token"`
}

type ResetTokenModel interface {
}

type ResetTokenImpl struct {
	DB *sql.DB
}

func NewResetTokenModel(db *sql.DB) *ResetTokenImpl {
	return &ResetTokenImpl{DB: db}
}

func (rt *ResetTokenImpl) GetResetToken(token string, id int) (string, error) {
	query := `SELECT * FROM ResetToken WHERE token = ? AND WHERE user_id = ?`

	var resetToken ResetToken

	row, err := rt.DB.Query(query, token, id)

	if err != nil {
		fmt.Println("Error querying the database", err)
		return "", err
	}

	row.Scan(&resetToken.ID, &resetToken.Token, &resetToken.UserID)

	return resetToken.Token, nil
}

func (rt *ResetTokenImpl) CreateResetToken(resetToken ResetToken) error {
	query := `INSERT INTO ResetToken(user_id, token) VALUES(?, ?)`

	_, err := rt.DB.Exec(query, resetToken.UserID, resetToken.Token)

	if err != nil {
		fmt.Println("Error inserting into database", err)
		return err
	}

	return nil
}

func (rt *ResetTokenImpl) DeleteResetToken(id int) error {
	query := `DELETE FROM ResetToken WHERE ResetToken.id = ?`

	_, err := rt.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting from database", err)
		return err
	}

	return nil
}
