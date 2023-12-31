package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/santos95mat/go-book-collection/initializer/database"
	"github.com/santos95mat/go-book-collection/internal/entity"
)

func Auth(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, _ := jwt.Parse(tokenStr, keyFunc)
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var user entity.User
		err := database.DB.First(&user, "id = ?", claims["sub"]).Error

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("user", user)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
}

func AuthADM(c *fiber.Ctx) error {
	tokenStr := c.Cookies("Authorization")
	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	token, _ := jwt.Parse(tokenStr, keyFunc)
	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		var user entity.User
		err := database.DB.First(&user, "id = ?", claims["sub"]).Error

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		if user.Role != "admin" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}
		c.Locals("user", user)
		return c.Next()
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
}

func keyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, nil
	}

	return []byte(os.Getenv("SECRET")), nil
}
