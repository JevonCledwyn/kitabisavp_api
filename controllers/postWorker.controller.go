package controllers

import (
	"net/http"
	"kitabisavp/helpers"
	"kitabisavp/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPW(c echo.Context) error {

	result, err := models.FetchAllPW()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StorePW(c echo.Context) error {
	post_worker_title := c.FormValue("post_worker_title")
	post_worker_description := c.FormValue("post_worker_description")
	post_worker_type := c.FormValue("post_worker_type")
	worker_id := helpers.ConvertStringToInt(c.FormValue("worker_id"))

	result, err := models.StorePW(post_worker_title, post_worker_description, post_worker_type, worker_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}

func UpdatePW(c echo.Context) error {

	post_worker_id := helpers.ConvertStringToInt(c.FormValue("post_worker_id"))
	post_worker_title := c.FormValue("post_worker_title")
	post_worker_description := c.FormValue("post_worker_description")
	post_worker_type := c.FormValue("post_worker_type")

	result, err := models.UpdatePW(post_worker_id, post_worker_title, post_worker_description, post_worker_type)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeletePW(c echo.Context) error {

	post_worker_id := helpers.ConvertStringToInt(c.FormValue("post_worker_id"))

	result, err := models.DeletePW(post_worker_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// fetch plan by id
func FetchPWById(c echo.Context) error {

	post_worker_id := c.Param("post_worker_id")

	result, err := models.FetchPWById(post_worker_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
