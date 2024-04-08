package main

import (
	"log"

	"github.com/emmaperez2197/go-fiber/database"
	"github.com/emmaperez2197/go-fiber/router"
	"github.com/gofiber/fiber/v2"
)

func main() {

	database.ConnnectDB()
	app := fiber.New()

	router.SetuoRoutes(app)
	log.Fatal(app.Listen(":8001"))
}
