package main

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type User struct {
	Id int `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	UserID := c.Param("id")
	// Find a user by ID
	for _, user := range users {
		if strconv.Itoa(user.Id) == UserID {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Successfully get user by ID",
				"user":    user,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "User not found",
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	UserID := c.Param("id")
	// Find and delete a user by ID if found
	for i, user := range users {
		if strconv.Itoa(user.Id) == UserID {
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Successfully deleted user",
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "User not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	UserID := c.Param("id")
	// Find and update a user by ID if found
	for i, user := range users {
		if strconv.Itoa(user.Id) == UserID {
			// Bind the data to be updated
			updatedUser := User{}
			c.Bind(&updatedUser)
			// Update the found user
			users[i] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"message": "Successfully updated user",
				"user":    updatedUser,
			})
		}
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "User not found",
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)
	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user": user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()

	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users/add", CreateUserController)
	e.PUT("/users/update/:id", UpdateUserController)
	e.DELETE("/users/delete/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
