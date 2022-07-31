package main

import (
	"log"

	"github.com/AlfarrelIndrawan/cmlabs-onboarding/backend/config"
	"github.com/AlfarrelIndrawan/cmlabs-onboarding/backend/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setupRoutes(app *fiber.App) {
	app.Get("/blogs", handlers.GetBlogs)
	app.Get("/blogs/:id", handlers.GetBlog)
	app.Post("/blogs", handlers.CreateBlog)
	app.Put("/blogs/:id", handlers.UpdateBlog)
	app.Delete("/blogs/:id", handlers.RemoveBlog)
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	config.Connect()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
