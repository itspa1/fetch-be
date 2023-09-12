package main

import (
	"fetch-be-assignment/routes" // Routes

	"github.com/gin-gonic/gin" // Gin web framework
)

func main() {
	r := gin.Default() // Initialize Gin

	routes.InitializeRoutes(r) // Initialize Routes

	r.Run(":8080") // Run on port 8080
}
