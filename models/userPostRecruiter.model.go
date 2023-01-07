package models

import (
	// "database/sql"
	// "fmt"
	"kitabisavp/db"
	"net/http"

	// "self_money_management_api_golang/helpers"
	"github.com/go-playground/validator/v10"
)

type PR struct {
	PostRecruiterId          int    `json:"post_recruiter_id"`
	PostRecruiterTitle       string `json:"post_recruiter_title"`
	PostRecruiterDescription string `json:"post_recruiter_description"` 
	PostRecruiterType        string `json:"post_recruiter_type"`
	RecruiterId              int    `json:"recruiter_id "`
}

//! CRUD START

func FetchAllPR() (Response, error) {
	var obj PR
	// digunakan untuk menampung data Post Recruiter (PR)
	var arrObj []PR
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM post_recruiter"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.PostRecruiterId, &obj.PostRecruiterTitle, &obj.PostRecruiterDescription, &obj.PostRecruiterType, &obj.RecruiterId)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}

// insert data plan
func StorePR(post_recruiter_title string, post_recruiter_description string, post_recruiter_type string, recruiter_id int) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	pr := PR{
		PostRecruiterTitle:   post_recruiter_title,
		PostRecruiterDescription:  post_recruiter_description,
		PostRecruiterType:   post_recruiter_type,
		RecruiterId: recruiter_id,
	}

	err := v.Struct(pr)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `post_recruiter`(`post_recruiter_title`, `post_recruiter_description`, `post_recruiter_type`, `recruiter_id`) VALUES (?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(post_recruiter_title, post_recruiter_description, post_recruiter_type, recruiter_id)

	if err != nil {
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	lastInsertedID, err := result.LastInsertId()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"last_inserted_id": lastInsertedID,
	}

	return res, nil

}

// func update user

func UpdatePR(post_recruiter_id int, post_recruiter_title string, post_recruiter_description string, post_recruiter_type string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE post_recruiter SET post_recruiter_title=?,post_recruiter_description=?,post_recruiter_type=? WHERE post_recruiter_id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(post_recruiter_title, post_recruiter_description, post_recruiter_type, post_recruiter_id)

	if err != nil {
		return res, err
	}

	rowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": rowAffectedID,
	}

	return res, nil

}

// func delete user
func DeletePR(post_recruiter_id int) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM post_recruiter WHERE post_recruiter_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(post_recruiter_id)

	if err != nil {
		return res, err
	}

	rowAffectedID, err := result.RowsAffected()

	if err != nil {
		return res, err
	}
	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int64{
		"row_affected_id": rowAffectedID,
	}

	return res, nil

}

func FetchPRById(post_recruiter_id string) (Response, error) {
	var obj PR
	var arrObj []PR
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM plans WHERE id_user = ?"

	rows, err := con.Query(sqlStatement, post_recruiter_id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.PostRecruiterId, &obj.PostRecruiterTitle, &obj.PostRecruiterDescription, &obj.PostRecruiterType, &obj.RecruiterId)

		if err != nil {
			return res, err
		}

		arrObj = append(arrObj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrObj

	return res, nil
}
