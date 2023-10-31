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

	// result, error := util.GenerateOTPCode()

	// if error != nil {
	// 	fmt.Println("There is an error generating OTP")
	// 	return
	// }

	// otp := util.SendSimpleMailWithHTML(result, "./util/test.html", []string{"josephe442@gmail.com"})

	// fmt.Println(otp)

	// Users
	r.POST("/auth", controllers.UserAuth)

	r.Run()
}
