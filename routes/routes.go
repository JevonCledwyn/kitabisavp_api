package routes

import (
	"net/http"

	"kitabisavp/controllers"
	"github.com/labstack/echo/v4"
)

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	name := c.Param("name")
	return c.String(http.StatusOK, "Hello, "+name+"!")
}

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})



	
	return e
}
