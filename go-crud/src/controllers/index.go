package routes

import (
	"go-crud/src/databases/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetAlbums(c *gin.Context) {
	var albums []models.Albums
	DB.Find(&albums)
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Albums

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	DB.Create(&newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	var album models.Albums
	DB.Find(&album, "id = ?", id)
	if album.ID != 0 {
		c.IndentedJSON(http.StatusOK, album)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}

func DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")
	var album models.Albums
	DB.Find(&album, "id = ?", id)
	if album.ID != 0 {
		DB.Where("id = ?", id).Delete(&album)
		c.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "message": "deleted"})
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}
func UpdateAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var newAlbum models.Albums
	if err := ctx.BindJSON(&newAlbum); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Bad request"})
		return
	}
	var album models.Albums
	DB.First(&album, "id = ?", id)
	if album.ID != 0 {
		album.Artist = newAlbum.Artist
		album.Price = newAlbum.Price
		album.Title = newAlbum.Title
		DB.Save(&album)
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"status": 202, "message": "updated"})
		return
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"status": 404, "message": "album not found"})
}
