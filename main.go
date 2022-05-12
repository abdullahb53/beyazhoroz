package main

import (
	"fmt"
	"log"
	"os"

	"github.com/abdullahb53/beyazhoroz/configs"
	"github.com/abdullahb53/beyazhoroz/controllers"
	_ "github.com/abdullahb53/beyazhoroz/models"
	_ "github.com/abdullahb53/beyazhoroz/responses"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

//func New(config Config) fiber.Handler

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "beyazhoroz",
		AppName:       "beyazhoroz-apiserver",
	})
	app.Use(cors.New())
	// },
	configs.ConnectDB()

	// Provide a minimal config
	app.Use(favicon.New())

	// Or extend your config for customization
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
	}))

	// app.Get("/index.html", func(c *fiber.Ctx) error {

	// 	return c.SendFile("././client/index.html")
	// })
	app.Static("/index", "client/")
	app.Static("/index.html", "client/")
	app.Static("/", "client/")

	// Parameters
	app.Get("/api/:city/:country", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("city"))
		fmt.Fprintf(c, "%s\n", c.Params("country"))
		// city := c.Params("city")
		// country := c.Params("country")
		controllers.GetAllUsers(c)

		//c.JSON(data)
		return c.SendStatus(200)

	})

	app.Post("/api/CREATEUSER", func(c *fiber.Ctx) error {

		controllers.CreateUser(c)
		return nil

	})

	app.Get("/api/LOCATIONS", func(c *fiber.Ctx) error {
		controllers.GetAllLocations(c)
		return c.SendStatus(200)
	})

	app.Post("/api/CREATELOCATION", func(c *fiber.Ctx) error {
		controllers.CreateLocation(c)
		return nil
	})

	app.Get("/api/CREATELOCATION", func(c *fiber.Ctx) error {
		controllers.CreateLocation(c)
		return nil
	})

	// Get the PORT from heroku env
	port := os.Getenv("PORT")

	// Verify if heroku provided the port or not
	if os.Getenv("PORT") == "" {
		port = "3000"
	}

	// Start server on http://${heroku-url}:${port}
	log.Fatal(app.Listen(":" + port))
	// app.Listen(":8080")
}
