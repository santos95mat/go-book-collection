package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/interfaces"
	"github.com/santos95mat/go-book-collection/internal/util"
)

type userHandler struct {
	userRepository interfaces.UserRepositoryInterface
}

func NewUserHandler(rep interfaces.UserRepositoryInterface) *userHandler {
	return &userHandler{
		userRepository: rep,
	}
}

func (h *userHandler) Create(c *fiber.Ctx) error {
	var createUserDTO dto.UserInputDTO
	err := c.BodyParser(&createUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	createUserDTO, err = util.ValidUser(createUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	user, err := h.userRepository.Create(createUserDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) GetMany(c *fiber.Ctx) error {
	q := c.Queries()
	search := q["search"]

	users, err := h.userRepository.GetMany(search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *userHandler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.userRepository.GetOne(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) Update(c *fiber.Ctx) error {
	var updateUserDTO dto.UserInputDTO
	id := c.Params("id")
	err := c.BodyParser(&updateUserDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	user, err := h.userRepository.Update(id, updateUserDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *userHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.userRepository.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deletado com sucesso",
	})
}
