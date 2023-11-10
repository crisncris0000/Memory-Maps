package models

import (
	"database/sql"
	"fmt"
	"log"
)

type PendingRequest struct {
	UserID      int `json:"userID" db:"user_id"`
	PendingUser int `json:"pendingUser" db:"pending_user"`
}

type PendingRequestModel interface {
}

type PendingRequestModelImpl struct {
	DB *sql.DB
}

func NewPendingRequestModel(db *sql.DB) *PendingRequestModelImpl {
	return &PendingRequestModelImpl{DB: db}
}

func (db *PendingRequestModelImpl) SendFriendRequest(pendingRequest PendingRequest) error {
	query := `INSERT INTO PendingRequest(user_id, pending_user) VALUES(?, ?)`

	_, err := db.DB.Exec(query, pendingRequest.UserID, pendingRequest.PendingUser)

	if err != nil {
		log.Println("Error Inserting into PendingRequestTable", err)
		return err
	}

	return nil
}

func (db *PendingRequestModelImpl) DeclineFriendRequest(id int) error {
	query := `DELETE FROM PendingRequest WHERE user_id = ?`

	_, err := db.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting Pending Request", err)
		return err
	}

	return nil
}
