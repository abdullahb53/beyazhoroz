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
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "cards")
var locationCollection *mongo.Collection = configs.GetCollection(configs.DB, "locations")

var validate = validator.New()

func CreateLocation(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var location models.Location
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&location); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.LocationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&location); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.LocationResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	// Id      primitive.ObjectID `json:"id,omitempty"`
	// Name    string             `json:"name,omitempty" validate:"required"`
	// Phone   string             `json:"phone,omitempty" validate:"required"`
	// City    string             `json:"city,omitempty" validate:"required"`
	// Country string             `json:"country,omitempty" validate:"required"`
	// Explain string             `json:"explain,omitempty" validate:"required"`

	newLocation := models.Location{
		Id:          primitive.NewObjectID(),
		Name:        location.Name,
		Information: location.Information,
		Phone:       location.Phone,
		Enlem:       location.Enlem,
		Boylam:      location.Boylam,
		Type:        location.Type,
		Initdate:    time.Now().String(),
	}

	result, err := locationCollection.InsertOne(ctx, newLocation)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.LocationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.LocationResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAllLocations(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var locations []models.Location

	defer cancel()

	results, err := locationCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.LocationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleLocation models.Location
		if err = results.Decode(&singleLocation); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.LocationResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		locations = append(locations, singleLocation)
	}
	return c.Status(http.StatusOK).JSON(
		responses.LocationResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": locations}},
	)
}

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

func GetAllUsers(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var users []models.User

	city := c.Params("city")
	country := c.Params("country")

	defer cancel()

	results, err := userCollection.Find(ctx, bson.M{"city": city, "country": country})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleUser models.User
		if err = results.Decode(&singleUser); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		users = append(users, singleUser)
	}
	return c.Status(http.StatusOK).JSON(
		responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": users}},
	)
}
