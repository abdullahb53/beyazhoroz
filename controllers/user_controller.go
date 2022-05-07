package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/abdullahb53/beyazhoroz/configs"
	"github.com/abdullahb53/beyazhoroz/models"
	"github.com/abdullahb53/beyazhoroz/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "cards")
var validate = validator.New()

func CreateUser(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Id      primitive.ObjectID `json:"id,omitempty"`
	// Name    string             `json:"name,omitempty" validate:"required"`
	// Phone   string             `json:"phone,omitempty" validate:"required"`
	// City    string             `json:"city,omitempty" validate:"required"`
	// Country string             `json:"country,omitempty" validate:"required"`
	// Explain string             `json:"explain,omitempty" validate:"required"`

	newUser := models.User{
		Id:      primitive.NewObjectID(),
		Name:    user.Name,
		Phone:   user.Phone,
		City:    user.City,
		Country: user.Country,
		Explain: user.Explain,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}
