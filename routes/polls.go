package routes

import (
	"github.com/faizan-glitch/vox-populi/controllers"
	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {

	app.Get("/polls", controllers.GetPolls)

	app.Get("/polls/:id", controllers.GetPoll)

	app.Post("/polls", controllers.CreatePoll)

	app.Patch("/polls/:id", controllers.UpdatePoll)

	app.Delete("/polls/:id", controllers.DeletePoll)

}
