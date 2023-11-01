package main

import (
	"github.com/gin-gonic/gin"
	"github.com/josephe44/go-passwordless-auth/controllers"
	"github.com/josephe44/go-passwordless-auth/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()

	// Users
	r.POST("/auth", controllers.UserAuth)

	r.Run()
}
