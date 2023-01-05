package controllers

import (
	"net/http"
	"kitabisavp/models"
	"kitabisavp/helpers"

	"github.com/labstack/echo/v4"
)

func FetchAllRecruiter(c echo.Context) error {

	result, err := models.FetchAllRecruiter()

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// Change Struct

func StoreRecruiter(c echo.Context) error {
	recruiter_id := helpers.ConvertStringToInt(c.FormValue("recruiter_id"))
	recruiter_name := c.FormValue("recruiter_name")
	recruiter_title := c.FormValue("recruiter_title")
	recruiter_description := c.FormValue("recruiter_description")
	recruiter_contact := c.FormValue("recruiter_contact")
	recruiter_password, _ := helpers.HashPassword(c.FormValue("recruiter_password"))

	result, err := models.StoreRecruiter(recruiter_id, recruiter_name, recruiter_title, recruiter_description, recruiter_contact, recruiter_password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, result)
	}

	return c.JSON(http.StatusOK, result)

}
//	RecruiterId          int    `json:"recruiter_id"`
// 	RecruiterName        string `json:"recruiter_name"`
// 	RecruiterPassword    string `json:"recruiter_password"`
// 	RecruiterTitle       string `json:"recruiter_title"`
// 	RecruiterDescription string `json:"recruiter_description"`
// 	RecruiterContact     string `json:"recruiter_contact"`

func UpdateRecruiter(c echo.Context) error {

	recruiter_id := helpers.ConvertStringToInt(c.FormValue("recruiter_id"))
	recruiter_name := c.FormValue("recruiter_name")
	recruiter_title := c.FormValue("recruiter_title")
	recruiter_description := c.FormValue("recruiter_description")
	recruiter_contact := c.FormValue("recruiter_contact")
	recruiter_password, _ := helpers.HashPassword(c.FormValue("recruiter_password"))

	result, err := models.UpdateRecruiter(recruiter_id, recruiter_name, recruiter_title, recruiter_description, recruiter_contact, recruiter_password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

func DeleteRecruiter(c echo.Context) error {

	recruiter_id := c.FormValue("recruiter_id")

	result, err := models.DeleteRecruiter(recruiter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

// function checklogin and get user id from model checklogin without token
func CheckLogin(c echo.Context) error {
	user_recruiter := c.FormValue("user_recruiter")
	recruiter_password := c.FormValue("recruiter_password")

	result, err := models.CheckLogin(user_recruiter, recruiter_password)

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
func FetchRecruiterById(c echo.Context) error {

	recruiter_id := c.Param("recruiter_id")

	result, err := models.FetchRecruiterById(recruiter_id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError,
			map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}