package helpers

import (
	"bookApp/internal/controllers"
	"bookApp/internal/models"
	"bookApp/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HelperReview struct {
	ReviewControllers controllers.ReviewControllersInterface
}

type HelperReviewInterface interface {
	CreateReview(c echo.Context) error
	GetReviews(c echo.Context) error
	GetReview(c echo.Context) error
	UpdateReview(c echo.Context) error
	DeleteReview(c echo.Context) error
}

func InitReview(ReviewControllers controllers.ReviewControllersInterface) HelperReviewInterface {
	return &HelperReview{ReviewControllers}
}

func (hr *HelperReview) CreateReview(c echo.Context) error {
	review := new(models.Review)
	response := new(utils.Response)
	c.Bind(review)

	if err := c.Validate(review); err != nil {
		return err
	}

	result, err := hr.ReviewControllers.CreateReview(*review)
	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to create review"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Result = result.Result
	return c.JSON(http.StatusCreated, response)

}

func (hr *HelperReview) GetReviews(c echo.Context) error {
	response := new(utils.Response)

	result, err := hr.ReviewControllers.GetReviews()
	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to get reviews"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Result = result.Result
	return c.JSON(http.StatusCreated, response)
}

func (hr *HelperReview) GetReview(c echo.Context) error {
	response := new(utils.Response)
	id := c.Param("id")
	result, err := hr.ReviewControllers.GetReview(id)
	if err != nil {
		response.Code = 404
		response.Status = "failed"
		response.Message = "Failed to get review"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Result = result.Result
	return c.JSON(http.StatusCreated, response)
}

func (hr *HelperReview) UpdateReview(c echo.Context) error {
	review := new(models.Review)
	response := new(utils.Response)
	c.Bind(review)
	id := c.Param("id")
	if err := c.Validate(review); err != nil {
		return err
	}

	result, err := hr.ReviewControllers.UpdateReview(review, id)
	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to update review"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Result = result.Result
	return c.JSON(http.StatusCreated, response)
}

func (hr *HelperReview) DeleteReview(c echo.Context) error {
	response := new(utils.Response)
	id := c.Param("id")
	result, err := hr.ReviewControllers.DeleteReview(id)
	if err != nil {
		response.Code = 400
		response.Status = "failed"
		response.Message = "Failed to delete review"
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Code = result.Code
	response.Status = result.Status
	response.Message = result.Message
	response.Result = result.Result
	return c.JSON(http.StatusCreated, response)
}
