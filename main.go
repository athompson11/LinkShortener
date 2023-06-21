package main

import (
	"os"

	"github.com/athompson11/LinkShortener/src/database"
	"github.com/athompson11/LinkShortener/src/endpoints/creation"
	"github.com/athompson11/LinkShortener/src/endpoints/link"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDatabase(os.Getenv("DATABASE_LOCATION")) // Initialize the database connection
	defer db.DBConn.Close()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	link.RegisterLinkRoutes(r)
	creation.RegisterCreationRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
