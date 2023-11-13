package models

import (
	"database/sql"
	"fmt"
)

type FriendsWith struct {
	UserID      int `json:"userID" db:"user_id"`
	FriendsWith int `json:"UserID" db:"user_id"`
}

type FriendsWithModel interface {
	GetUserFriends(id int) ([]FriendsWith, error)
	DeleteFriend(id int) error
}

type FriendsWithImpl struct {
	DB *sql.DB
}

func NewFriendsWithModel(db *sql.DB) *FriendsWithImpl {
	return &FriendsWithImpl{DB: db}
}

func (fModel *FriendsWithImpl) GetUserFriends(id int) ([]FriendsWith, error) {
	query := `SELECT * FROM FriendsWith WHERE user_id = ?`

	var friends []FriendsWith

	rows, err := fModel.DB.Query(query, id)

	if err != nil {
		fmt.Println("Error querying database", err)
		return nil, err
	}

	for rows.Next() {
		var friend FriendsWith

		err = rows.Scan(&friend.UserID, &friend.FriendsWith)

		if err != nil {
			fmt.Println("Error getting friends of user", err)
			return nil, err
		}

		friends = append(friends, friend)
	}

	return friends, nil
}

func (fModel *FriendsWithImpl) DeleteFriend(id int) error {
	query := `DELETE FriendsWith WHERE friends_with = ?`

	_, err := fModel.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting", err)
		return err
	}

	return nil
}
