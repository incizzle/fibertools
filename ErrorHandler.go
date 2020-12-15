package fibertools

import (
	"github.com/gofiber/fiber/v2"
)

// Error handler that is passed into `fiber.New()`. Set x-debug in request header to see stack trace of error.
func ErrorHandler(c *fiber.Ctx, err error) error {
	isDebug := GetHeader(c, "x-debug")
	if isDebug == "true" {
		richErr, ok := err.(*RichError)
		if !ok {
			richErr = NewError(err)
		}
		return c.Status(richErr.Code).JSON(fiber.Map{
			"status": richErr.Message,
			"code":   richErr.Code,
			"error":  richErr.StackTrace(),
		})
	}

	fiberErr, ok := err.(*fiber.Error)
	if ok {
		return c.Status(fiberErr.Code).JSON(fiber.Map{
			"status": fiberErr.Message,
		})
	}

	return c.Status(500).JSON(fiber.Map{
		"status": err.Error(),
	})
}
