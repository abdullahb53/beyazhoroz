package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Name    string             `json:"name,omitempty" validate:"required"`
	Phone   string             `json:"phone,omitempty" validate:"required"`
	City    string             `json:"city,omitempty" validate:"required"`
	Country string             `json:"country,omitempty" validate:"required"`
	Explain string             `json:"explain,omitempty" validate:"required"`
}

type Location struct {
	Id          primitive.ObjectID `json:"id,omitempty"`
	Name        string             `json:"name,omitempty" validate:"required"`
	Information string             `json:"information,omitempty" validate:"required"`
	Phone       string             `json:"phone,omitempty" validate:"required"`
	Enlem       string             `json:"enlem,omitempty" validate:"required"`
	Boylam      string             `json:"boylam,omitempty" validate:"required"`
	Type        string             `json:"type,omitempty" validate:"required"`
	Initdate    string             `json:"initdate,omitempty" validate:"required"`
}

type Organik struct {
	Id             primitive.ObjectID `json:"id,omitempty"`
	Sehir          string             `json:"sehir,omitempty" validate:"required"`
	Urun           string             `json:"urun,omitempty" validate:"required"`
	Ciftci_sayisi  int32              `json:"ciftci_sayisi,omitempty" validate:"required"`
	Uretim_alani   int32              `json:"uretim_alani,omitempty" validate:"required"`
	Uretim_miktari int32              `json:"uretim_miktari,omitempty" validate:"required"`
}
