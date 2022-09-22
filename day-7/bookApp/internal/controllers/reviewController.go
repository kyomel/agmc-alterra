package controllers

import (
	"bookApp/internal/models"
	"bookApp/internal/repository"
	"bookApp/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReviewControllers struct {
	LibReview repository.ReviewRepo
}

type ReviewControllersInterface interface {
	CreateReview(review models.Review) (*utils.Response, error)
	GetReviews() (*utils.Response, error)
	GetReview(id string) (*utils.Response, error)
	DeleteReview(id string) (*utils.Response, error)
	UpdateReview(review *models.Review, id string) (*utils.Response, error)
}

func InitReview(LibReview repository.ReviewRepo) ReviewControllersInterface {
	return &ReviewControllers{LibReview}
}

func (rc ReviewControllers) CreateReview(review models.Review) (*utils.Response, error) {
	reviewMapping := models.Review{
		Id:      primitive.NewObjectID(),
		Comment: review.Comment,
		Like:    review.Like,
	}

	err := rc.LibReview.CreateReview(reviewMapping)
	if err != nil {
		return nil, err
	}

	result := &utils.Response{
		Code:    201,
		Status:  "Success",
		Message: "Success Create Review",
		Result:  err,
	}

	return result, nil
}

func (rc ReviewControllers) GetReviews() (*utils.Response, error) {
	reviews, err := rc.LibReview.GetReviews()
	if err != nil {
		return nil, err
	}

	result := &utils.Response{
		Code:    200,
		Status:  "Success",
		Message: "Success Get Reviews",
		Result:  reviews,
	}

	return result, nil
}

func (rc ReviewControllers) GetReview(id string) (*utils.Response, error) {
	review, err := rc.LibReview.GetReview(id)
	if err != nil {
		return nil, err
	}

	result := &utils.Response{
		Code:    200,
		Status:  "success",
		Message: "Success get review",
		Result:  review,
	}

	return result, nil
}

func (rc ReviewControllers) DeleteReview(id string) (*utils.Response, error) {
	err := rc.LibReview.DeleteReview(id)
	if err != nil {
		return nil, err
	}

	result := &utils.Response{
		Code:    200,
		Status:  "success",
		Message: "Success delete review",
		Result:  err,
	}

	return result, nil
}

func (rc ReviewControllers) UpdateReview(review *models.Review, id string) (*utils.Response, error) {
	reviewMapping := bson.M{"comment": review.Comment, "like": review.Like}

	err := rc.LibReview.UpdateReview(reviewMapping, id)
	if err != nil {
		return nil, err
	}

	result := &utils.Response{
		Code:    200,
		Status:  "success",
		Message: "Success update review",
		Result:  err,
	}

	return result, nil
}
