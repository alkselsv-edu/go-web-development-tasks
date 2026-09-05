package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

var postLikes = map[string]int64{}

func SetupApp() *fiber.App {
	webApp := fiber.New(fiber.Config{
		Immutable:      true,
		ReadBufferSize: 16 * 1024,
	})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Go to /likes/12345")
	})
	// BEGIN (write your solution here)
	// END
	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
