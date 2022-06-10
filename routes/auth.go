package routes

import (
	"goFiberRestApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupAuthRoutes(router fiber.Router){
	router.Post("/login", controllers.AuthLoginUser)
	router.Post("/signup", controllers.AuthSignUpUser)
}