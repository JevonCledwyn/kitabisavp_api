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

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is user!")
	})

	// e.GET("/user/:name", getUser)

	// login
	e.POST("/login", controllers.CheckLoginWorker)
	e.POST("/login", controllers.CheckLoginRecruiter)

	// CRUD RECRUITERS

	e.GET("/recruiter", controllers.FetchAllRecruiter)
	e.POST("/recruiter", controllers.StoreRecruiter)
	e.PATCH("/recruiter", controllers.UpdateRecruiter)
	e.DELETE("/recruiter", controllers.DeleteRecruiter)
	e.GET("/recruiter/:id", controllers.FetchRecruiterById)

	// e.GET("/user", controllers.FetchAllUser,middleware.IsAuthenticated)
	e.GET("/worker", controllers.FetchAllRecruiter)
	e.POST("/worker", controllers.StoreRecruiter)
	e.PATCH("/worker", controllers.UpdateRecruiter)
	e.DELETE("/worker", controllers.DeleteRecruiter)
	e.GET("/worker/:id", controllers.FetchRecruiterById)

	// post Recruiter
	e.GET("/pr", controllers.FetchAllPR)
	e.GET("/pr/:id", controllers.FetchPRById)
	e.POST("/pr", controllers.StorePR)
	e.PATCH("/pr", controllers.UpdatePR)
	e.DELETE("/pr", controllers.DeletePR)

	// post Worker
	e.GET("/pw", controllers.FetchAllPW)
	e.GET("/pw/:id", controllers.FetchPWById)
	e.POST("/pw", controllers.StorePW)
	e.PATCH("/pw", controllers.UpdatePW)
	e.DELETE("/pw", controllers.DeletePW)



	return e
}