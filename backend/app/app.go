package app

import (
	"log"

	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/adityaladwa/todoapp/server"
	"github.com/gofiber/fiber/v2"
)

func Init() error {
	db.Connect()
	setupRoutes()
	return nil
}

func setupRoutes() {
	log.Printf("Setting up routes")
	s := fiber.New()
	s.Get("/users", server.GetUsers)
	s.Post("/users", server.AddUser)
	log.Printf("Server is listening on 8080")
	s.Listen(":8080")
}
