package routes

import (
	"github.com/bluestoneag/log4j-collector/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/reports", controllers.GetReports)
	v1.Post("/reports", controllers.PostReport)
}
