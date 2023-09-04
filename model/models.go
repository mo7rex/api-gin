package model

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Email       string  `json:"email" validate:"required" gorm:"unique"`
	Password    string  `json:"password" validate:"required"`
	FullName    string  `json:"full_name" validate:"required"`
	PhoneNumber string  `json:"phone_number" validate:"required,min=4,max=13"`
	Gender      string  `json:"gender" validate:"required,eq=MALE|eq=FEMALE"`
	Post        []*Post `json:"post,omitempty" gorm:"constraint:OnDelete:CASCADE;"`
}
type Post struct {
	gorm.Model
	AccountID uint   `json:"account_id"`
	Title     string `json:"title,omitempty"`
	Body      string `json:"body,omitempty"`
}
