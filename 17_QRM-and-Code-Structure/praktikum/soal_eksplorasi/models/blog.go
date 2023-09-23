package models

import (
	"fmt"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
}

func InitDB() {
	config := config.LoadConfig()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type Blog struct {
	gorm.Model
	UserID  uint   `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (bl *Blog) TableName() string {
	return "blog"
}

func (b *Blog) User() (*User, error) {
    if b.UserID == 0 {
        return nil, fmt.Errorf("Blog record has no associated user")
    }
    
    var user User
    if err := DB.First(&user, b.UserID).Error; err != nil {
        return nil, err
    }
    return &user, nil
}