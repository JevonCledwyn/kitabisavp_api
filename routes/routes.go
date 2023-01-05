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
	e.POST("/login", controllers.CheckLogin)


	// RecruiterId          int    `json:"recruiter_id"`
	// RecruiterName        string `json:"recruiter_name"`
	// RecruiterPassword    string `json:"recruiter_password"`
	// RecruiterTitle       string `json:"recruiter_title"`
	// RecruiterDescription string `json:"recruiter_description"`
	// RecruiterContact     string `json:"recruiter_contact"`
	// CRUD RECRUITERS


	// e.GET("/user", controllers.FetchAllUser,middleware.IsAuthenticated)
	e.GET("/recruiter", controllers.FetchAllRecruiter)
	e.POST("/recruiter", controllers.StoreRecruiter)
	e.PATCH("/recruiter", controllers.UpdateRecruiter)
	e.DELETE("/recruiter", controllers.DeleteRecruiter)
	e.GET("/recruiter/:id", controllers.FetchRecruiterById)

	// // plan
	// e.GET("/plan", controllers.FetchAllPlan)
	// e.GET("/plan/:id", controllers.FetchPlanById)
	// e.POST("/plan", controllers.StorePlan)
	// e.PATCH("/plan", controllers.UpdatePlan)
	// e.DELETE("/plan", controllers.DeletePlan)

	// //money
	// e.GET("/money/:id", controllers.FetchMoneyById)
	// e.GET("/moneyTotalPemasukan/:id", controllers.FetchTotalPemasukanById)
	// e.GET("/moneyTotalPengeluaran/:id", controllers.FetchTotalPengeluaranById)
	// e.PATCH("/money", controllers.UpdateMoney)
	// e.DELETE("/money", controllers.DeleteMoney)
	// e.GET("/currency", controllers.FetchMoneyAPI)
	// //1pemasukan
	// e.GET("/moneyPemasukan/:id", controllers.FetchDataPemasukanByUserId)
	// e.GET("/moneyPengeluaran/:id", controllers.FetchDataPengeluaranByUserId)

	// // validation
	// e.POST("/test-validation", controllers.TestStructValidation)
	// e.POST("/test-validation-var", controllers.TestVarValidation)
	return e
}