package employee

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// ModelEmployee : Model mapping for tbl_employee
type ModelEmployee struct {
	ID          int64       `gorm:"column:id;" json:"id" form:"id" binding:"exists"`
	Userid      int64       `gorm:"column:userid;" json:"userid" form:"userid" binding:"exists"`
	Firstname   string      `gorm:"column:firstname;" json:"firstname" form:"firstname" binding:"exists"`
	Lastname    string      `gorm:"column:lastname;" json:"lastname" form:"lastname" binding:"exists"`
	Jobtitle    string      `gorm:"column:jobtitle;" json:"jobtitle" form:"jobtitle" binding:"exists"`
	Salary      float32     `gorm:"column:salary;" json:"salary" form:"salary" binding:"exists"`
	Createddate time.Time   `gorm:"column:created_date;" json:"created_date" form:"created_date" binding:"exists"`
	Createdby   string      `gorm:"column:created_by;" json:"created_by" form:"created_by" binding:"exists"`
	Updateddate pq.NullTime `gorm:"column:updated_date;" json:"updated_date" form:"updated_date"`
	Updatedby   *string     `gorm:"column:updated_by;" json:"updated_by" form:"updated_by"`
}

// TableName : Mapping table tbl_user in postgreSQL to struct
func (ModelEmployee) TableName() string {
	return "tbl_employee"
}

// SerializerEmployee : model for serializer
type SerializerEmployee struct {
	c *gin.Context
}

// ResponseEmployee : model for response
type ResponseEmployee struct {
	ID          int64       `json:"id"`
	Userid      int64       `json:"userid"`
	Firstname   string      `json:"firstname"`
	Lastname    string      `json:"lastname"`
	Jobtitle    string      `json:"jobtitle"`
	Salary      float32     `json:"salary"`
	Createddate time.Time   `json:"created_date"`
	Createdby   string      `json:"created_by"`
	Updateddate pq.NullTime `json:"updated_date"`
	Updatedby   *string     `json:"updated_by"`
}

// ModelValidatorEmployee : *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type ModelValidatorEmployee struct {
	Employee struct {
		ID          int64     `json:"id" form:"id"`
		Firstname   string    `json:"firstname" form:"firstname" binding:"exists,min=4,max=255"`
		Lastname    string    `json:"lastname" form:"lastname" binding:"exists,min=4,max=255"`
		Jobtitle    string    `json:"jobtitle" form:"jobtitle" binding:"exists,min=4,max=255"`
		Salary      float32   `json:"salary" form:"salary" binding:"exists"`
		Updateddate time.Time `json:"updated_date" form:"updated_date" binding:"exists"`
		Updatedby   string    `json:"updated_by" form:"updated_by" binding:"exists"`
	} `json:"employee"`
	EmployeeModel ModelEmployee `json:"employeemodel"`
}
