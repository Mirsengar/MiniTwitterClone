package util

import (
          `github.com/gofiber/fiber/v2`
)

func ResponseHelperJSON(c *fiber.Ctx, data any, dataType string, dataError string) {
          if data != nil {
                    c.Status(200).JSON(&fiber.Map{"success": true, dataType: data})
          } else {
                    c.Status(404).JSON(&fiber.Map{"success": false, "error": dataError})
          }
}
