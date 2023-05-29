package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}

func DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	indexToRemove := -1
	for i, item := range albums {
		if item.ID == id {
			indexToRemove = i
			break
		}
	}
	if indexToRemove != -1 {
		albums = append(albums[:indexToRemove], albums[indexToRemove+1:]...)
		c.IndentedJSON(http.StatusOK, albums)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}
func UpdateAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var newAlbum album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Bad request"})
		return
	}

	indexToUpdate := -1
	for i, item := range albums {
		if item.ID == id {
			indexToUpdate = i
			break
		}
	}
	if indexToUpdate != -1 {
		albums[indexToUpdate] = newAlbum
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "message": "Updated successfully"})
		return
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}
