package main

import (
	common "kreditplus/kreditplus-api/commons"
	middleware "kreditplus/kreditplus-api/middleware"
	router "kreditplus/kreditplus-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	PostgreSQLConn := common.InitPostgreSQL()
	SQLServerConn := common.InitSQLServer()

	defer PostgreSQLConn.Close()
	defer SQLServerConn.Close()

	// Gin Web Server
	r := gin.Default()

	// For Test Purpose
	testAuth := r.Group("/api/ping")

	testAuth.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Gin Group Router
	v1 := r.Group("/api")

	v1.Use(middleware.AuthMiddleware(false))
	router.BeforeLogin(v1.Group("/beforelogin"))

	v1.Use(middleware.AuthMiddleware(true))
	router.AfterLogin(v1.Group("/afterlogin"))

	// listen and serve on 0.0.0.0:8080
	r.Run(":9090")
}
