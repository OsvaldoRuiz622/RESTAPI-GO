package main

import (
	"github.com/OsvaldoRuiz622/APIREST_GO/handlers"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/sabores", handlers.ListarSabores)

	app.Get("/sabores/:id", handlers.GetSabor)

	app.Post("/sabores", handlers.Create)

	app.Delete("/sabores/:id", handlers.DeleteSabor)

	app.Put("/sabores/:id", handlers.PutSabor)

}
