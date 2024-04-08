package router

import (
	"github.com/emmaperez2197/go-fiber/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetuoRoutes(app *fiber.App) {

	app.Post("api/user", handlers.CreateUser)
	app.Post("api/user/log-in", handlers.LogIn)
}
