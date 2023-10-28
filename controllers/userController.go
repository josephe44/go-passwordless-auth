package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/josephe44/go-passwordless-auth/initializers"
	"github.com/josephe44/go-passwordless-auth/models"
	"github.com/josephe44/go-passwordless-auth/util"
)

type CleanUser struct {
	Email string
	Token string
}

func Signup(c *gin.Context) {
	// Get the email off req body
	var body struct {
		Email string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	user := models.User{Email: body.Email}
	result := initializers.DB.Create(&user)

	// Respond
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Generate Token
	otpCode, err := util.GenerateOTPCode()

	if err != nil {
		fmt.Println("There is an error generating OTP")
		return
	}

	otp := util.SendSimpleMailWithHTML(otpCode, "../util/test.html", []string{user.Email})

	// josephe442@gmail.com
	if otp != "" {
		c.Set("otp", otpCode)
		// Return it
		c.JSON(http.StatusCreated, gin.H{
			"message": "OTP is sent to your email",
		})

	}

}
