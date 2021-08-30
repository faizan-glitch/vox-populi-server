package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetPolls(c *fiber.Ctx) error {
	return c.SendString("Get Polls")
}

func GetPoll(c *fiber.Ctx) error {
	return c.SendString("Get Poll")
}

func CreatePoll(c *fiber.Ctx) error {
	return c.SendString("Create Poll")
}

func UpdatePoll(c *fiber.Ctx) error {
	return c.SendString("Update Poll")
}

func DeletePoll(c *fiber.Ctx) error {
	return c.SendString("Delete Poll")
}
