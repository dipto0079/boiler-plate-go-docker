package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func ListFacts(c *fiber.Ctx) error {
	facts := "all_ok"
	
	return c.Status(200).JSON(facts)
}
