package user

import (
	employee "kreditplus/kreditplus-api/employee"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// ModelUser : Model mapping for tbl_user
type ModelUser struct {
	ID           int64       `gorm:"column:id;" json:"id" form:"id" binding:"exists"`
	Username     string      `gorm:"column:username;" json:"username" form:"username" binding:"exists"`
	Password     string      `gorm:"column:password;" json:"password" form:"password" binding:"exists"`
	PasswordHash string      `gorm:"column:password_hash;" json:"password_hash" form:"password_hash" binding:"exists"`
	Role         string      `gorm:"column:role;" json:"role" form:"role" binding:"exists"`
	Createddate  time.Time   `gorm:"column:created_date;" json:"created_date" form:"created_date" binding:"exists"`
	Createdby    string      `gorm:"column:created_by;" json:"created_by" form:"created_by" binding:"exists"`
	Updateddate  pq.NullTime `gorm:"column:updated_date;" json:"updated_date" form:"updated_date"`
	Updatedby    *string     `gorm:"column:updated_by;" json:"updated_by" form:"updated_by"`
}

// TableName : Mapping table tbl_user in postgreSQL to struct
func (ModelUser) TableName() string {
	return "tbl_user"
}

// ModelValidatorUser : *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type ModelValidatorUser struct {
	User struct {
		ID           int64     `json:"id" form:"id"`
		Username     string    `form:"username" json:"username" binding:"exists,min=4,max=255"`
		Password     string    `form:"password" json:"password" binding:"exists,min=5,max=255"`
		Role         string    `form:"role" json:"role" binding:"max=1024"`
		PasswordHash string    `form:"password_hash" json:"password_hash" binding:"exists"`
		Firstname    string    `json:"firstname" form:"firstname" binding:"exists,min=4,max=255"`
		Lastname     string    `json:"lastname" form:"lastname" binding:"exists,min=4,max=255"`
		Jobtitle     string    `json:"jobtitle" form:"jobtitle" binding:"exists,min=4,max=255"`
		Salary       float32   `json:"salary" form:"salary" binding:"exists"`
		Createddate  time.Time `form:"created_date" json:"created_date" binding:"exists"`
		Createdby    string    `form:"created_by" json:"created_by" binding:"exists"`
	} `json:"user"`
	Employee struct {
		Createddate time.Time `json:"created_date" form:"created_date" binding:"exists"`
		Createdby   string    `json:"created_by" form:"created_by" binding:"exists"`
	} `json:"employee"`
	UserModel     ModelUser              `json:"usermodel"`
	EmployeeModel employee.ModelEmployee `json:"employeemodel"`
}

// ModelValidatorUpdateUser : *ModelValidator containing two parts:
// - Validator: write the form/json checking rule according to the doc https://github.com/go-playground/validator
// - DataModel: fill with data from Validator after invoking common.Bind(c, self)
// Then, you can just call model.save() after the data is ready in DataModel.
type ModelValidatorUpdateUser struct {
	User struct {
		ID           int64     `json:"id" form:"id"`
		Username     string    `form:"username" json:"username" binding:"exists,min=4,max=255"`
		Password     string    `form:"password" json:"password" binding:"exists,min=5,max=255"`
		Role         string    `form:"role" json:"role" binding:"max=1024"`
		PasswordHash string    `form:"password_hash" json:"password_hash" binding:"exists"`
		Updateddate  time.Time `json:"updated_date" form:"updated_date" binding:"exists"`
		Updatedby    string    `json:"updated_by" form:"updated_by" binding:"exists"`
	} `json:"user"`
	UserModel ModelUser `json:"usermodel"`
}

// LoginValidator : model for login validator
type LoginValidator struct {
	User struct {
		Username string `form:"username" json:"username" binding:"exists,min=4"`
		Password string `form:"password" json:"password" binding:"exists,min=5,max=255"`
	} `json:"user"`
	UserModel ModelUser `json:"usermodel"`
}

// SerializerUser : model for serializer
type SerializerUser struct {
	c *gin.Context
}

// ResponseUser : model for response
type ResponseUser struct {
	ID           int64       `json:"id"`
	Username     string      `json:"username"`
	Password     string      `json:"password"`
	PasswordHash string      `json:"password_hash"`
	Role         string      `json:"role"`
	Createddate  time.Time   `json:"created_date"`
	Createdby    string      `json:"created_by"`
	Updateddate  pq.NullTime `json:"updated_date"`
	Updatedby    *string     `json:"updated_by"`
	Token        string      `json:"token"`
}
