package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/santos95mat/go-book-collection/internal/dto"
	"github.com/santos95mat/go-book-collection/internal/interfaces"
)

type authHandler struct {
	authRepository interfaces.AuthRepositoryInterface
}

func NewAuthHandler(rep interfaces.AuthRepositoryInterface) *authHandler {
	return &authHandler{
		authRepository: rep,
	}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var data dto.UserLoginDTO

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	user, token, err := h.authRepository.Login(data)

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

func (h *authHandler) RecoveryPassword(c *fiber.Ctx) error {
	var data dto.UserRecoveryPasswordDTO

	err := c.BodyParser(&data)

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	err = h.authRepository.VerifyEmail(data.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// aqui vamos criar um novo recovery password e enviarmos um email para validar a troca da senha
	// go mail.SendMailHTML("", []string{data.Email})

	return c.Status(200).JSON(fiber.Map{
		"Message": "OK",
	})
}
