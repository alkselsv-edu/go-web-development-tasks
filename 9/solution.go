package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"
)

func SetupApp() *fiber.App {
	file, err := os.OpenFile(".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("error opening file: %v", err)
	}
	// defer file.Close() // Не закрываем файл здесь, чтобы не закрыть до завершения работы сервера

	webApp := fiber.New(fiber.Config{
		ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	// BEGIN (write your solution here) (write your solution here)
	// END

	webApp.Get("/foo", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	webApp.Get("/bar", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
