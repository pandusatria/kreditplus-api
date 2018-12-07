package user

import (
	"fmt"
	common "kreditplus/kreditplus-api/commons"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Bind : There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (usermodel *ModelValidatorUser) Bind(c *gin.Context) error {
	err := common.Bind(c, usermodel)
	if err != nil {
		return err
	}

	usermodel.UserModel.Username = usermodel.User.Username
	usermodel.UserModel.Password = usermodel.User.Password
	usermodel.UserModel.Role = usermodel.User.Role

	if usermodel.User.Password != common.NBRandomPassword {
		usermodel.UserModel.setPassword(usermodel.User.Password)
	}

	usermodel.UserModel.Createddate = time.Now().Local()
	usermodel.UserModel.Createdby = "system"

	usermodel.EmployeeModel.Firstname = usermodel.User.Firstname
	usermodel.EmployeeModel.Lastname = usermodel.User.Lastname
	usermodel.EmployeeModel.Jobtitle = usermodel.User.Jobtitle
	usermodel.EmployeeModel.Salary = usermodel.User.Salary
	usermodel.EmployeeModel.Createddate = time.Now().Local()
	usermodel.EmployeeModel.Createdby = "system"

	return nil
}

// NewUserModelValidator : You can put the default value of a Validator here
func NewUserModelValidator() ModelValidatorUser {
	userModelValidator := ModelValidatorUser{}
	return userModelValidator
}

// BindUpdate : There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (usermodel *ModelValidatorUpdateUser) BindUpdate(c *gin.Context) error {
	fmt.Println("Test - BindUpdate")
	err := common.Bind(c, usermodel)
	if err != nil {
		return err
	}

	usermodel.UserModel.Username = usermodel.User.Username
	usermodel.UserModel.Password = usermodel.User.Password
	usermodel.UserModel.Role = usermodel.User.Role

	if usermodel.User.Password != common.NBRandomPassword {
		usermodel.UserModel.setPassword(usermodel.User.Password)
	}

	usermodel.UserModel.Updateddate = pq.NullTime{Time: time.Now().Local(), Valid: true}

	var updateby string
	updateby = "system"

	usermodel.UserModel.Updatedby = &updateby

	return nil
}

// UpdateUserModelValidator : You can put the default value of a Validator here
func UpdateUserModelValidator() ModelValidatorUpdateUser {
	userModelValidator := ModelValidatorUpdateUser{}
	fmt.Println("Test - UpdateUserModelValidator")
	return userModelValidator
}

// Bind : if error happen
func (loginmodel *LoginValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, loginmodel)
	if err != nil {
		return err
	}

	loginmodel.UserModel.Username = loginmodel.User.Username
	return nil
}

// NewLoginValidator : You can put the default value of a Validator here
func NewLoginValidator() LoginValidator {
	loginValidator := LoginValidator{}
	return loginValidator
}
