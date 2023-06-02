package databases

import (
	"fmt"
	"go-crud/src/databases/models"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

var albums = []models.Albums{
	{Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// This function will make a connection to the database only once.
func InitDB() (*gorm.DB, error) {
	var err error

	dsn := url.URL{
		User:     url.UserPassword(USER, PASSWORD),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", HOST, PORT),
		Path:     DATABASE,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	var albumsList []models.Albums

	db.AutoMigrate(&models.Albums{}, &models.User{}, &models.Project{}, &models.Space{}, &models.Task{})

	db.Find(&albumsList)
	if len(albumsList) <= 0 {
		for index := range albums {
			db.Create(&albums[index])
		}
	}

	fmt.Println("The database is connected")
	return db, nil
}
