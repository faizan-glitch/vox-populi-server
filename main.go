package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faizan-glitch/vox-populi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not load .env file")
	}

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(limiter.New())

	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	routes.Register(app)

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(PORT)
}
