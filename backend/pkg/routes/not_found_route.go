package routes

import "github.com/gofiber/fiber/v2"

func NotFoundRoute(c *fiber.Ctx) error {
    c.Status(fiber.StatusNotFound)
    response := &ApiResponse {
        Code: fiber.StatusNotFound,
        Msg: "endpoint is not found",
        Error: true,
        Data: nil,
    }
    return c.JSON(response)
}

