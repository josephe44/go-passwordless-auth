package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/josephe44/go-passwordless-auth/initializers"
	"github.com/josephe44/go-passwordless-auth/models"
	"github.com/josephe44/go-passwordless-auth/util"
)

type CleanUser struct {
	Email string
	OTP   string
}

func UserAuth(c *gin.Context) {
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

	// Check if the user already exists in the database
	user := models.User{}
	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			// Email doesn't exist, so create a new user
			// Generate Token
			otpCode, err := util.GenerateOTPCode()

			if err != nil {
				fmt.Println("There is an error generating OTP")
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to generate OTP",
				})
				return
			}

			user := models.User{Email: body.Email}
			result := initializers.DB.Create(&user)

			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to create user",
				})
				return
			}

			sendOTPAndRespond(c, user, otpCode)
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
		}
	} else {
		// User exists, proceed with login
		sendOTPAndRespond(c, user, "")
	}
}

func sendOTPAndRespond(c *gin.Context, user models.User, otpCode string) {
	// Generate Token
	if otpCode == "" {
		var err error
		otpCode, err = util.GenerateOTPCode()
		if err != nil {
			fmt.Println("There is an error generating OTP")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate OTP",
			})
			return
		}

		returnData := CleanUser{
			Email: user.Email,
			OTP:   otpCode,
		}

		util.SendSimpleMailHTML("OTP Sent", []string{user.Email}, otpCode)

		if otpCode != "" {
			c.Set("otp", otpCode)

			// Generate a jwt token
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"sub": user.ID,
				"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
			})

			tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Failed to create token",
				})

				return
			}
			// Set the JWT token as a cookie
			c.SetSameSite(http.SameSiteLaxMode)
			c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

			// Return it
			c.JSON(http.StatusCreated, gin.H{
				"data": returnData,
			})
		}
	}
}
