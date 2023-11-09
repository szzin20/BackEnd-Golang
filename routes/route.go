package routes

import (
	"healthcare/middlewares"
	"github.com/labstack/echo/v4"
)

func SetupRoutes() *echo.Echo {

	e := echo.New()

	middlewares.RemoveTrailingSlash(e)
	middlewares.Logger(e)
	middlewares.RateLimiter(e)
	middlewares.Recover(e)
	middlewares.CORS(e)
	
	return e

}
