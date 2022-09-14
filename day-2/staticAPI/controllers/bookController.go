package controllers

import (
	"net/http"
	"staticAPI/models"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var (
	books = map[int]*models.Book{}
	seq   = 1
)

func GetBookAllControllers(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func CreateBookControllers(c echo.Context) error {
	b := &models.Book{
		ID:        seq,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	c.Request().Header.Get("Content-type")

	if err := c.Bind(b); err != nil {
		return err
	}

	books[b.ID] = b
	seq++
	return c.JSON(http.StatusCreated, b)
}

func GetBooksByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, books[id])
}

func UpdateBooksByID(c echo.Context) error {
	b := new(models.Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	books[id].Title = b.Title
	return c.JSON(http.StatusOK, books[id])
}

func DeleteBooks(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(books, id)
	return c.NoContent(http.StatusNoContent)
}
