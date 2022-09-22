package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Comment string             `json:"comment" form:"comment"`
	Like    bool               `json:"like" form:"like" validate:"like"`
}
