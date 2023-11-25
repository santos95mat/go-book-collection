package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/interfaces"
	"github.com/santos95mat/go-book-collection/internal/util"
)

type bookHandler struct {
	bookRepository interfaces.BookRepositoryInterface
}

func NewBookHandler(rep interfaces.BookRepositoryInterface) *bookHandler {
	return &bookHandler{
		bookRepository: rep,
	}
}

func (h *bookHandler) Create(c *fiber.Ctx) error {
	var createBookDTO dto.BookInputDTO
	err := c.BodyParser(&createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	createBookDTO, err = util.ValidBook(createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	book, err := h.bookRepository.Create(createBookDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *bookHandler) GetMany(c *fiber.Ctx) error {
	q := c.Queries()
	search := q["search"]

	books, err := h.bookRepository.GetMany(search)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(books)
}

func (h *bookHandler) GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	book, err := h.bookRepository.GetOne(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *bookHandler) Update(c *fiber.Ctx) error {
	var updateBookDTO dto.BookInputDTO
	id := c.Params("id")
	err := c.BodyParser(&updateBookDTO)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	book, err := h.bookRepository.Update(id, updateBookDTO)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(book)
}

func (h *bookHandler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.bookRepository.Delete(id)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Livro deletado com sucesso",
	})
}
