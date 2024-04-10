package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model

	Title    string   `json:"title"`
	Year     int      `json:"year"`
	Director string   `json:"director"`
	Desc     string   `json:"desc"`
	Subtitle string   `json:"subtitle"`
	Actors   []*Actor `gorm:"many2many:actor_movies;"` // many to many relationship
}
