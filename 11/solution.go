package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/sirupsen/logrus"
)

type (
	CreateItemRequest struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}

	Item struct {
		Name  string `json:"name"`
		Price uint   `json:"price"`
	}
)

var (
	items []Item
)

func SetupApp() *fiber.App {
	viewsEngine := html.New("./templates", ".tmpl")
	webApp := fiber.New(fiber.Config{
		Views:          viewsEngine,
		ReadBufferSize: 16 * 1024,
	})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	// BEGIN (write your solution here)
	// END
	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
