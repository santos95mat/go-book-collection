package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/repository"
	"github.com/santos95mat/go-book-collection/internal/util"
)

type UserController struct {
	userRepository repository.UserRepository
}

func (b UserController) Create(c *fiber.Ctx) error {
	var createUserDTO dto.UserInputDTO
	err := c.BodyParser(&createUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	createUserDTO, err = util.ValidUser(createUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	user, err := b.userRepository.Create(createUserDTO)

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

	users, err := b.userRepository.GetMany(search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (b UserController) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := b.userRepository.GetOne(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (b UserController) Update(c *fiber.Ctx) error {
	var updateUserDTO dto.UserInputDTO
	id := c.Params("id")
	err := c.BodyParser(&updateUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	user, err := b.userRepository.Update(id, updateUserDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (b UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := b.userRepository.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deletado com sucesso",
	})
}
