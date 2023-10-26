package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(driverName, dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Println("Error pinging database", err)
		return nil, err
	}

	return db, nil
}
