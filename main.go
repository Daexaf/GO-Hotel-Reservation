package main

import (
	"flag"

	"github.com/daexaf/GO-Hotel-Reservation/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listen", ":5000", "The listen address of the API server")

	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}
