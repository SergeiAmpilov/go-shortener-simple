package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var (
	port = ":8080"
)

func main() {
	webApp := fiber.New()

	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("Web application started")
	})

	webApp.Post("/links", func(c *fiber.Ctx) error {
		return nil
	})

	webApp.Get("/links/:external", func(c *fiber.Ctx) error {
		return nil
	})

	logrus.Info("Starting web app on port " + port)
	logrus.Fatal(webApp.Listen(port))

}
