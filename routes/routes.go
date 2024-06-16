package routes

import (
	"go_gin_boilerplate/controllers"
	"go_gin_boilerplate/utils"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
	// Public routes
	router.POST("/signup", userController.Signup)
	router.POST("/login", userController.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(utils.AuthMiddleware())
	{
		auth.GET("/users", userController.GetAllUsers)
		auth.GET("/users/:id", userController.GetUserByID)
		auth.PUT("/users/:id", userController.UpdateUser)
		auth.DELETE("/users/:id", userController.DeleteUser)
	}
}
