package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase() // Initialize the database connection
	defer DBConn.Close()
	r := gin.Default()

	RegisterLinkRoutes(r)
	RegisterCreationRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
