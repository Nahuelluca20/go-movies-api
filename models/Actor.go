package models

type Actor struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Movies []Movie `gorm:"many2many:actor_movies;"` // many to many relationship
}
