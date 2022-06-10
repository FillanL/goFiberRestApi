package controllers

import (
	"goFiberRestApi/models"

	"github.com/gofiber/fiber/v2"
)


func UserGetUserDetails(c *fiber.Ctx) error{
	id, err :=  c.ParamsInt("id")
	if err != nil {
		return gernerateJsonError("something was missing", c, fiber.StatusInternalServerError)
	}
	
	user, err := models.GetUserById(id)
	if err != nil{
		return gernerateJsonError("something was missing", c, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"user": user,
		},
	)
}
