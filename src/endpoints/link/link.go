package link

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/athompson11/LinkShortener/src/database"
	"github.com/gin-gonic/gin"
)

func RegisterLinkRoutes(r *gin.Engine) {
	r.GET("/link", redirectCreation)
	r.GET("/link/:id", getLink)
}

func redirectCreation(c *gin.Context) {
	// If user goes to /link without an ID, redirect to create
	c.Redirect(http.StatusMovedPermanently, "/create")
}

func getLink(c *gin.Context) {
	givenid := c.Param("id")
	row := database.DBConn.QueryRow("SELECT link FROM links WHERE LinkID = ?", givenid)

	var redirectlink string
	err := row.Scan(&redirectlink)

	if err != nil {
		if err == sql.ErrNoRows {
			// 404, link not found
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			// There was some other error with the database
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
	}
	c.JSON(http.StatusMovedPermanently, redirectlink)
}
