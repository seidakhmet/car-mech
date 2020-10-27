package person

import "github.com/gofiber/fiber"

// PersonRoutes is urls
func PersonRoutes(app *fiber.App) {
	app.Get("/api/v1/test", GetPerson)
}
