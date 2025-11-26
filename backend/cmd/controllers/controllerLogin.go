package controllers

import (
	"myapp-test-btech/cmd/models"
	"myapp-test-btech/cmd/models/login"
	"myapp-test-btech/cmd/models/register"
	"myapp-test-btech/cmd/repos"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(c *gin.Context) {
	var input login.LoginRequest
	var response register.RegisterResponse
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ResponseCode = "98"
		response.ResponseMessage = "Invalid Request -" + err.Error()
		c.JSON(http.StatusBadRequest, response)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	user := models.User{
		Email:        input.Email,
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

func LoginController2(c *gin.Context) {
	var request login.LoginRequest
	var response login.LoginResponse
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repos.GetUserByEmail(request.Email)
	if err != nil {
		response.ResponseCode = "99"
		response.ResponseMessage = "Invalid email or password"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		response.ResponseCode = "99"
		response.ResponseMessage = "Invalid email or password"
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	// Generate JWT
	jwtSecret := os.Getenv("JWT_SECRET")
	expireMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRE_MINUTES")) // fixed 15 minutes
	if err != nil {
		expireMinutes = 15
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Minute * time.Duration(expireMinutes)).Unix(),
	})

	tokenString, _ := token.SignedString([]byte(jwtSecret))
	response.ResponseCode = "00"
	response.ResponseMessage = "Successful login"
	response.Token = tokenString
	c.JSON(http.StatusOK, response)
}
