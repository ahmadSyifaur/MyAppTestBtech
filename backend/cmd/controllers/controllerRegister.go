package controllers

import (
	"myapp-test-btech/cmd/models"
	"myapp-test-btech/cmd/models/register"
	"myapp-test-btech/cmd/repos"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterController(c *gin.Context) {
	var request register.RegisterRequest
	var response register.RegisterResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		response.ResponseCode = "99"
		response.ResponseMessage = "Invalid Request -" + err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	user := models.User{
		Email:        request.Email,
		PasswordHash: string(hashedPassword),
	}

	InsertUser := repos.InsertUser(user)
	if InsertUser != nil {
		response.ResponseCode = "99"
		response.ResponseMessage = "General Error -" + InsertUser.Error()
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.ResponseCode = "00"
	response.ResponseMessage = "User registered successfully"
	c.JSON(http.StatusCreated, response)
}
