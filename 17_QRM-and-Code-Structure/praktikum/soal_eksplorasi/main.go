package main

import (
	"fmt"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/config"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/models"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/route"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	InitDB()
	InitialMigration()
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

func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
	DB.AutoMigrate(&models.Blog{})

}

func main() {
	e := echo.New()
	route.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
