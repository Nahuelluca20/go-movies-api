package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model

	ID       uint    `json:"id" gorm:"primaryKey;not null"`
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	Director string  `json:"director"`
	Desc     string  `json:"desc"`
	Subtitle string  `json:"subtitle"`
	Actor    []Actor `gorm:"many2many:actor_movies;"` // many to many relationship
}
