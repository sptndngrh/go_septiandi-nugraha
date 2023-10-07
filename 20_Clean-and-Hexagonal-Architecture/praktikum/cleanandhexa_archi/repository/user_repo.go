package repository

import (
	"cleanandhexa_archi/model"

	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user model.User) error
	Find() ([]model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (u *userRepo) Create(user model.User) error {
	return u.db.Save(&user).Error
}

func (u *userRepo) Find() ([]model.User, error) {
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}
