package models

import (
	"database/sql"
	"fmt"
	"time"
)

type Comments struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"firstName" db:"first_name"`
	LastName  string    `json:"lastName" db:"last_name"`
	MarkerID  int       `json:"markerID" db:"marker_id"`
	Comment   string    `json:"comment" db:"comment"`
	Likes     int       `json:"likes" db:"likes"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CommentsModel interface {
	CreateComment(comment Comments) error
	GetAllComments() ([]Comments, error)
	UpdateComment(comment string) error
	DeleteCommentByID(id int) error
}

type CommentsModelImpl struct {
	DB *sql.DB
}

func NewCommentsModel(db *sql.DB) *CommentsModelImpl {
	return &CommentsModelImpl{DB: db}
}

func (cModel *CommentsModelImpl) GetAllComments() ([]Comments, error) {
	query := `SELECT * FROM Comments`
	var comments []Comments

	rows, err := cModel.DB.Query(query)

	if err != nil {
		fmt.Println("Error retrieving comments", err)
		return nil, err
	}

	for rows.Next() {
		var comment Comments

		err = rows.Scan(&comment.ID, &comment.FirstName, &comment.LastName, &comment.MarkerID, &comment.Comment,
			&comment.Comment)

		if err != nil {
			fmt.Println("Error scanning row", err)
			return nil, err
		}
	}

	return comments, nil
}

func (cModel *CommentsModelImpl) GetCommentsByMarkerID(id int) ([]Comments, error) {
	query := `SELECT * FROM Comments WHERE marker_id = ?`

	var comments []Comments

	rows, err := cModel.DB.Query(query, id)

	if err != nil {
		fmt.Println("Error querying Comments by ID", err)
		return nil, err
	}

	for rows.Next() {
		var comment Comments

		err = rows.Scan(&comment.ID, &comment.FirstName, &comment.LastName, &comment.MarkerID, &comment.Comment,
			&comment.Likes, &comment.CreatedAt, &comment.UpdatedAt)

		if err != nil {
			fmt.Println("Error Scanning comment by marker ID", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (cModel *CommentsModelImpl) CreateComment(comment Comments) error {
	query := `INSERT INTO Comments (first_name, last_name, marker_id, comment, likes, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`

	createdAt := time.Now()
	updatedAt := time.Now()

	_, err := cModel.DB.Exec(query, comment.FirstName, comment.LastName,
		comment.MarkerID, comment.Comment, comment.Likes, createdAt, updatedAt)

	if err != nil {
		fmt.Println("Error inserting into table comments", err)
		return err
	}

	return nil
}

func (cModel *CommentsModelImpl) UpdateComment(comment string) error {
	query := "UPDATE Comments SET comment = ?, updated_at = ?"

	updatedAt := time.Now()

	_, err := cModel.DB.Exec(query, comment, updatedAt)

	if err != nil {
		fmt.Println("Error updating Comment", err)
		return err
	}

	return nil
}

func (cModel *CommentsModelImpl) DeleteCommentByID(id int) error {
	query := "DELETE Comments WHERE id = ?"

	_, err := cModel.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting comment with the id: ", id, err)
		return err
	}

	return nil
}
