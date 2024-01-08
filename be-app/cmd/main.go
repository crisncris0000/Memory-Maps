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

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	uModel := models.NewUserModel(database)
	uHandler := handlers.NewUserHandler(uModel)
	uRouter := routes.NewUserRouter(uHandler)

	rModel := models.NewRoleModel(database)
	rHandler := handlers.NewRoleHandler(rModel)
	rRouter := routes.NewRoleRouter(rHandler)

	mModel := models.NewMarkerPost(database)
	mHandler := handlers.NewMarkerPostHandler(mModel)
	mRouter := routes.NewMarkerPostRouter(mHandler)

	iModel := models.NewMarkerPostImageModel(database)
	iHandler := handlers.NewMarkerPostImageHandler(iModel)
	iRouter := routes.NewMarkerPostImageRouter(iHandler)

	vModel := models.NewVisibilityModel(database)
	vHandler := handlers.NewVisibilityHandler(vModel)
	vRouter := routes.NewVisibilityRouter(vHandler)

	pModel := models.NewPendingRequestModel(database)
	pHandler := handlers.NewPendingRequestHandler(pModel)
	pRouter := routes.NewPendingRequestRouter(pHandler)

	fModel := models.NewFriendsWithModel(database)
	fHandler := handlers.NewFriendsWithHandler(fModel)
	fRouter := routes.NewFriendsWithRouter(fHandler)

	cModel := models.NewCommentsModel(database)
	cHandler := handlers.NewCommentsHandler(cModel)
	cRouter := routes.NewCommentsRouter(cHandler)

	rtModel := models.NewResetTokenModel(database)
	rtHandler := handlers.NewResetTokenHandler(rtModel)
	rtRouter := routes.NewResetTokenRouter(rtHandler)

	uRouter.InitializeUserRouter(r)
	rRouter.InitializeRouter(r)
	mRouter.InitializeRouter(r)
	iRouter.InitializeRoutes(r)
	vRouter.InitializeRouter(r)
	pRouter.InitializeRouter(r)
	fRouter.InitializeRouter(r)
	cRouter.InitializeRouter(r)
	rtRouter.InitializeRouter(r)

	log.Fatal(r.Run(cf.Port))
}
