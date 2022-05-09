package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Phone   string             `json:"phone,omitempty" validate:"required"`
	City    string             `json:"city,omitempty" validate:"required"`
	Country string             `json:"country,omitempty" validate:"required"`
	Explain string             `json:"explain,omitempty" validate:"required"`
}

type Location struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Uid    string             `json:"name,omitempty" validate:"required"`
	Enlem  string             `json:"enlem,omitempty" validate:"required"`
	Boylam string             `json:"boylam,omitempty" validate:"required"`
}
