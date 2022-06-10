package controllers

import "github.com/gofiber/fiber/v2"

func gernerateJsonError(message string, c *fiber.Ctx, httpStatus int) error{
	return c.Status(httpStatus).JSON(
		fiber.Map{
			"message": message,
		},
	)
}