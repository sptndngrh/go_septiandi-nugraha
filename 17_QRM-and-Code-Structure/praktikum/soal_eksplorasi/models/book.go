package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func (bl *Book) TableName() string {
	return "book"
}