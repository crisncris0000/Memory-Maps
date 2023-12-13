package models

import (
	"database/sql"
	"fmt"
)

type MarkerPostImage struct {
	ID       int    `json:"id" db:"id"`
	Image    []byte `json:"image" db:"image"`
	MimeType string `json:"mimeType" db:"mime_type"`
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

func (iModel *MarkerPostImageImpl) CreateSingleImage(markerPostImage MarkerPostImage) error {
	query := `INSERT INTO MarkerPostImage (image, mime_type, marker_id) VALUES(?, ?, ?)`

	_, err := iModel.DB.Exec(query, markerPostImage.Image, markerPostImage.MimeType, markerPostImage.MarkerID)

	if err != nil {
		fmt.Println("error inserting image", err)
		return err
	}

	return nil
}

func (iModel *MarkerPostImageImpl) CreateMultipleImages(markerPostImage []MarkerPostImage) error {
	query := `INSERT INTO MarkerPostImage (image, mime_type, marker_id) VALUES(?, ?, ?)`

	for i := 0; i < len(markerPostImage); i++ {
		_, err := iModel.DB.Exec(query, markerPostImage[i].Image, markerPostImage[i].MimeType, markerPostImage[i].MarkerID)

		if err != nil {
			fmt.Println("error inserting images", err)
			return err
		}
	}

	return nil
}
