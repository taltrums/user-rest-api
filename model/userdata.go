package model

import (
	"gorm.io/gorm"
)

type Details struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Image  []byte `json:"image"`
	Pdf    []byte `json:"pdf"`
	UserID uint
}
