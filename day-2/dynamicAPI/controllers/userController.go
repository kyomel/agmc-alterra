package controllers

import (
	"dynamicAPI/lib/database"
	"dynamicAPI/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserControllers struct {
	Lib database.UserRepo
}

type UserControllersInterface interface {
	CreateUserControllers(echo.Context) error
	GetUsersControllers(echo.Context) error
	GetUserByIDControllers(echo.Context) error
	UpdateUserControllers(echo.Context) error
	DeleteUserControllers(echo.Context) error
}

func Init(Lib database.UserRepo) UserControllersInterface {
	return &UserControllers{Lib}
}

func (uc UserControllers) CreateUserControllers(c echo.Context) error {
	var data models.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	uc.Lib.CreateUser(data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": data,
	})
}

func (uc UserControllers) GetUsersControllers(c echo.Context) error {
	data, err := uc.Lib.GetUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": data,
	})
}

func (uc UserControllers) GetUserByIDControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	data, err := uc.Lib.GetUserByID(idNumber)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": data,
	})
}

func (uc UserControllers) UpdateUserControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)
	var data models.User

	if err := c.Bind(&data); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	uc.Lib.UpdateUser(idNumber, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": data,
	})
}

func (uc UserControllers) DeleteUserControllers(c echo.Context) error {
	id := c.Param("id")
	idNumber, _ := strconv.Atoi(id)

	_, err := uc.Lib.DeleteUser(idNumber)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Data User!",
	})
}
