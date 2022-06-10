package controllers

import (
	"goFiberRestApi/models"

	"github.com/gofiber/fiber/v2"
)

func AuthLoginUser(c *fiber.Ctx) error{
	var user models.User
	var err error

	if err := c.BodyParser(&user); err != nil{
		return gernerateJsonError("invalid", c,fiber.StatusInternalServerError)
	}
	// check if username and password is not null
	if len(user.Password) < 6 || len(user.Email) < 6{
		return gernerateJsonError("unable to create user", c, fiber.StatusNotAcceptable) 
	}
	if user, err = models.LoginUser(&user); err != nil{
		return gernerateJsonError("unable to create user", c, fiber.StatusNotAcceptable)
	}

	// if err := models.LoginUser(&user); err != nil{
	// 	return gernerateJsonError("unable to create user", c, fiber.StatusNotAcceptable)
	// }
	return c.Status(fiber.StatusOK).JSON(
		fiber.Map{
			"message": "user logged in",
			"user": user,
		},
	)
	// return c.SendStatus(fiber.StatusAccepted)
}

func AuthSignUpUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil{
		return gernerateJsonError("invalid",c,fiber.StatusInternalServerError)
	}
	if err := models.CreateUser(user); err != nil{
		return gernerateJsonError("unable to create user", c, fiber.StatusNotAcceptable)
	}
	return c.Status(fiber.StatusCreated).JSON(
		fiber.Map{
			"Ok": true,
		},
	)
}