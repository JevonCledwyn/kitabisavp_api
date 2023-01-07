package controllers

import (
	"kitabisavp/helpers"
	"kitabisavp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func FetchAllWorker(c echo.Context) error {

	result, err := models.FetchAllWorker()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// Change Struct

func StoreWorker(c echo.Context) error {
	worker_id := helpers.ConvertStringToInt(c.FormValue("worker_id"))
	worker_name := c.FormValue("worker_name")
	worker_password, _ := helpers.HashPassword(c.FormValue("worker_password"))
	worker_title := c.FormValue("worker_title")
	worker_description := c.FormValue("worker_description")
	worker_contact := c.FormValue("worker_contact")

	result, err := models.StoreWorker(worker_id, worker_name, worker_password, worker_title, worker_description, worker_contact)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}

func UpdateWorker(c echo.Context) error {

	worker_id := helpers.ConvertStringToInt(c.FormValue("worker_id"))
	worker_name := c.FormValue("worker_name")
	worker_password, _ := helpers.HashPassword(c.FormValue("worker_password"))
	worker_title := c.FormValue("worker_title")
	worker_description := c.FormValue("worker_description")
	worker_contact := c.FormValue("worker_contact")

	result, err := models.UpdateWorker(worker_id, worker_name, worker_password, worker_title, worker_description, worker_contact)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteWorker(c echo.Context) error {

	worker_id := c.FormValue("worker_id")
	worker_name := c.FormValue("worker_name")
	worker_password, _ := helpers.HashPassword(c.FormValue("worker_password"))
	worker_title := c.FormValue("worker_title")
	worker_description := c.FormValue("worker_description")
	worker_contact := c.FormValue("worker_contact")

	result, err := models.DeleteWorker(worker_id, worker_name, worker_password, worker_title, worker_description, worker_contact)
	

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// function checklogin and get user id from model checklogin without token
func CheckLoginWorker(c echo.Context) error {
	user_worker := c.FormValue("user_worker")
	worker_password := c.FormValue("worker_password")

	result, err := models.CheckLoginWorker(user_worker, worker_password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	//make return data to json
	return c.JSON(http.StatusOK,
		map[string]interface{}{
			"user_id": result,
			"message": "login success",
		})

}

// // function get user data from id
func FetchWorkerById(c echo.Context) error {

	worker_id := c.Param("worker_id")

	result, err := models.FetchWorkerById(worker_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
