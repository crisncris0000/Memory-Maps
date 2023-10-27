package main

import (
	"log"

	"github.com/crisncris0000/Memory-Maps/be-app/config"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/handlers"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/models"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/pkg/db"
	"github.com/crisncris0000/Memory-Maps/be-app/internal/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	cf, err := config.StartConfiguration()

	if err != nil {
		log.Fatal(err)
	}

	database, err := db.Connect(cf.Database.Drivername, cf.Database.DataSourceName)
	defer database.Close()

	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	uModel := models.NewUserModel(database)
	uHandler := handlers.NewUserHandler(uModel)
	uRouter := routes.NewUserRouter(uHandler)

	rModel := models.NewRoleModel(database)
	rHandler := handlers.NewRoleHandler(rModel)
	rRouter := routes.NewRoleRouter(rHandler)

	uRouter.InitializeUserRouter(r)
	rRouter.InitializeRouter(r)

	log.Fatal(r.Run(cf.Port))
}
