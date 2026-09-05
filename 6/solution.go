package main

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	CreateLinkRequest struct {
		External string `json:"external"`
		Internal string `json:"internal"`
	}

	GetLinkResponse struct {
		Internal string `json:"internal"`
	}
)

var links = make(map[string]string)

func SetupApp() *fiber.App {
	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	webApp.Post("/links", func(c *fiber.Ctx) error {
		// BEGIN (write your solution here)
		// END
	})
	webApp.Get("/links/:external", func(c *fiber.Ctx) error {
		// BEGIN (write your solution here)
		// END
	})
	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
