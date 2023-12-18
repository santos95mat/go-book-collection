package handler

import (
	"encoding/json"
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

func (h *authHandler) Token(c *fiber.Ctx) error {
	var userCreateToken dto.UserTokenInputDTO
	err := c.BodyParser(&userCreateToken)

	if err != nil {
		panic(err)
	}

	id, err := h.authRepository.VerifyEmail(userCreateToken.Email)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	userCreateToken.UserID = id

	req := fiber.Post("http://localhost:3030/token")
	// to set JSON BODY
	req.JSON(userCreateToken)

	statusCode, data, errs := req.Bytes()

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	var userResponseToken dto.UserTokenOutputDTO
	jsonErr := json.Unmarshal(data, &userResponseToken)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return c.Status(statusCode).JSON(userResponseToken)
}

func (h *authHandler) RecoveryPassword(c *fiber.Ctx) error {
	var userValidateToken dto.UserRecoverPasswordDTO
	err := c.BodyParser(&userValidateToken)

	if err != nil {
		panic(err)
	}

	req := fiber.Post("http://localhost:3030/token/validate")
	// to set JSON BODY
	req.JSON(userValidateToken)

	statusCode, data, errs := req.Bytes()

	if len(errs) > 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": errs,
		})
	}

	type resValidate struct {
		Message string `json:"message"`
	}

	var isvalid resValidate
	jsonErr := json.Unmarshal(data, &isvalid)
	if jsonErr != nil {
		panic(jsonErr)
	}

	return c.Status(statusCode).JSON(isvalid)
}
