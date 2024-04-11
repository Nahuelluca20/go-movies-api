package models

import "gorm.io/gorm"

type Actor struct {
	gorm.Model

	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Movies []Movie `gorm:"many2many:actor_movies;"` // many to many relationship
}
