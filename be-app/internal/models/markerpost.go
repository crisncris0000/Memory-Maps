package models

import (
	"database/sql"
	"fmt"
	"time"
)

type MarkerPost struct {
	ID           int       `json:"id" db:"id"`
	Lattitude    float32   `json:"latitude" db:"latitude"`
	Longitude    float32   `json:"longitude" db:"longitude"`
	Description  string    `json:"description" db:"description"`
	Likes        int       `json:"likes" db:"likes"`
	VisibilityID int       `json:"visibilityID" db:"visibility_id"`
	UserID       int       `json:"userID" db:"user_id"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

type MarkerPostModel interface {
	CreateMarkerPost(post MarkerPost) error
	GetMarkerPosts() ([]MarkerPost, error)
	GetPostsByDate(startDate, endDate time.Time) ([]MarkerPost, error)
	UpdatePost(post MarkerPost) error
	DeletePost(id int) error
}

type MarkerPostImpl struct {
	DB *sql.DB
}

func NewMarkerPost(db *sql.DB) *MarkerPostImpl {
	return &MarkerPostImpl{DB: db}
}

func (mModel *MarkerPostImpl) GetMarkerPosts() ([]MarkerPost, error) {
	query := `SELECT * FROM MarkerPost`

	var posts []MarkerPost

	rows, err := mModel.DB.Query(query)

	if err != nil {
		fmt.Println("Error getting posts", err)
	}

	for rows.Next() {
		var post MarkerPost

		if err := rows.Scan(&post.ID, &post.Lattitude, &post.Longitude, &post.Description,
			&post.Likes, &post.VisibilityID, &post.UserID, &post.CreatedAt, &post.UpdatedAt); err != nil {

			fmt.Println("Error getting posts", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows", err)
		return nil, err
	}

	return posts, nil
}

func (mModel *MarkerPostImpl) GetPostsByDate(startDate, endDate time.Time) ([]MarkerPost, error) {
	query := `SELECT * FROM MarkerPost WHERE created_at >= ?  AND ? <= created_at`

	var posts []MarkerPost

	rows, err := mModel.DB.Query(query, startDate, endDate)

	if err != nil {
		var post MarkerPost

		if err := rows.Scan(&post.ID, &post.Lattitude, &post.Longitude, post.Description,
			&post.Likes, &post.VisibilityID, &post.UserID, &post.CreatedAt, &post.UpdatedAt); err != nil {

			fmt.Println("Error getting posts", err)
			return nil, err
		}

		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows", err)
		return nil, err
	}

	return posts, nil
}

func (mModel *MarkerPostImpl) CreateMarkerPost(post MarkerPost) (int, error) {

	query := `INSERT INTO MarkerPost(latitude, longitude, description, likes, visibility_id, user_id, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	created := time.Now()
	updated := time.Now()

	res, err := mModel.DB.Exec(query, post.Lattitude, post.Longitude, post.Description,
		post.Likes, post.VisibilityID, post.UserID, created, updated)

	if err != nil {
		fmt.Println("Error inserting marker post within the database", err)
		return -1, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println("Error getting id of last inserted", err)
	}

	return int(id), nil
}

func (mModel *MarkerPostImpl) UpdatePost(post MarkerPost) error {
	query := `UPDATE MarkerPost SET likes = ? WHERE id = ?`

	_, err := mModel.DB.Exec(query, post.Likes, post.ID)

	if err != nil {
		fmt.Println("Error updating post", err)
		return err
	}

	return nil
}

func (mModel *MarkerPostImpl) DeletePost(id int) error {
	imageQuery := `DELETE FROM MarkerPostImage WHERE marker_id = ?`

	_, err := mModel.DB.Exec(imageQuery, id)

	if err != nil {
		fmt.Println("Error deleting images", err)
		return err
	}

	commentsQuery := `DELETE FROM Comments WHERE marker_id = ?`

	_, err = mModel.DB.Exec(commentsQuery, id)

	if err != nil {
		fmt.Println("Error deleting comments", err)
		return err
	}

	markerQuery := `DELETE FROM MarkerPost WHERE id = ?`

	_, err = mModel.DB.Exec(markerQuery, id)

	if err != nil {
		fmt.Println("Error deleting post", err)
		return err
	}

	return nil
}
