package fibertools

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Simple fiber recover middleware based around default fiber recover but with the addiction of stack tarcing. Should be used with fibertools.ErrorHandler().
func Recover() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = NewError(fmt.Errorf("%v", r))
			}
		}()

		return c.Next()
	}
}
