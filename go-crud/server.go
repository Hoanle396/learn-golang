package main

import (
	"flag"
	routes "go-crud/src/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	router := gin.Default()
	router.GET("/albums", routes.GetAlbums)
	router.GET("/albums/:id", routes.GetAlbumByID)
	router.POST("/albums", routes.PostAlbums)
	router.PUT("/albums/:id", routes.UpdateAlbumByID)
	router.DELETE("/albums/:id", routes.DeleteAlbumById)
	log.Fatal(router.Run(*addr))
}
