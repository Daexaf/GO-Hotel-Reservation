package api

import (
	"github.com/daexaf/GO-Hotel-Reservation/models"

	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error {
	u := models.User{
		FirstName: "James",
		LastName:  "Bond",
	}
	return c.JSON(u)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("James")
}
