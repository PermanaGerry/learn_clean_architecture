package route

import (
	"github.com/gofiber/fiber/v2"
	"learn-hexagonal-architecture/internal/delivery/http"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *http.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("api/users", c.UserController.Register)
	c.App.Post("api/users/login", c.UserController.Login)
}
