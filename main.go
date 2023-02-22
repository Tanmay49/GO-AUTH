package main

import (
	"main/Init"
	"main/Middleware"
	Route "main/Routes"

	"github.com/gin-gonic/gin"
)

func init() {
	Init.LoadEnv()
	Init.ConnectDB()
}

func main() {
	router := gin.Default()

	router.GET("/ping", Route.Test)
	router.POST("/signup", Route.Signup)
	router.POST("/login", Route.Login)
	router.GET("/validate", Middleware.RequireAuth, Route.Validate)
	router.Run()
}
