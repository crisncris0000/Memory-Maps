package models

type MarkerPostImage struct {
	ID       int    `json:"id" db:"id"`
	Image    []byte `json:"image" db:"image"`
	MarkerID int    `json:"markerID" db:"marker_id"`
}
