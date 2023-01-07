package helpers

import (
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(recruiter_password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(recruiter_password), 14)
	return string(bytes), err
   }
   
   func CheckPasswordHash(recruiter_password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(recruiter_password))
   
	if err != nil {
	 return false, err
	}
   
	return true, nil
   }

// convert string to int
func ConvertStringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		// handle error
		return 0
	}
	return i
}
