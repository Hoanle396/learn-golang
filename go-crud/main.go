package main

import (
	"flag"
	routes "go-crud/src/controllers"
	"go-crud/src/databases"
	"go-crud/src/middleware"
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

	router := initRouter()
	log.Fatal(router.Run(*addr))
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/login", routes.Login)
		api.POST("/register", routes.RegisterUser)
		secured := api.Use(middleware.Auth())
		{
			secured.GET("/me", routes.Me)
			secured.GET("/albums", routes.GetAlbums)
			secured.GET("/albums/:id", routes.GetAlbumByID)
			secured.POST("/albums", routes.PostAlbums)
			secured.PUT("/albums/:id", routes.UpdateAlbumByID)
			secured.DELETE("/albums/:id", routes.DeleteAlbumById)
			secured.GET("/project/:id", routes.GetProjectWithSpaces)
		}
	}
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"status": 404, "message": "404 - Page Not Found"})
		ctx.Abort()
		return
	})
	return router
}
