package user

import (
	"errors"
	"fmt"
	common "kreditplus/kreditplus-api/commons"
	employee "kreditplus/kreditplus-api/employee"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Registration : func for user reqistration
func Registration(c *gin.Context) {
	userModelValidator := NewUserModelValidator()
	if err := userModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	if err := SaveUser(&userModelValidator.UserModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	usrModel, err := GetLatestUser()

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("register", errors.New("Not user exist")))
		return
	}

	userModelValidator.EmployeeModel.Userid = usrModel.ID
	if err := employee.SaveEmployee(&userModelValidator.EmployeeModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	c.Set("my_user_model", userModelValidator.UserModel)
	serializer := SerializerUser{c}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

// Login : func for user login
func Login(c *gin.Context) {
	loginValidator := NewLoginValidator()
	if err := loginValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}
	userModel, err := FindOneUser(&ModelUser{Username: loginValidator.UserModel.Username})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered username or invalid password")))
		return
	}

	if userModel.checkPassword(loginValidator.User.Password) != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Not Registered username or invalid password")))
		return
	}
	UpdateContextUserModel(c, userModel.ID)
	serializer := SerializerUser{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

// GetAll : func for get user all
func GetAll(c *gin.Context) {
	listUserModel := ListAll()
	c.JSON(http.StatusOK, gin.H{"user": listUserModel})
}

// ModifyUser : func Update User
func ModifyUser(c *gin.Context) {
	fmt.Println("Test - ModifyUser")
	userModelValidator := UpdateUserModelValidator()
	if err := userModelValidator.BindUpdate(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	userModel, err := FindOneUser(&ModelUser{ID: userModelValidator.User.ID})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("User not found")))
		return
	}

	if err := userModel.UpdateUser(userModelValidator.UserModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_user_model", userModel)
	serializer := SerializerUser{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

// GetDetail : func for user login
func GetDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userModel, err := FindOneUser(&ModelUser{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("detail", errors.New("Invalid ID or User not found")))
		return
	}
	c.Set("my_user_model", userModel)
	serializer := SerializerUser{c}
	c.JSON(http.StatusOK, gin.H{"user": serializer.Response()})
}

// RemoveUser : func for user login
func RemoveUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	userModel, err := FindOneUser(&ModelUser{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("detail", errors.New("Invalid ID or User not found")))
		return
	}

	myUserModel := c.MustGet("my_user_model").(ModelUser)

	err = myUserModel.DeleteUser(userModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	employeeModel, err := employee.FindOneEmployee(&employee.ModelEmployee{Userid: myUserModel.ID})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("detail", errors.New("Invalid Employee ID or Employee not found")))
		return
	}

	err = employee.DeleteEmployee(employeeModel)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := SerializerUser{c}
	c.JSON(http.StatusOK, gin.H{"profile": serializer.Response()})
}
