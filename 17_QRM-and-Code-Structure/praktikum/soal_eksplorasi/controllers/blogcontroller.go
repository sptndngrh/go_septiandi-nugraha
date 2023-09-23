package controllers

import (
	"go_septiandi-nugraha/17_QRM-and-Code-Structure/praktikum/soal_eksplorasi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllBlogController retrieves all blog data
func GetAllBlogController(c echo.Context) error {
	var blogs []models.Blog
	if err := DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"blogs":   blogs,
	})
}

// GetBlogController retrieves blog data by ID
func GetBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get blog",
		"blog":    blog,
	})
}

// CreateBlogController creates a new blog entry
func CreateBlogController(c echo.Context) error {
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := DB.Create(blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new blog entry",
		"blog":    blog,
	})
}

// DeleteBlogController deletes a blog entry by ID
func DeleteBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	var blog models.Blog
	if err := DB.First(&blog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}
	if err := DB.Delete(&blog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete blog entry",
	})
}

// UpdateBlogController updates a blog entry by ID
func UpdateBlogController(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}
	blog := new(models.Blog)
	if err := c.Bind(blog); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var existingBlog models.Blog
	if err := DB.First(&existingBlog, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Blog not found")
	}
	existingBlog.Title = blog.Title
	existingBlog.Content = blog.Content
	// You can update other fields as needed
	if err := DB.Save(&existingBlog).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update blog entry",
		"blog":    existingBlog,
	})
}
