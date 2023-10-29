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
	OTP   string
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

	// Check if the email already exists in the database
	existingUser := models.User{}
	result := initializers.DB.Where("email = ?", body.Email).First(&existingUser)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			// Email doesn't exist, so create a new user
			// Generate Token
			otpCode, err := util.GenerateOTPCode()

			if err != nil {
				fmt.Println("There is an error generating OTP")
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

			returnData := CleanUser{
				Email: user.Email,
				OTP:   otpCode,
			}

			if otpCode != "" {
				c.Set("otp", otpCode)
				// Return it
				c.JSON(http.StatusCreated, gin.H{
					"data": returnData,
				})
			}
		} else {
			// Some other error occurred while trying to fetch the record
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
		}
	} else {
		// Email already exists, so return an error
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email already exists in the database",
		})
	}
}
