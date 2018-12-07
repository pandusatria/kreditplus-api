package employee

import (
	common "kreditplus/kreditplus-api/commons"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// Bind : There are some difference when you create or update a model, you need to fill the DataModel before
// update so that you can use your origin data to cheat the validator.
// BTW, you can put your general binding logic here such as setting password.
func (employeemodel *ModelValidatorEmployee) Bind(c *gin.Context) error {
	err := common.Bind(c, employeemodel)
	if err != nil {
		return err
	}

	employeemodel.EmployeeModel.Firstname = employeemodel.Employee.Firstname
	employeemodel.EmployeeModel.Lastname = employeemodel.Employee.Lastname
	employeemodel.EmployeeModel.Jobtitle = employeemodel.Employee.Jobtitle
	employeemodel.EmployeeModel.Salary = employeemodel.Employee.Salary
	employeemodel.EmployeeModel.Updateddate = pq.NullTime{Time: time.Now().Local(), Valid: true}

	var updateby string
	updateby = "system"

	employeemodel.EmployeeModel.Updatedby = &updateby

	return nil
}

// NewEmployeeModelValidator : You can put the default value of a Validator here
func NewEmployeeModelValidator() ModelValidatorEmployee {
	employeeModelValidator := ModelValidatorEmployee{}
	return employeeModelValidator
}
