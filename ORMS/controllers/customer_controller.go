package controllers

import (
	"github.com/gofiber/fiber/v2"

	"bun-spreader/dto"
	"bun-spreader/services"
	"bun-spreader/utils"
)

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) InsertUser(c *fiber.Ctx) error {
	var body dto.Customer
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	_, err := uc.UserService.CreateUser(body)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"response": "Customer Has Been Created Successfully",
	})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	getAllUsers, err := uc.UserService.GetAllUsers()
	utils.HandleError(err)
	return c.Status(fiber.StatusCreated).JSON(getAllUsers)
}

func (uc *UserController) GetUserById(c *fiber.Ctx) error {
	userId := c.Params("userId")
	getUser, err := uc.UserService.GetUserByID(userId)
	utils.HandleError(err)
	return c.Status(fiber.StatusCreated).JSON(getUser)
}

func (uc *UserController) GetUserByName(c *fiber.Ctx) error {
	userName := c.Query("name")
	getUser, err := uc.UserService.GetUserByID(userName)
	utils.HandleError(err)
	return c.Status(fiber.StatusCreated).JSON(getUser)
}
