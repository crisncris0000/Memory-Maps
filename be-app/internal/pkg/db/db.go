package db

import (
	"database/sql"
	"fmt"
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
