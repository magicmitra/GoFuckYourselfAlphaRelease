package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album
var albums = []album{
	{ID: "1", Title: "Fractals", Artist: "Subtronics", Price: 6.99},
	{ID: "2", Title: "Pay Attention EP", Artist: "um..", Price: 20.99},
}

func main() {
	router := gin.Default()

	// router
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

// getAlbums responds with the list of all albums as JSON
// AKA handler/controller
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
