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
	GetMarkerPostImages(id int) ([]MarkerPostImage, error)
	CreateImage(markerPostImage []MarkerPostImage) error
}

type MarkerPostImageImpl struct {
	DB *sql.DB
}

func NewMarkerPostImageModel(db *sql.DB) *MarkerPostImageImpl {
	return &MarkerPostImageImpl{DB: db}
}

func (iModel *MarkerPostImageImpl) GetMarkerPostImages(id int) ([]MarkerPostImage, error) {

	var markerPostImages []MarkerPostImage

	query := `SELECT * FROM MarkerPostImage WHERE MarkerPostImage.marker_id = ?`

	rows, err := iModel.DB.Query(query, id)

	if err != nil {
		fmt.Println("error retrieving MarkerPostImages", err)
		return nil, err
	}

	for rows.Next() {
		var markerPostImage MarkerPostImage

		if err := rows.Scan(&markerPostImage.ID, &markerPostImage.Image,
			&markerPostImage.MimeType, &markerPostImage.MarkerID); err != nil {
			fmt.Println("error scanning MarkerPostImages", err)
			return nil, err
		}

		markerPostImages = append(markerPostImages, markerPostImage)
	}

	return markerPostImages, nil
}

func (iModel *MarkerPostImageImpl) CreateImages(markerPostImage []MarkerPostImage, markerID int) error {
	query := `INSERT INTO MarkerPostImage (image, mime_type, marker_id) VALUES(?, ?, ?)`

	for i := 0; i < len(markerPostImage); i++ {
		_, err := iModel.DB.Exec(query, markerPostImage[i].Image, markerPostImage[i].MimeType, markerID)

		if err != nil {
			fmt.Println("error inserting images", err)
			return err
		}
	}

	return nil
}
