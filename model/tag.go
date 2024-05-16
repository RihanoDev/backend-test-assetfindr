package model

import (
	"time"
)

type Tag struct {
	ID        uint       `json:"id" form:"id" gorm:"primaryKey"`
	CreatedAt time.Time  `json:"created_at" form:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" form:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" form:"deleted_at"`
	Label     string     `json:"label" form:"label" gorm:"unique"`
	Posts     []Post     `json:"posts" form:"posts" gorm:"many2many:post_tags"`
}
