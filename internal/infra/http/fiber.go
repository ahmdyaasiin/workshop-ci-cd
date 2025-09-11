package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		ErrorHandler:      errorHandler,
	})
}

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(fiber.Map{
		"error": err.Error(),
	})
}
