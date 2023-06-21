package creation

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/athompson11/LinkShortener/src/database"
	"github.com/athompson11/LinkShortener/src/structs"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func RegisterCreationRoutes(r *gin.Engine) {
	r.POST("/create", createLink)
}

func createLink(c *gin.Context) {
	//Take JSON from frontend, create Link, frontend should handle fields and sizes
	//Table ID (autogenned)- Link ID (autogenned if left blank) - Link - DateTime
	//
	var link structs.Link
	if err := c.BindJSON(&link); err != nil {
		fmt.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if link.LinkID == "" {
		link.LinkID = String(16)
	}
	stmt, err := db.DBConn.Prepare("INSERT INTO links(linkid, link, creationtime) VALUES(?, ?, ?)")
	if err != nil {
		fmt.Printf("Failed to prepare JSON: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if stmt == nil {
		print(link.LinkID)
		print(link.Link)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(link.LinkID, link.Link, link.CreatedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "Created", "link": "https://shortener.mymanyprojects.dev/link/" + link.LinkID})
}
