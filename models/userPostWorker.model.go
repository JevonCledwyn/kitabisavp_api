package models

import (
	// "database/sql"
	// "fmt"
	"kitabisavp/db"
	"net/http"

	// "self_money_management_api_golang/helpers"
	"github.com/go-playground/validator/v10"
)

type PW struct {
	PostWorkerId          int    `json:"post_worker_id "`
	PostWorkerTitle       string `json:"post_worker_title"`
	PostWorkerDescription string `json:"post_worker_description"` 
	PostWorkerType        string `json:"post_worker_type"`
	WorkerId              int    `json:"worker_id "`
}

//! CRUD START

func FetchAllPW() (Response, error) {
	var obj PW
	// digunakan untuk menampung data Post Worker (PW)
	var arrObj []PW
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM post_worker"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.PostWorkerId, &obj.PostWorkerTitle, &obj.PostWorkerDescription, &obj.PostWorkerType, &obj.WorkerId)

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
func StorePW(post_worker_title string, post_worker_description string, post_worker_type string, worker_id int) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	pw := PW{
		PostWorkerTitle:   post_worker_title,
		PostWorkerDescription:  post_worker_description,
		PostWorkerType:   post_worker_type,
		WorkerId: worker_id,
	}

	err := v.Struct(pw)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `post_worker`(`post_worker_title`, `post_worker_description`, `post_worker_type`, `worker_id`) VALUES (?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(post_worker_title, post_worker_description, post_worker_type, worker_id)

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

func UpdatePW(post_worker_id int, post_worker_title string, post_worker_description string, post_worker_type string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE post_worker SET post_worker_title=?,post_worker_description=?,post_worker_type=? WHERE post_worker_id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(post_worker_title, post_worker_description, post_worker_type, post_worker_id)

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
func DeletePW(post_worker_id int) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM post_worker WHERE post_worker_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(post_worker_id)

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

func FetchPWById(post_worker_id string) (Response, error) {
	var obj PW
	var arrObj []PW
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM plans WHERE id_user = ?"

	rows, err := con.Query(sqlStatement, post_worker_id)

	defer rows.Close()

	if err != nil {
		return res, err
	}

	for rows.Next() {
		err = rows.Scan(&obj.PostWorkerId, &obj.PostWorkerTitle, &obj.PostWorkerDescription, &obj.PostWorkerType, &obj.WorkerId)

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
