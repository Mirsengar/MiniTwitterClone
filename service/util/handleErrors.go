package util

import (
          "github.com/gofiber/fiber/v2"
)

var DefaultErrorHandler = func(c *fiber.Ctx, err error) error {
          code := fiber.StatusInternalServerError
          if e, ok := err.(*fiber.Error); ok {
                    code = e.Code
          }
          c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
          return c.Status(code).JSON(&fiber.Map{
                    "success": false,
                    "err":     err.Error(),
          })
}
