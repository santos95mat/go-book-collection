package controller

import "github.com/gofiber/fiber/v2"

type StatusController struct{}

func (StatusController) Get(c *fiber.Ctx) error {
	return c.SendString("Book-collection API is running")
}
