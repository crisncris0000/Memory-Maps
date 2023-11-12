package models

import "database/sql"

type FriendsWith struct {
	UserID      int `json:"userID" db:"user_id"`
	FriendsWith int `json:"UserID" db:"user_id"`
}

type FriendsWithModel interface {
}

type FriendsWithImpl struct {
	DB *sql.DB
}

func NewFriendsWithModel(db *sql.DB) *FriendsWithImpl {
	return &FriendsWithImpl{DB: db}
}
