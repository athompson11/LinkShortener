package main

import (
	"github.com/athompson11/LinkShortener/src/database"
	"github.com/athompson11/LinkShortener/src/endpoints/creation"
	"github.com/athompson11/LinkShortener/src/endpoints/link"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase() // Initialize the database connection
	defer database.DBConn.Close()
	r := gin.Default()

	link.RegisterLinkRoutes(r)
	creation.RegisterCreationRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
