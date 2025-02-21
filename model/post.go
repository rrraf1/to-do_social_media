package model_todoSM

import (
	"gorm.io/gorm"
	"time"
)

type PostInput struct {
	Title    *string `json:"title" example:"My Post"`
	Brand    *string `json:"brand" example:"Brand A"`
	Platform *string `json:"platform" example:"Platform X"`
	DueDate  *string `json:"due_date" example:"2025-01-10"` // Harus sebagai string di input
}

type PostResponse struct {
	Id       int    `json:"id" example:"1"`
	Title    string `json:"title" example:"My Post"`
	Brand    string `json:"brand" example:"Brand A"`
	Platform string `json:"platform" example:"Platform X"`
	DueDate  string `json:"due_date" example:"2025-01-10"` // Harus dalam format YYYY-MM-DD
}

type Post struct {
	Id       int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title    string    `json:"title"`
	Brand    string    `json:"brand"`
	Platform string    `json:"platform"`
	DueDate  time.Time `gorm:"type:date" json:"due_date"`
}

func MigratePost(db *gorm.DB) error {
	return db.AutoMigrate(&Post{})
}
