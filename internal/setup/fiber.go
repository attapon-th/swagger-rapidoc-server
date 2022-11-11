package setup

import (
	"fmt"
	"runtime"

	"github.com/attapon-th/go-pkgs/zlog/log"

	json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// NewFiberConfig fiber server initailize
func NewFiberConfig() fiber.Config {
	config := fiber.Config{
		AppName:       "fiber-api",
		Prefork:       false,
		StrictRouting: false,
		CaseSensitive: true,
		Immutable:     false,
		BodyLimit:     4 * 1024 * 1024, // 4mb body size
		Concurrency:   256 * 1024,
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	}

	if nc := viper.GetInt("app.cpu"); nc > 1 {
		config.Prefork = true
		runtime.GOMAXPROCS(nc)
	} else {
		config.Prefork = false
	}
	return config
}

// Listen start server api
func Listen(app *fiber.App) error {
	l := fmt.Sprintf("%s:%s", viper.GetString("app.listen"), viper.GetString("app.port"))
	log.Info().Str("Listener", l).Msg("API server started")
	return app.Listen(l)
}

func errorHandlerResponseJSON(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	msgError := "Error!!!"
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	if err != nil {
		msgError = err.Error()
	}
	resError := fiber.Map{
		"msg": msgError,
		"ok":  false,
	}
	return ctx.Status(code).JSON(resError)
}
