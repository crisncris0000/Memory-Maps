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
	GetUserPendingRequests(id int) ([]PendingRequest, error)
	SendFriendRequest(pendingRequest PendingRequest) error
	DeclineFriendRequest(id int) error
}

type PendingRequestModelImpl struct {
	DB *sql.DB
}

func NewPendingRequestModel(db *sql.DB) *PendingRequestModelImpl {
	return &PendingRequestModelImpl{DB: db}
}

func (db *PendingRequestModelImpl) GetUserPendingRequests(id int) ([]PendingRequest, error) {
	query := `SELECT * FROM PendingRequest`

	var requests []PendingRequest

	rows, err := db.DB.Query(query)

	if err != nil {
		fmt.Println("Error querying database", err)
		return nil, err
	}

	for rows.Next() {
		var request PendingRequest

		err = rows.Scan(&request.UserID, &request.PendingUser)

		if err != nil {
			fmt.Println("Error scanning into PendingRequest", err)
			return nil, err
		}
		requests = append(requests, request)
	}

	return requests, nil
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
