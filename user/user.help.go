package user

import (
	"errors"
	common "kreditplus/kreditplus-api/commons"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// What's bcrypt? https://en.wikipedia.org/wiki/Bcrypt
// Golang bcrypt doc: https://godoc.org/golang.org/x/crypto/bcrypt
// You can change the value in bcrypt.DefaultCost to adjust the security index.
// setPassword : err := userModel.setPassword("password0")
func (u *ModelUser) setPassword(password string) error {
	if len(password) == 0 {
		return errors.New("Password should not be empty")
	}
	bytePassword := []byte(password)
	// Make sure the second param `bcrypt generator cost` between [4, 32)
	passwordHash, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	u.PasswordHash = string(passwordHash)
	return nil
}

// checkPassword : Database will only save the hashed string, you should check it by util function.
// 	if err := userModel.checkPassword("password0"); err != nil { password error }
func (u *ModelUser) checkPassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.PasswordHash)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// UpdateContextUserModel : A helper to write user_id and user_model to the context
func UpdateContextUserModel(c *gin.Context, myUserID int64) {
	var myUserModel ModelUser
	if myUserID != 0 {
		db := common.GetPostgreSQLDB()
		db.First(&myUserModel, myUserID)
	}
	c.Set("my_user_id", myUserID)
	c.Set("my_user_model", myUserModel)
}
