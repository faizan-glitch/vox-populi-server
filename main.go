package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	PORT := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(PORT)
}
