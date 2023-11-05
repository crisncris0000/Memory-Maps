package models

import (
	"database/sql"
	"fmt"
	"time"
)

type MarkerPost struct {
	ID          int       `json:"id" db:"id"`
	Lattitude   float32   `json:"latitude" db:"latitude"`
	Longitude   float32   `json:"longitude" db:"longitude"`
	Image       []byte    `json:"image" db:"image"`
	Description string    `json:"description" db:"description"`
	Likes       int       `json:"likes" db:"likes"`
	Visibility  int       `json:"visibility" db:"visibility"`
	UserId      int       `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type MarkerPostModel interface {
	CreateMarkerPost(post MarkerPost) error
	GetMarkerPosts() ([]MarkerPost, error)
}

type MarkerPostImpl struct {
	*sql.DB
}

func NewMarkerPost(db *sql.DB) *MarkerPostImpl {
	return &MarkerPostImpl{DB: db}
}

func (mModel *MarkerPostImpl) CreateMarkerPost(post MarkerPost) error {
	query := `INSERT INTO MarkerPost(latitude, longitude, image, description, likes, visibility, user_id, created_at, updated_at)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	created := time.Now()
	updated := time.Now()

	fmt.Println(created, updated)

	res, err := mModel.Exec(query, post.Lattitude, post.Longitude, post.Image, post.Description,
		post.Likes, post.Visibility, post.UserId, created, updated)

	if err != nil {
		fmt.Println("Error inserting marker post within the database", err)
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println("Error getting id of last inserted", err)
	}

	fmt.Println("Last inserted id", id)

	return nil
}

func (mModel *MarkerPostImpl) GetMarkerPosts() ([]MarkerPost, error) {
	query := `SELECT * FROM MarkerPost`

	var posts []MarkerPost

	rows, err := mModel.Query(query)

	if err != nil {
		fmt.Println("Error getting posts", err)
	}

	for rows.Next() {
		var post MarkerPost

		if err := rows.Scan(&post.ID, &post.Lattitude, &post.Longitude, &post.Image, post.Description,
			&post.Likes, &post.Visibility, &post.UserId, &post.CreatedAt, &post.UpdatedAt); err != nil {

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

	rows, err := mModel.Query(query, startDate, endDate)

	if err != nil {
		var post MarkerPost

		if err := rows.Scan(&post.ID, &post.Lattitude, &post.Longitude, &post.Image, post.Description,
			&post.Likes, &post.Visibility, &post.UserId, &post.CreatedAt, &post.UpdatedAt); err != nil {

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

func (mModel *MarkerPostImpl) UpdatePost(post MarkerPost) error {
	query := `UPDATE MarkerPost SET latitude = ?, longitude = ?, image = ?, 
	description = ?, likes = ?, visibility = ?, user_id = ?, created_at = ?, updated_at = ? WHERE id = ?`

	_, err := mModel.Exec(query, post.Lattitude, post.Longitude, post.Image, post.Description,
		post.Likes, post.Visibility, post.UserId, post.CreatedAt, post.UpdatedAt, post.ID)

	if err != nil {
		fmt.Println("Error updating post", err)
		return err
	}

	return nil
}

func (mModel *MarkerPostImpl) DeletePost(id int) error {
	query := `DELETE FROM MarkerPost WHERE id = ?`

	_, err := mModel.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting post", err)
		return err
	}

	return nil
}
