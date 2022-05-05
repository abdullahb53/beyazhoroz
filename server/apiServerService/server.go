package main

import (
	"fmt"

	"github.com/abdullahb53/beyazhoroz/server/apiServerService/database"
	_ "github.com/abdullahb53/beyazhoroz/server/apiServerService/database"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

//func New(config Config) fiber.Handler

func main() {

	db := database.CreateDBEngine()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "beyazhoroz",
		AppName:       "beyazhoroz-apiserver",
	})

	// Provide a minimal config
	app.Use(favicon.New())

	// Or extend your config for customization
	app.Use(favicon.New(favicon.Config{
		File: "./favicon.ico",
	}))

	// app.Get("/index.html", func(c *fiber.Ctx) error {

	// 	return c.SendFile("././client/index.html")
	// })
	app.Static("/index", "../../client/")
	app.Static("/index.html", "../../client/")
	app.Static("/", "../../client/")

	// Parameters
	app.Get("/api/:city/:country", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("city"))
		fmt.Fprintf(c, "%s\n", c.Params("country"))
		city := c.Params("city")
		country := c.Params("country")

		fmt.Println(country, city)

		data, err := database.GetUseCiCo(db, city, country)
		if err != nil {
			return err
		}

		c.JSON(data)
		return c.SendStatus(200)

	})

	app.Post("/api/:city/:country", func(c *fiber.Ctx) error {
		fmt.Fprintf(c, "%s\n", c.Params("city"))
		fmt.Fprintf(c, "%s\n", c.Params("country"))
		city := c.Params("city")
		country := c.Params("country")

		fmt.Println(country, city)

		err := database.InsertUser(db, "ABDULLAH BIYIK", city, country, "none", "KTU CEC ADMIN")
		if err != nil {
			return err
		}

		return nil

	})

	app.Listen(":8080")
}
