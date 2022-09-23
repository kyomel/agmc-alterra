package helpers

import (
	"bookApp/internal/controllers"
	mid "bookApp/internal/middleware"
	"bookApp/internal/models"
	"bookApp/pkg/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HelperBook struct {
	BookControllers controllers.BookControllersInterface
}

type HelperBookInterface interface {
	CreateBook(c echo.Context) error
	GetAllBooks(c echo.Context) error
	GetBookById(c echo.Context) error
	UpdateBook(c echo.Context) error
	DeleteBook(c echo.Context) error
}

func InitBook(BookControllers controllers.BookControllersInterface) HelperBookInterface {
	return &HelperBook{BookControllers}
}

func (hb HelperBook) CreateBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)

	if err := c.Validate(book); err != nil {
		return err
	}

	response := new(utils.Response)
	result, err := hb.BookControllers.CreateBook(*book)

	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create book"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}

	return c.JSON(http.StatusCreated, response)
}

func (hb HelperBook) UpdateBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	extractToken := mid.ExtractTokenUserId(c)
	result, err := hb.BookControllers.UpdateBook(idInt, book)

	if err != nil || float64(idInt) != extractToken {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	} else {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to update book"
	}

	return c.JSON(http.StatusOK, response)
}

func (hb HelperBook) DeleteBook(c echo.Context) error {
	book := new(models.Book)
	c.Bind(book)
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	extractToken := mid.ExtractTokenUserId(c)
	result, err := hb.BookControllers.DeleteBook(idInt)

	if err != nil || float64(idInt) != extractToken {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	} else {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to delete book"
	}

	return c.JSON(http.StatusOK, response)
}

func (hb HelperBook) GetBookById(c echo.Context) error {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	response := new(utils.Response)
	result, err := hb.BookControllers.GetBookById(idInt)

	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Book not found"
	} else {
		response.Code = result.Code
		response.Status = result.Status
		response.Message = result.Message
		response.Result = result.Result
	}

	return c.JSON(http.StatusOK, response)
}

func (hb HelperBook) GetAllBooks(c echo.Context) error {
	response := new(utils.Response)
	result, err := hb.BookControllers.GetAllBooks()

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
