package main

import (
	"github.com/josephe44/go-passwordless-auth/initializers"
	"github.com/josephe44/go-passwordless-auth/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.User{})

}
