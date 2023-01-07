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
	e.DELETE("/recruiter:recruiter_id", controllers.DeleteRecruiter)
	e.GET("/recruiter/:recruiter_id", controllers.FetchRecruiterById)

	// e.GET("/user", controllers.FetchAllUser,middleware.IsAuthenticated)
	e.GET("/worker", controllers.FetchAllWorker)
	e.POST("/worker", controllers.StoreWorker)
	e.PATCH("/worker", controllers.UpdateWorker)
	e.DELETE("/worker:worker_id", controllers.DeleteWorker)
	e.GET("/worker/:worker_id", controllers.FetchWorkerById)

	// post Recruiter
	e.GET("/pr", controllers.FetchAllPR)
	e.GET("/pr/:post_recruiter_id", controllers.FetchPRById)
	e.POST("/pr", controllers.StorePR)
	e.PATCH("/pr", controllers.UpdatePR)
	e.DELETE("/pr:post_recruiter_id", controllers.DeletePR)

	// post Worker
	e.GET("/pw", controllers.FetchAllPW)
	e.GET("/pw/:post_worker_id", controllers.FetchPWById)
	e.POST("/pw", controllers.StorePW)
	e.PATCH("/pw", controllers.UpdatePW)
	e.DELETE("/pw:post_worker_id", controllers.DeletePW)

	return e
}