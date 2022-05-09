package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/abdullahb53/beyazhoroz/api"
	"github.com/abdullahb53/beyazhoroz/responses"
	"google.golang.org/grpc"

	"github.com/abdullahb53/beyazhoroz/configs"
	"github.com/abdullahb53/beyazhoroz/controllers"

	_ "github.com/abdullahb53/beyazhoroz/controllers"
	_ "github.com/abdullahb53/beyazhoroz/models"
	_ "github.com/abdullahb53/beyazhoroz/responses"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

//func New(config Config) fiber.Handler

type myFurkanService struct {
	api.UnimplementedFurkanServiceServer
}

func main() {

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "beyazhoroz",
		AppName:       "beyazhoroz-apiserver",
	})
	app.Use(cors.New())
	var grpc_port string = "localhost:443"
	lis, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	api.RegisterFurkanServiceServer(srv, &myFurkanService{})
	fmt.Println("starting gRPC server...", grpc_port)
	go func() {
		panic(srv.Serve(lis))
	}()

	func (c *myFurkanService) ListUsers(fiber.Ctx , req *api.UserListReq) (*api.UserListRes, error) {
		res := new(api.UserListRes)
		err := c.cc.Invoke(ctx, "/FurkanService/ListUsers", req, res, opts...)
		if err != nil {
			return nil, err
		}
		res
		return out, nil
	}


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
		city := c.Params("city")
		country := c.Params("country")

		fmt.Println(country, city)

		//c.JSON(data)
		return c.SendStatus(200)

	})

	app.Post("/api/user", func(c *fiber.Ctx) error {

		controllers.CreateUser(c)
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
