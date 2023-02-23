package api

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/rateApp/internal/helper"
	"github.com/michaelgbenle/rateApp/internal/models"
)

func (u *HTTPHandler) SignUpHandler(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		helper.Response(c, "error", 400, nil, []string{"invalid request"})
		return
	}
	//check if user already exists
	_, err = u.Repository.FindUserByEmail(user.Email)
	if err == nil {
		helper.Response(c, "error", 400, nil, []string{"user already exists"})
		return
	}

	//check if password is valid
	if !helper.IsValidPassword(user.Password) {
		helper.Response(c, "error", 400, nil, []string{"minimum of 8 characters containing upper case, lower case, number and special character required"})
	}

	//hash password
	err = user.HashPassword()
	if err != nil {
		helper.Response(c, "error", 500, nil, []string{"internal server error"})
		return
	}

	//save user to database
	err = u.Repository.CreateUser(user)
	if err != nil {
		helper.Response(c, "unable to sign up user", 500, nil, []string{"internal server error"})
		return
	}
	// successful sign up
	helper.Response(c, "sign up successful", 201, nil, nil)
}
