package server

import (
	"net/http"

	"github.com/adityaladwa/todoapp/db/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupServer() {
	s := fiber.New()
	s.Use(logger.New())
	s.Get("api/v1/users", getUsers)
	s.Post("api/v1/users", addUser)
	s.Listen(":8080")
}

func getUsers(c *fiber.Ctx) error {
	userResult := users.GetUsers()
	c.JSON(userResult)
	return nil
}

func addUser(c *fiber.Ctx) error {
	user := new(users.User)
	if err := c.BodyParser(user); err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	if err := user.InsertUser(); err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	user.Password = ""
	c.Status(http.StatusCreated).JSON(user)
	return nil
}
