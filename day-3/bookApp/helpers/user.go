package helpers

import (
	"bookApp/controllers"
	mid "bookApp/middleware"
	"bookApp/models"
	"bookApp/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HelperUser struct {
	UserControllers controllers.UserControllersInterface
}

type HelperUserInterface interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetUserById(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	UserLogin(c echo.Context) error
}

func Init(UserControllers controllers.UserControllersInterface) HelperUserInterface {
	return &HelperUser{UserControllers}
}

func (h HelperUser) CreateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)

	if err := c.Validate(user); err != nil {
		return err
	}

	response := new(utils.Response)
	result, err := h.UserControllers.CreateUser(*user)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create user"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}
	return c.JSON(http.StatusCreated, response)
}

func (h HelperUser) UpdateUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	extractToken := mid.ExtractTokenUserId(c)
	result, err := h.UserControllers.UpdateUser(idInt, user)

	if err != nil || float64(idInt) != extractToken {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to update user"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}
	return c.JSON(http.StatusOK, response)
}

func (h HelperUser) DeleteUser(c echo.Context) error {
	user := new(models.User)
	c.Bind(user)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	extractToken := mid.ExtractTokenUserId(c)
	result, err := h.UserControllers.DeleteUser(idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
	} else {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to delete user"
	}
	return c.JSON(http.StatusOK, response)
}

func (h HelperUser) GetUserById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	result, err := h.UserControllers.GetUserById(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "User not found"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}
	return c.JSON(http.StatusOK, response)
}

func (h HelperUser) GetAllUsers(c echo.Context) error {
	response := new(utils.Response)
	result, err := h.UserControllers.GetAllUser()

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Users not found"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}
	return c.JSON(http.StatusOK, response)
}

func (h HelperUser) UserLogin(c echo.Context) error {
	response := new(utils.Response)
	user := new(models.User)
	c.Bind(user)
	result, err := h.UserControllers.Login(user.Email, user.Password)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Login failed"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}
	return c.JSON(http.StatusOK, response)
}
