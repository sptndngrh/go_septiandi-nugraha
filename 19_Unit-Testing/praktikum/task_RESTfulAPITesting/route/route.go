package route

import (
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/controllers"
	m "go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/middleware"

	"github.com/labstack/echo/v4"
)

func SetupRoutes() *echo.Echo {
	e := echo.New()
	m.LogMiddleware(e)

	UserProtect := e.Group("/users")

	UserProtect.GET("", controllers.GetUsersController)
	UserProtect.GET("/:id", controllers.GetUserController)
	UserProtect.DELETE("/:id", controllers.DeleteUserController)
	UserProtect.PUT("/:id", controllers.UpdateUserController)
	e.POST("/users", controllers.CreateUserController)
	e.POST("/users/login", controllers.LoginUserController)

	BooksProtect := e.Group("/books", m.JWTMiddleware())

	BooksProtect.GET("", controllers.GetAllBooksController)
	BooksProtect.GET("/:id", controllers.GetBookController)
	BooksProtect.PUT("/:id", controllers.UpdateBookController)
	BooksProtect.DELETE("/:id", controllers.DeleteBookController)
	BooksProtect.POST("", controllers.CreateBookController)

	return e
}
