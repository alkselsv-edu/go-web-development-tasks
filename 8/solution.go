package main

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type User struct {
	Username string
	Email    string
	Age      int
	Country  string
}

var users = map[string]User{}

type (
	CreateUserRequest struct {
		// BEGIN (write your solution here)
		// END
	}
)

var usernameRegexp = regexp.MustCompile(`^[a-z0-9]+$`)

func SetupApp() *fiber.App {
	webApp := fiber.New(fiber.Config{ReadBufferSize: 16 * 1024})
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("1334414")
	})

	// BEGIN (write your solution here)
	// END

	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
