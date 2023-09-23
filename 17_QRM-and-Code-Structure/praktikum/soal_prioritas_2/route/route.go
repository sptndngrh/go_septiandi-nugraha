package route

import (
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_prioritas_2/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Untuk pengelolaan data user
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	// Untuk pengelolaan data user
	e.GET("/books", controllers.GetAllBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
}
