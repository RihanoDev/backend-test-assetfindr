package model

import (
	"time"
)

type Post struct {
	ID        uint       `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`
	Title     string     `json:"title" form:"title"`
	Content   string     `json:"content" form:"content"`
	Tags      []Tag      `json:"tags" form:"tags" gorm:"many2many:post_tags"`
}
