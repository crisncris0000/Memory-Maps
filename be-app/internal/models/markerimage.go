package models

import (
	"database/sql"
	"fmt"
)

type MarkerPostImage struct {
	ID       int    `json:"id" db:"id"`
	Image    []byte `json:"image" db:"image"`
	MarkerID int    `json:"markerID" db:"marker_id"`
}

type MarkerPostImageModel interface {
}

type MarkerPostImageImpl struct {
	DB *sql.DB
}

func NewMarkerPostImageModel(db *sql.DB) *MarkerPostImageImpl {
	return &MarkerPostImageImpl{DB: db}
}

func (iModel *MarkerPostImageImpl) CreateSingleImage(image []byte, markerID int) error {
	query := `INSERT INTO MarkerPostImage (image, marker_id) VALUES(?, ?)`

	_, err := iModel.DB.Exec(query, image, markerID)

	if err != nil {
		fmt.Println("error inserting image into marker post", err)
		return err
	}

	return nil
}
