package controllers

import (
	"fetch-be-assignment/common"
	"fetch-be-assignment/models"
	"fetch-be-assignment/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]*models.Receipt) // In memory database for receipts

// ProcessReceipt processes a receipt
func ProcessReceipt(c *gin.Context) {

	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"message": "Invalid JSON",
		})
		return
	}

	// validate receipt (if invalid, return error)
	// we use the Validate function from the model, make changes there if you want to change the validation rules
	if err := receipt.Validate(); err != nil {
		switch err.(type) {
		case *common.ValidationError:
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		default:
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
	}

	// generate id to store receipt
	id := uuid.New().String()
	receipt.ID = id

	// calculate total points
	receipt.TotalPoints = services.CalculatePoints(&receipt)
	fmt.Println(receipt.TotalPoints)
	receipts[id] = &receipt

	c.JSON(200, gin.H{
		"id": id,
	})
}

// GetReceiptPoints returns a receipt points by id (if it exists)
func GetReceiptPoints(c *gin.Context) {
	id := c.Param("id") // get id from url

	receipt, ok := receipts[id]
	if !ok {
		c.JSON(404, gin.H{
			"message": "Receipt not found",
		})
		return
	}

	// return points
	c.JSON(200, gin.H{
		"points": receipt.TotalPoints,
	})
}
