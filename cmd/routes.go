package main

import (
	"github.com/gofiber/fiber/v2"
	"api/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/",handlers.ListFacts)
}
