package models

import "database/sql"

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
