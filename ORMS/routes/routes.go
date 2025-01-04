package routes

import (
	"github.com/gofiber/fiber/v2"

	"bun-spreader/controllers"
	"bun-spreader/services"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("api/v1")

	userController := &controllers.UserController{
		UserService: &services.UserService{},
	}

	api.Post("/user", userController.InsertUser)
	api.Get("/getUsers", userController.GetAllUsers)
	api.Get("/getUser/:userId", userController.GetUserById)
	api.Get("/getUser", userController.GetUserByName)
}
