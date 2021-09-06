package controllers

import (
	"fmt"

	"github.com/faizan-glitch/vox-populi/database"
	"github.com/faizan-glitch/vox-populi/models"
	"github.com/gofiber/fiber/v2"
)

func GetPolls(c *fiber.Ctx) error {
	return c.SendString("Get Polls")
}

func GetPoll(c *fiber.Ctx) error {
	return c.SendString("Get Poll")
}

func CreatePoll(c *fiber.Ctx) error {
	payload := struct {
		Question string `json:"question"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	poll := models.Poll{
		Question: payload.Question,
		Replies:  1,
	}

	result := database.DBConnection.Create(&poll)

	fmt.Println(result)

	return c.JSON(&fiber.Map{
		"success": true,
		"message": "poll created successfully",
	})
}

func UpdatePoll(c *fiber.Ctx) error {
	return c.SendString("Update Poll")
}

func DeletePoll(c *fiber.Ctx) error {
	return c.SendString("Delete Poll")
}
