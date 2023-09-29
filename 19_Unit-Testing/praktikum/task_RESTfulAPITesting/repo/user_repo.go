package repo

import (
	"errors"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/helpers"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/middleware"
	"go_septiandi-nugraha/19_Unit-Testing/praktikum/task_RESTfulAPITesting/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (userRepo *UserRepository) CheckLogin(email string, password string) (models.User, string, error) {
	var data models.User

	// Get email from user
	useremail := userRepo.db.Where("email = ?", email).First(&data)
	if useremail.Error != nil {
		return models.User{}, "", useremail.Error
	}

	// Password verify
	_, err := helpers.VerifyPassword(data.Password, password)
	if err != nil {
		return models.User{}, "", err
	}

	var token string
	if useremail.RowsAffected > 0 {
		var errToken error
		token, errToken = middleware.CreateToken(int(data.ID), data.Name)
		if errToken != nil {
			return models.User{}, "", errToken
		}
	}
	return data, token, nil
}

func (userRepo *UserRepository) SelectUsers() ([]models.User, error) {
	var users []models.User

	useremail := userRepo.db.Order("user_id desc").Find(&users)
	if useremail.Error != nil {
		return nil, useremail.Error
	}
	return users, nil
}

func (userRepo *UserRepository) SelectUserById(Id uint) (models.User, error) {
	var user models.User
	useremail := userRepo.db.First(&user, Id)
	if useremail.Error != nil {
		return models.User{}, useremail.Error
	}

	return user, nil
}

func (userRepo *UserRepository) InsertUser(user models.User) error {
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		return errors.New("Hash password failed")
	}
	user.Password = hash

	useremail := userRepo.db.Create(&user)
	if useremail.Error != nil {
		return useremail.Error
	}
	return nil
}

func (userRepo *UserRepository) DeleteUser(Id uint) error {
	useremail := userRepo.db.Delete(&models.User{}, Id)
	if useremail.Error != nil {
		return errors.New("Delete user failed")
	}

	if useremail.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}

func (userRepo *UserRepository) UpdateUser(user models.User) error {
	updatedUser := models.User{
		Name:  user.Name,
		Email: user.Email,
	}

	useremail := userRepo.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(updatedUser)
	if useremail.Error != nil {
		return useremail.Error
	}

	if useremail.RowsAffected == 0 {
		return errors.New("User not found")
	}

	return nil
}
