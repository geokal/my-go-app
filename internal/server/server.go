package server

import (
	"github.com/gofiber/fiber/v2"

	"my-go-app/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "my-go-app",
			AppName:      "my-go-app",
		}),

		db: database.New(),
	}

	return server
}
