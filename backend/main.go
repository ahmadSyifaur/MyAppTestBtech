package main

import (
	"fmt"
	"myapp-test-btech/cmd/controllers"
	"myapp-test-btech/cmd/helpers"
	"myapp-test-btech/configs"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.ConnectDB()

	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.RegisterController)
		v1.POST("/login", controllers.LoginController2)
	}
	// Protected routes (with JWT)
	protected := r.Group("/api/v1")
	protected.Use(helpers.JWTAuthMiddleware())
	{
		protected.GET("/dashboard", func(c *gin.Context) {
			email, _ := c.Get("email")
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Hello %s, welcome back", email),
			})
		})
	}

	r.Run(":9909")
}
