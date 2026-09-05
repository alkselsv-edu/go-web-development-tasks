package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

type (
	SendPushNotificationRequest struct {
		Message string `json:"message"`
		UserID  int64  `json:"user_id"`
	}

	PushNotification struct {
		Message string `json:"message"`
		UserID  int64  `json:"user_id"`
	}
)

var pushNotificationsQueue []PushNotification

func SetupApp() *fiber.App {
	// BEGIN (write your solution here)
	// END
	webApp.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})
	webApp.Post("/push/send", func(c *fiber.Ctx) error {
		var req SendPushNotificationRequest
		if err := c.BodyParser(&req); err != nil {
			// BEGIN (write your solution here)
			// END
		}

		pushNotificationsQueue = append(pushNotificationsQueue, PushNotification{
			Message: req.Message,
			UserID:  req.UserID,
		})
		if len(pushNotificationsQueue) > 3 {
			panic("Queue is full")
		}

		return c.SendStatus(fiber.StatusOK)
	})
	return webApp
}

func main() {
	logrus.Fatal(SetupApp().Listen(":3000"))
}
