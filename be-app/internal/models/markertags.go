package models

import (
	"database/sql"
	"fmt"
)

type MarkerTags struct {
	TagID        int `json:"tagID" db:"tag_id"`
	MarkerPostID int `json:"markerPostID" db:"marker_id"`
}

type MarkerTagsModel interface {
	CreateMarkerTag(markerID, tagID int) error
	DeleteMarkerTag(tagID int) error
}

type MarkerTagsImpl struct {
	DB *sql.DB
}

func NewMarkerTagsModel(db *sql.DB) *MarkerTagsImpl {
	return &MarkerTagsImpl{DB: db}
}

func (mtModel *MarkerTagsImpl) CreateMarkerTag(markerID, tagID int) error {
	query := `INSERT INTO MarkerPostTags VALUES (marker_id, tag_id) VALUES (?, ?)`

	_, err := mtModel.DB.Exec(query, markerID, tagID)

	if err != nil {
		fmt.Println("Error inserting into MarkerPostTags", err)
		return err
	}

	return nil
}

func (mtModel *MarkerTagsImpl) DeleteMarkerTag(tagID int) error {
	query := `DELETE MarkerPostTag WHERE tag_id = ?`

	_, err := mtModel.DB.Exec(query, tagID)

	if err != nil {
		fmt.Println("Error deleting in MarkerPostTag", err)
		return err
	}

	return nil
}
