package main

import (
	"fmt"
	"log"
	"praktikum_26/handler"
	"praktikum_26/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load .env file. Err: %s", err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	ailaptopUsecase := usecase.NewAILaptopUsecase()
	AILaptopHandler := handler.NewAILaptopHandler(ailaptopUsecase)
	e.POST("/welcome", AILaptopHandler.AIRecommended)
	port := ":8000"
	fmt.Printf("Server is running on port %s\n", port)
	e.Start(port)
}
