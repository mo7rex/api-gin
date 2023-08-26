package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}
