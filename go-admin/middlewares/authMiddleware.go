package middlewares

import (
	"govue/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	_, err := util.ParseJwt(cookie)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"messages": "Unautenticated",
		})
	}

	return c.Next()
}
