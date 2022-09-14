package database

import (
	"dynamicAPI/config"
	"dynamicAPI/models"

	"gorm.io/gorm"
)

type LibUser struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateUser(models.User) (*models.User, error)
	GetUser() (*[]models.User, error)
	GetUserByID(int) (*models.User, error)
	UpdateUser(int, models.User) (*models.User, error)
	DeleteUser(int) (*models.User, error)
}

func Init(DB *gorm.DB) UserRepo {
	return &LibUser{DB}
}

func (l *LibUser) CreateUser(user models.User) (*models.User, error) {
	if err := l.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (l *LibUser) GetUser() (data *[]models.User, err error) {
	if err := l.DB.Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (l *LibUser) GetUserByID(id int) (data *models.User, err error) {
	param := config.InitDB()

	if err = param.Where(`id=?`, id).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (l *LibUser) UpdateUser(id int, input models.User) (data *models.User, err error) {
	param := config.InitDB()

	if err = param.Where(`id=?`, id).Updates(input).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (l *LibUser) DeleteUser(id int) (data *models.User, err error) {
	param := config.InitDB()

	if err = param.Where(`id=?`, id).Delete(&data, id).Error; err != nil {
		return nil, err
	}

	return data, nil
}
