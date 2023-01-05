package models

import (
	"database/sql"
	"fmt"
	"kitabisavp/db"
	"net/http"

	"github.com/go-playground/validator/v10"
	"kitabisavp/helpers"
)

type Recruiter struct {
	RecruiterId          int    `json:"recruiter_id"`
	RecruiterName        string `json:"recruiter_name"`
	RecruiterPassword    string `json:"recruiter_password"`
	RecruiterTitle       string `json:"recruiter_title"`
	RecruiterDescription string `json:"recruiter_description"`
	RecruiterContact     string `json:"recruiter_contact"`
}

//! CRUD START

func FetchAllRecruiter() (Response, error) {
	var obj Recruiter
	// digunakan untuk menampung data user
	var arrObj []Recruiter
	var res Response

	con := db.Createcon()

	sqlStatement := "SELECT * FROM user_recruiter"

	rows, err := con.Query(sqlStatement)

	// defer digunakan untuk menutup koneksi database
	defer rows.Close()

	// kalau ada error di return
	if err != nil {
		return res, err
	}

	// looping untuk menampung data user, lalu di cek apakah ada error
	for rows.Next() {
		err = rows.Scan(&obj.RecruiterId, &obj.RecruiterName, &obj.RecruiterPassword, &obj.RecruiterTitle, &obj.RecruiterDescription, &obj.RecruiterContact)

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
func StoreRecruiter(recruiter_id int, recruiter_name string, recruiter_password string, recruiter_title string, recruiter_description string, recruiter_contact string) (Response, error) {
	var res Response

	// !validasi

	v := validator.New()

	rct := Recruiter{
	RecruiterId:			recruiter_id,
	RecruiterName:			recruiter_name,	
	RecruiterPassword:		recruiter_password,
	RecruiterTitle:			recruiter_title,
	RecruiterDescription:	recruiter_description,
	RecruiterContact:		recruiter_contact,
	}

	err := v.Struct(rct)
	if err != nil {
		res.Status = http.StatusBadRequest
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	con := db.Createcon()

	sqlStatement := "INSERT INTO `user_recruiter`(`recruiter_id`, `recruiter_name`, `recruiter_password`,`recruiter_title`, `recruiter_description`, `recruiter_contact`) VALUES (?,?,?,?,?,?)"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		res.Status = http.StatusInternalServerError
		res.Message = "Error"
		res.Data = map[string]string{
			"errors": err.Error(),
		}
		return res, err
	}

	result, err := stmt.Exec(recruiter_id, recruiter_name, recruiter_password, recruiter_title, recruiter_description, recruiter_contact)

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

func UpdateRecruiter(recruiter_id int, recruiter_name string, recruiter_password string, recruiter_title string, recruiter_description string, recruiter_contact string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "UPDATE users SET recruiter_id=?,recruiter_password=?,recruiter_title=?,recruiter_description=?,recruiter_contact=? WHERE recruiter_id=?"

	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(recruiter_id, recruiter_name, recruiter_password, recruiter_title, recruiter_description, recruiter_contact)

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
func DeleteRecruiter(id string) (Response, error) {
	var res Response

	con := db.Createcon()

	sqlStatement := "DELETE FROM users WHERE id=?"
	stmt, err := con.Prepare(sqlStatement)

	if err != nil {
		return res, err
	}

	result, err := stmt.Exec(id)

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
func CheckLogin(recruiter_name, recruiter_password string) (int, error) {
	var obj Recruiter
	var pwd string
	var id int
	con := db.Createcon()

	sqlStatement := "SELECT * FROM users WHERE email = ?"
	err := con.QueryRow(sqlStatement, recruiter_name).Scan(
		&id, &obj.RecruiterName, &obj.RecruiterTitle, &obj.RecruiterDescription, &obj.RecruiterContact, &pwd,
	)

	if err == sql.ErrNoRows {
		fmt.Print("Email not found!")
		return 0, err
	}

	if err != nil {
		fmt.Print("Query error!")
		return 0, err
	}

	match, err := helpers.CheckPasswordHash(recruiter_password, pwd)
	if !match {
		fmt.Print("Hash and password doesn't match!")
		return 0, err
	}

	return id, nil
}

	// RecruiterId          int    `json:"recruiter_id"`
	// RecruiterName        string `json:"recruiter_name"`
	// RecruiterPassword    string `json:"recruiter_password"`
	// RecruiterTitle       string `json:"recruiter_title"`
	// RecruiterDescription string `json:"recruiter_description"`
	// RecruiterContact     string `json:"recruiter_contact"`

func FetchRecruiterById(recruiter_id string) (Recruiter, error) {
    var obj Recruiter

    con := db.Createcon()

    sqlStatement := "SELECT * FROM users WHERE id = ?"

    rows := con.QueryRow(sqlStatement, recruiter_id)

    err := rows.Scan(&obj.RecruiterId, &obj.RecruiterName, &obj.RecruiterPassword, &obj.RecruiterTitle, &obj.RecruiterDescription, &obj.RecruiterContact)

    if err != nil {
        return obj, err
    }

    return obj, nil
}

