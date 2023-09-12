package routes

import (
	"fetch-be-assignment/controllers" // Controllers

	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes all routes for the application
func InitializeRoutes(router *gin.Engine) {
	// Initialize receipts routes
	receipts := router.Group("/receipts")
	{
		// Process a receipt
		receipts.POST("/process", controllers.ProcessReceipt)

		// Get a receipt points by id
		receipts.GET("/:id/points", controllers.GetReceiptPoints)
	}
}
