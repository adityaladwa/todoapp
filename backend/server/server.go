package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/adityaladwa/todoapp/cache"
	"github.com/adityaladwa/todoapp/users"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v2"
)

var jwtSigningKey = "supersecret"

func SetupServer() {
	app := fiber.New()
	app.Use(logger.New())
	app.Get("api/v1/users", authMiddleware(), getUsers)
	app.Post("api/v1/users/signup", signUp)
	app.Post("api/v1/users/login", login)
	app.Listen(":8080")
}

func authMiddleware() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSigningKey),
	})
}

func getUsers(c *fiber.Ctx) error {
	userResult := users.GetUsers()
	c.JSON(userResult)
	return nil
}

func signUp(c *fiber.Ctx) error {
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

func login(c *fiber.Ctx) error {
	userLogin := new(users.User)
	if err := c.BodyParser(userLogin); err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	password := userLogin.Password
	userLogin.FindOne()
	ok := users.IsPasswordValid(userLogin.Password, password)
	if !ok {
		c.Status(http.StatusUnprocessableEntity).Send([]byte("Password is not ok"))
		return nil
	}
	jwtToken, err := generateJWTToken(userLogin.Email)
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	ctx := context.TODO()
	err = cache.RedisClient.Set(
		ctx,
		userLogin.Email,
		jwtToken,
		time.Minute*30).Err()
	if err != nil {
		c.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
		return err
	}
	userResponse := new(users.UserLoginResponse)
	userResponse.Token = jwtToken
	c.JSON(userResponse)
	return nil
}

func generateJWTToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["authorized"] = true
	claims["expire"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(jwtSigningKey))
	if err != nil {
		log.Printf("Something went wrong: %v", err)
	}
	return tokenString, nil
}
