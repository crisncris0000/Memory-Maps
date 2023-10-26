package main

import (
	"log"

	"github.com/crisncris0000/Memory-Maps/be-app/config"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/pkg/db"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cf, err := config.StartConfiguration()

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Connect(cf.Database.Drivername, cf.Database.DataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Run(cf.Port)
}
