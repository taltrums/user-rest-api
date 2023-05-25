package model

import (
	"rest-api/database"

	"gorm.io/gorm"
)

type Detail struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Image  []byte `json:"image"`
	Pdf    []byte `json:"pdf"`
	UserID uint
}

func (detail *Detail) Save() (*Detail, error) {
	err := database.Database.Create(&detail).Error
	if err != nil {
		return &Detail{}, err
	}
	return detail, nil
}
