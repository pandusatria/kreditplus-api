package router

import (
	employee "kreditplus/kreditplus-api/employee"
	user "kreditplus/kreditplus-api/user"

	"github.com/gin-gonic/gin"
)

// BeforeLogin : Router handler for all function before login
func BeforeLogin(router *gin.RouterGroup) {
	router.POST("/register", user.Registration)
	router.POST("/login", user.Login)
}

// AfterLogin : Router handler for all function after login
func AfterLogin(router *gin.RouterGroup) {
	router.GET("/user", user.GetAll)
	router.GET("/user/:id", user.GetDetail)
	common.Logger("info", "AfterLogin - user.ModifyUser", "Modul Employee : ListAll")
	router.PUT("/user", user.ModifyUser)
	router.DELETE("/user/:id", user.RemoveUser)

	router.GET("/employee", employee.GetAll)
	router.GET("/employee/:id", employee.GetDetail)
	router.PUT("/employee", employee.ModifyEmployee)
}
