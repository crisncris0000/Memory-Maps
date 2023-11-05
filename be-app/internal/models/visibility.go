package models

import (
	"database/sql"
	"fmt"
)

type Visibility struct {
	ID   int    `json:"id" db:"id"`
	View string `json:"view" db:"view"`
}

type VisibilityModel interface {
}

type VisibilityImpl struct {
	DB *sql.DB
}

func NewVisibilityModel(db *sql.DB) *VisibilityImpl {
	return &VisibilityImpl{DB: db}
}

func (vModel *VisibilityImpl) GetVisibilityByID(id int) (*Visibility, error) {
	var visibility Visibility

	err := vModel.DB.QueryRow("SELECT * FROM visibility WHERE id = ?", id).Scan(&visibility.ID, &visibility.View)

	if err != nil {
		fmt.Println("Error searching for visibility", err)
		return nil, err
	}

	return &visibility, nil
}
