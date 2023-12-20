package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".tpl")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/stream", StreamVideo)

	app.Listen(":9000")
}

func StreamVideo(c *fiber.Ctx) error {
	return c.SendFile("video.mp4", false)
}
