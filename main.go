package main

import (
	"fmt"
	"log"
	"os"

	"github.com/faizan-glitch/vox-populi/database"
	"github.com/faizan-glitch/vox-populi/models"
	"github.com/faizan-glitch/vox-populi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	dsn := fmt.Sprintf("host=database user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Karachi", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	database.DBConnection, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	database.DBConnection.AutoMigrate(&models.Poll{})
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
		log.Fatal("Could not load .env file")
	}

	initDatabase()

	app := fiber.New(fiber.Config{
		Prefork: true,
	})

	app.Use(limiter.New(limiter.Config{
		Max: 15,
	}))

	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	routes.Register(app)

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))

	app.Listen(port)
}
