package models

import (
	"database/sql"
	"fmt"
)

type Tags struct {
	TagID   int    `json:"id" db:"id"`
	TagName string `json:"tagName" db:"tag_name"`
}

type TagsModel interface {
	CreateTag(tagName string) error
	DeleteTag(id int) error
}

type TagsModelImpl struct {
	DB *sql.DB
}

func NewTagsModel(db *sql.DB) *TagsModelImpl {
	return &TagsModelImpl{DB: db}
}

func (tModel *TagsModelImpl) CreateTag(tagName string) error {
	query := `INSERT INTO Tags(tag_name) VALUES(?)`

	_, err := tModel.DB.Exec(query, tagName)

	if err != nil {
		fmt.Println("Error inserting into tag", err)
		return err
	}

	return nil
}

func (tModel *TagsModelImpl) DeleteTag(id int) error {
	query := `DELETE Tags WHERE tag_name = ?`

	_, err := tModel.DB.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting tag", err)
		return err
	}

	return nil
}
