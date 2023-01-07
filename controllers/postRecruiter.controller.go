package controllers

import (
	"net/http"
	"kitabisavp/helpers"
	"kitabisavp/models"

	"github.com/labstack/echo/v4"
)

func FetchAllPR(c echo.Context) error {

	result, err := models.FetchAllPR()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func StorePR(c echo.Context) error {
	post_recruiter_title := c.FormValue("post_recruiter_title")
	post_recruiter_description := c.FormValue("post_recruiter_description")
	post_recruiter_type := c.FormValue("post_recruiter_type")
	recruiter_id := helpers.ConvertStringToInt(c.FormValue("recruiter_id"))

	result, err := models.StorePR(post_recruiter_title, post_recruiter_description, post_recruiter_type, recruiter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}

func UpdatePR(c echo.Context) error {

	post_recruiter_id := helpers.ConvertStringToInt(c.FormValue("post_recruiter_id"))
	post_recruiter_title := c.FormValue("post_recruiter_title")
	post_recruiter_description := c.FormValue("post_recruiter_description")
	post_recruiter_type := c.FormValue("post_recruiter_type")

	result, err := models.UpdatePR(post_recruiter_id, post_recruiter_title, post_recruiter_description, post_recruiter_type)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeletePR(c echo.Context) error {

	post_recruiter_id := helpers.ConvertStringToInt(c.FormValue("post_recruiter_id"))

	result, err := models.DeletePR(post_recruiter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// fetch plan by id
func FetchPRById(c echo.Context) error {

	post_recruiter_id := c.Param("post_recruiter_id")

	result, err := models.FetchPRById(post_recruiter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}
