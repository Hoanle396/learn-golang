package models

import "gorm.io/gorm"

type Albums struct {
	gorm.Model

	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
