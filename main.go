package main

import (
	"go_gin_boilerplate/controllers"
	"go_gin_boilerplate/migrations"
	"go_gin_boilerplate/repository"
	"go_gin_boilerplate/routes"
	"go_gin_boilerplate/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Connect to PostgreSQL database
	db, err := utils.ConnectDB()
	if err != nil {
		panic("Failed to connect to database")
	}

	// Run migrations
	if err := migrations.Migrate(db); err != nil {
		panic("Failed to migrate database")
	}

	// Create repository and controller instances
	userRepo := repository.UserRepository{DB: db}
	userController := controllers.UserController{Repository: &userRepo}

	// Set up Gin router
	router := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Authorization", "Content-Type"},
	}))

	// Initialize routes
	routes.SetupRoutes(router, &userController)

	// Start server
	router.Run(":8000")
}
