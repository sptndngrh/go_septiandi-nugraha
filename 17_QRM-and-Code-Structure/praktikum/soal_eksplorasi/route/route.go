package route

import (
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Untuk pengelolaan data user
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)
	// Untuk pengelolaan data buku
	e.GET("/books", controllers.GetAllBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	// Untuk pengelolaan data blog
	e.GET("/blogs", controllers.GetAllBlogController)
	e.GET("/blogs/:id", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogController)
	e.PUT("/blogs/:id", controllers.UpdateBlogController)
}
