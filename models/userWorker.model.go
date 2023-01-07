package models

import (
	"database/sql"
	"fmt"
	"kitabisavp/db"
	"net/http"

	"kitabisavp/helpers"

	"github.com/go-playground/validator/v10"
)

type Worker struct {
	WorkerId          int    `json:"worker_id"`
	WorkerName        string `json:"worker_name"`
	WorkerPassword    string `json:"worker_password"`
	WorkerTitle       string `json:"worker_title"`
	WorkerDescription string `json:"worker_description"`
	WorkerContact     string `json:"worker_contact"`
}

//! CRUD START

func FetchAllWorker() (Response, error) {
	var obj Worker
	// digunakan untuk menampung data user
	var arrObj []Worker
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM user_worker"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.WorkerId, &obj.WorkerName, &obj.WorkerPassword, &obj.WorkerTitle, &obj.WorkerDescription, &obj.WorkerContact)

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

// insert data user
func StoreWorker(worker_id int, worker_name string, worker_password string, worker_title string, worker_description string, worker_contact string) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	wrk := Worker{
		WorkerId:          worker_id,
		WorkerName:        worker_name,
		WorkerPassword:    worker_password,
		WorkerTitle:       worker_title,
		WorkerDescription: worker_description,
		WorkerContact:     worker_contact,
	}

	err := v.Struct(wrk)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `user_worker`(`worker_id`, `worker_name`, `worker_password`,`worker_title`, `worker_description`, `worker_contact`) VALUES (?,?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(worker_id, worker_name, worker_password, worker_title, worker_description, worker_contact)

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

func UpdateWorker(worker_id int, worker_name string, worker_password string, worker_title string, worker_description string, worker_contact string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE user_worker SET worker_name=?,worker_password=?,worker_title=?,worker_description=?,worker_contact=? WHERE worker_id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(worker_id, worker_name, worker_password, worker_title, worker_description, worker_contact)

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
func DeleteWorker(worker_id string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM user_worker WHERE worker_id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(worker_id)

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

// check login and return user id
func CheckLoginWorker(worker_name, worker_password string) (int, error) {
	var obj Worker
	var pwd string
	var id int
	con := db.Createcon()

	sqlStatement := "SELECT * FROM user_worker WHERE worker_name = ?"
	err := con.QueryRow(sqlStatement, worker_name).Scan(
		&id, &obj.WorkerName, &pwd, &obj.WorkerTitle, &obj.WorkerDescription, &obj.WorkerContact,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Email not found!")
		return 0, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return 0, err
	}

	match, err := helpers.CheckPasswordHash(worker_password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return 0, err
	}

	return id, nil
}

// 	WorkerId          int    `json:"worker_id"`
// 	WorkerName        string `json:"worker_name"`
// 	WorkerPassword    string `json:"worker_password"`
// 	WorkerTitle       string `json:"worker_title"`
// 	WorkerDescription string `json:"worker_description"`
// 	WorkerContact     string `json:"worker_contact"`

func FetchWorkerById(worker_id string) (Worker, error) {
	var obj Worker

	con := db.Createcon()

	sqlStatement := "SELECT * FROM user_worker WHERE worker_id = ?"

	rows := con.QueryRow(sqlStatement, worker_id)

	err := rows.Scan(&obj.WorkerId, &obj.WorkerName, &obj.WorkerPassword, &obj.WorkerTitle, &obj.WorkerDescription, &obj.WorkerContact)

	if err != nil {
		return obj, err
	}

	return obj, nil
}
