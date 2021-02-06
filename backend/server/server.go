package server

import (
	"net/http"

	"github.com/adityaladwa/todoapp/db/users"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	userResult := users.GetUsers()
	c.JSON(userResult)
	return nil
}

func AddUser(c *fiber.Ctx) error {
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
