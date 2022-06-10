package server

import (
	"goFiberRestApi/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetUpServer() *fiber.App{
	server := fiber.New()

	server.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Orgin, Content-Type, Accept",
	}))
	
	server.Use(limiter.New(limiter.Config{
		Max: 100,
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(&fiber.Map{
				"status":  "fail",
				"message": "You have requested too many in a single time-frame! Please wait another minute!",
			})
		},
	}))

	server.Use(compress.New())
	server.Use(logger.New())
	routes.SetupRoutes(server)

	return server
}