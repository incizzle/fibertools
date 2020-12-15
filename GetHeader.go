package fibertools

import "github.com/gofiber/fiber/v2"

// Pass in fiber context and the specific header you want to pull from the context.
func GetHeader(c *fiber.Ctx, header string) string {
	return string(c.Request().Header.Peek(header))
}
