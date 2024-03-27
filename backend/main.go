package main

import (
	"embed"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	//go:embed configs
	embedFS embed.FS
)

// controller
// service
// repository
// model

func main() {

	// Setup routes
	r := gin.Default()

	api := r.Group("/api")
	{
		// Health check
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong ",
			})
		})

		// Car routes
		carRoutes := api.Group("/cars")
		{
			carRoutes.GET("/", GetCars)
			carRoutes.GET("/:id", GetCar)
			carRoutes.POST("/", CreateCar)
			carRoutes.PUT("/:id", UpdateCar)
			carRoutes.DELETE("/:id", DeleteCar)
		}
	}

	r.Run(":3000")
}
