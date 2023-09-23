package main

import (
	"fmt"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_prioritas_2/route"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_prioritas_2/config"
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_prioritas_2/models"

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
	configData := config.LoadConfig()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configData.DB_Username,
		configData.DB_Password,
		configData.DB_Host,
		configData.DB_Port,
		configData.DB_Name,
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
}

func main() {
	e := echo.New()
	route.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
