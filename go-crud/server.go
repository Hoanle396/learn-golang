package main

import (
	"flag"
	routes "go-crud/src/controllers"
	"go-crud/src/databases"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var addr = flag.String("addr", ":8080", "http service address")

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	var db, e = databases.InitDB()
	if e != nil {
		log.Fatal("Error loading database")
	}
	routes.DB = db

	router := gin.Default()
	router.GET("/albums", routes.GetAlbums)
	router.GET("/albums/:id", routes.GetAlbumByID)
	router.POST("/albums", routes.PostAlbums)
	router.PUT("/albums/:id", routes.UpdateAlbumByID)
	router.DELETE("/albums/:id", routes.DeleteAlbumById)
	log.Fatal(router.Run(*addr))
}
