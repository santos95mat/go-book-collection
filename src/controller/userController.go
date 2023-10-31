package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/src/dto"
	"github.com/santos95mat/go-book-collection/src/service"
	"github.com/santos95mat/go-book-collection/src/util"
)

type UserController struct {
	userService service.UserService
	userBody    dto.UserBody
}

func (b UserController) Create(c *fiber.Ctx) error {
	err := c.BodyParser(&b.userBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	data, err := util.ValidUser(b.userBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	user, err := b.userService.Create(data)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (b UserController) GetMany(c *fiber.Ctx) error {
	q := c.Queries()
	search := q["search"]

	users, err := b.userService.GetMany(search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (b UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := b.userService.GetOne(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (b UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	err := c.BodyParser(&b.userBody)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	user, err := b.userService.Update(id, b.userBody)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (b UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := b.userService.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deletado com sucesso",
	})
}
