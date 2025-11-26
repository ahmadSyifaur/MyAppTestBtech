package routes

import (
	"myapp-test-btech/cmd/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterController)
	r.POST("/login", controllers.LoginController2)
}
