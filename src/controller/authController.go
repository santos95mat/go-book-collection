package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/dto"
)

func (b UserController) Login(c *fiber.Ctx) error {
	var data dto.UserLogin

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	user, token, err := b.userService.Login(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":  user,
		"token": token,
	})
}
