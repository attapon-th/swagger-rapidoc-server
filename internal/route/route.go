// Package route api service
package route

import (
	"github.com/attapon-th/swagger-rapidoc/internal/controller"
	"github.com/attapon-th/swagger-rapidoc/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// New api router
func New(app fiber.Router) {
	pathPrefix := viper.GetString("app.prefix")

	// Setup Middleware
	r := app.Use(pathPrefix, middleware.LogAccess(), middleware.Cache(), middleware.Compress())

	// Create Group route
	routePublic(r.Group(pathPrefix))

}

func routePublic(rt fiber.Router) {
	//  app public route handler
	// import controller public
	controller.EndpointPing(rt)
	controller.EndpointSwagger(rt)
}
