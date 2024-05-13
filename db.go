package main

import (
	"fmt"

	log "github.com/zenlit/go-shared-module/logger"
	"github.com/zenlit/go-shared-module/validator"
)

// User represents a sample struct with validation rules.
type User struct {
	Username string `validate:"required,min=4"`
	Email    string `validate:"required,email"`
	Age      int    `validate:"required,numeric,min=18"`
}

func _() {
	log.Logger.Debug("DB test", "test", true)

	// Create a User instance with validation
	user := User{
		Username: "jon",
		Email:    "john@example.org",
		Age:      16,
	}

	// Validate the User struct
	if err := validator.Validate(user); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Validation passed!")
	}

	// object_store.ListBuckets()

	// http.Boot(routes, 8000)
}
