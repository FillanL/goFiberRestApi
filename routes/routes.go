package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) *fiber.App{
	fmt.Println("Setting up routes")

	v1 := app.Group("/api/v1")
	v1.Use(func(c *fiber.Ctx)error{
		// add middleware code here
		return c.Next()
	})

	auth := v1.Group("/auth")
	setupAuthRoutes(auth)


	user := v1.Group("/user")
	user.Use(func (c *fiber.Ctx) error {
		// add only user middlewares here
		return c.Next()
	})
	setUpUserRoutes(user)


	v2 := app.Group("/api/v2")
	v2.Use(func (c *fiber.Ctx) error {
		// add only v2 middlewares here
		return c.Next()
	})

	return app
}