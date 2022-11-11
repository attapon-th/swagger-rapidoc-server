// Package controller api controller
package controller

import (
	"github.com/gofiber/fiber/v2"
)

type pingResponse struct {
	OK  bool
	Msg string
}

// EndpointPing ping endpoint
//
//	@param r fiber.Router
func EndpointPing(r fiber.Router) {
	// init endpoint
	// ...
	// ...

	r.Get("/ping", getPing)
	// app more routers
}

func getPing(c *fiber.Ctx) error {
	return c.JSON(pingResponse{
		OK:  true,
		Msg: "Ping successfully",
	})
}
