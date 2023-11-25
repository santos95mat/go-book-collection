package handler

import "github.com/gofiber/fiber/v2"

type statusHandler struct{}

func NewStatusHandler() *statusHandler {
	return &statusHandler{}
}

func (*statusHandler) Get(c *fiber.Ctx) error {
	return c.SendString("Book-collection API is running")
}
