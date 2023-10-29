package models

import (
	"database/sql"
	"fmt"
)

type Role struct {
	ID       int    `json:"id"`
	RoleName string `json:"roleName" db:"role_name"`
}

type RoleModel interface {
}

type RoleModelImpl struct {
	*sql.DB
}

func NewRoleModel(db *sql.DB) *RoleModelImpl {
	return &RoleModelImpl{DB: db}
}

func (rModel *RoleModelImpl) GetRole(id int) (*Role, error) {
	query := "SELECT * FROM Roles WHERE id = ?"

	var role Role

	err := rModel.DB.QueryRow(query, id).Scan(&role.ID, &role.RoleName)

	if err != nil {
		fmt.Println("Error converting to role struct", err)
		return nil, err
	}

	return &role, nil
}
