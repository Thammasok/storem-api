package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// Inject fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` into model `User`
type User struct {
	gorm.Model
	Name string
}

// Declaring model w/o gorm.Model
// type User struct {
// 	ID   int
// 	Name string
// }
