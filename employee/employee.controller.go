package employee

import (
	"errors"
	common "kreditplus/kreditplus-api/commons"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAll : func for get user all
func GetAll(c *gin.Context) {
	listEmployeeModel := ListAll()
	c.JSON(http.StatusOK, gin.H{"employee": listEmployeeModel})
}

// ModifyEmployee : func Update Employee
func ModifyEmployee(c *gin.Context) {
	employeeModelValidator := NewEmployeeModelValidator()
	if err := employeeModelValidator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	employeeModel, err := FindOneEmployee(&ModelEmployee{ID: employeeModelValidator.Employee.ID})

	if err != nil {
		c.JSON(http.StatusForbidden, common.NewError("login", errors.New("Employee not found")))
		return
	}

	if err := employeeModel.UpdateEmployee(employeeModelValidator.EmployeeModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}
	c.Set("my_employee_model", employeeModel)
	serializer := SerializerEmployee{c}
	c.JSON(http.StatusOK, gin.H{"employee": serializer.Response()})
}

// GetDetail : func for user login
func GetDetail(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	employeeModel, err := FindOneEmployee(&ModelEmployee{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, common.NewError("detail", errors.New("Invalid ID or User not found")))
		return
	}
	c.Set("my_employee_model", employeeModel)
	serializer := SerializerEmployee{c}
	c.JSON(http.StatusOK, gin.H{"employee": serializer.Response()})
}
