package routes

import (
	"goFiberRestApi/controllers"

	"github.com/gofiber/fiber/v2"
)

func setUpUserRoutes(router fiber.Router){
	router.Get("/:id", controllers.UserGetUserDetails)
	// router.Patch("/", controllers.UserCreateUser)
	// router.Delete("/:id", controllers.DeleteProject)
}
