package main

import (
	"github.com/NaheedRayan/mini_rest_api_shikho/config"
	"github.com/NaheedRayan/mini_rest_api_shikho/infrastructure"
	"github.com/NaheedRayan/mini_rest_api_shikho/infrastructure/database"
	"github.com/NaheedRayan/mini_rest_api_shikho/interfaces"
	"github.com/NaheedRayan/mini_rest_api_shikho/usecases"
)

func main() {
	// Initialize database
	config.InitDB()

	// Initialize repository
	courseRepo := database.NewCourseRepository(config.DB)

	// Initialize use case
	courseUseCase := usecases.NewCourseUseCase(courseRepo)

	// Initialize handler
	courseHandler := interfaces.NewCourseHandler(courseUseCase)

	// Setup router
	router := infrastructure.SetupRouter(courseHandler)

	// Start server
	router.Run(":8080")
}
