package repository

import (
	"bookApp/internal/models"

	"gorm.io/gorm"
)

type LibUser struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(user models.User) (*models.User, error)
	GetAllUser() (*[]models.User, error)
	GetUserById(id int) (models.User, error)
	UpdateUser(id int, user *models.User) error
	DeleteUser(id int) error
	Login(email string) (models.User, error)
}

func Init(DB *gorm.DB) UserRepo {
	return &LibUser{DB}
}

func (u *LibUser) CreateUser(user models.User) (*models.User, error) {
	if err := u.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *LibUser) GetAllUser() (*[]models.User, error) {
	var (
		users *[]models.User
		err   error
	)

	if err = u.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *LibUser) GetUserById(id int) (models.User, error) {
	var user models.User

	res := u.DB.Where("id = ?", id).First(&user)
	return user, res.Error
}

func (u *LibUser) UpdateUser(id int, user *models.User) error {
	if err := u.DB.Model(user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}

	return nil
}

func (u *LibUser) DeleteUser(id int) error {
	var user *models.User

	if err := u.DB.Where("id = ?", id).Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}

func (u *LibUser) Login(email string) (models.User, error) {
	var user models.User

	result := u.DB.Where("email = ?", email).First(&user)
	return user, result.Error
}
