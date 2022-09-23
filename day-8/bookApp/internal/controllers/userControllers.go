package controllers

import (
	mid "bookApp/internal/middleware"
	"bookApp/internal/models"
	"bookApp/internal/repository"
	"bookApp/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserControllers struct {
	Lib repository.UserRepo
}

type UserControllersInterface interface {
	CreateUser(user models.User) (*utils.Response, error)
	GetAllUser() (*utils.Response, error)
	GetUserById(id int) (*utils.Response, error)
	UpdateUser(id int, user *models.User) (*utils.Response, error)
	DeleteUser(id int) (*utils.Response, error)
	Login(email string, password string) (*utils.Response, error)
}

func Init(Lib repository.UserRepo) UserControllersInterface {
	return &UserControllers{Lib}
}

func (uc UserControllers) CreateUser(user models.User) (*utils.Response, error) {
	hash, _ := utils.PasswordHash(user.Password)
	userData := &models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: hash,
	}

	dataUser, err := uc.Lib.CreateUser(*userData)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    201,
		Status:  "Success",
		Message: "Success to create new user",
		Result:  dataUser,
	}

	return data, err
}

func (uc UserControllers) GetAllUser() (*utils.Response, error) {
	var arrUsers []utils.UserData

	user, err := uc.Lib.GetAllUser()
	if err != nil {
		return nil, err
	}

	for _, v := range *user {
		dataUser := utils.UserData{
			ID:        int(v.ID),
			FullName:  v.FullName,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		}

		arrUsers = append(arrUsers, dataUser)
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success to get all user",
		Result:  arrUsers,
	}

	return data, nil
}

func (uc UserControllers) GetUserById(id int) (*utils.Response, error) {
	user, err := uc.Lib.GetUserById(id)
	if err != nil {
		return nil, err
	}

	dataUser := &utils.UserData{
		ID:        int(user.ID),
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success get data user by ID",
		Result:  dataUser,
	}

	return data, nil
}

func (uc UserControllers) UpdateUser(id int, user *models.User) (*utils.Response, error) {
	userData := &models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: user.Password,
	}

	err := uc.Lib.UpdateUser(id, userData)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    201,
		Status:  "Success",
		Message: "Success to update user",
		Result:  err,
	}

	return data, nil
}

func (uc UserControllers) DeleteUser(id int) (*utils.Response, error) {
	err := uc.Lib.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success data user delete",
		Result:  err,
	}

	return data, nil
}

func (uc UserControllers) Login(email string, password string) (*utils.Response, error) {
	user, err := uc.Lib.Login(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	userLogin := utils.LoginResponse{}
	userLogin.ID = int(user.ID)
	userLogin.Email = user.Email
	userLogin.FullName = user.FullName
	userLogin.Token, err = mid.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	data := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success Login",
		Result:  userLogin,
	}

	return data, nil
}
