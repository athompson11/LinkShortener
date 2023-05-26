package link

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/athompson11/LinkShortener/src/database"
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
	print(givenid)
	row := db.DBConn.QueryRow("SELECT link FROM links WHERE linkid = ?", givenid)
	var redirectlink string
	err := row.Scan(&redirectlink)

	if err != nil {
		if err == sql.ErrNoRows {
			// 404, link not found
			fmt.Printf("Failed to prepare link: %v\n", err)
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		} else {
			// There was some other error with the database
			fmt.Printf("Failed to prepare link: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
	}
	c.Redirect(http.StatusMovedPermanently, redirectlink)
}
