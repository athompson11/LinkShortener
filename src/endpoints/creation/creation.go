package creation

import (
	"github.com/gin-gonic/gin"
)

func RegisterCreationRoutes(r *gin.Engine) {
	r.GET("/create", serveCreationPage)
	r.POST("/create", createLink)
}

func serveCreationPage(c *gin.Context) {
	//Serve page from template, todo: make frontend, maybe angular?
}

func createLink(c *gin.Context) {
	//Take JSON from frontend, create Link, frontend should handle fields and sizes
}
