package tests

import (
	"fetch-be/models"
	"testing"
)

func TestValidateValidReceipt(t *testing.T) {
	// Create a valid receipt
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "2022-03-20",
		Time:        "14:33",
		Total:       "9.00",
		TotalPoints: 50, // TotalPoints should not be checked in this test
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if there are no errors
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestValidateInvalidRetailer(t *testing.T) {
	// Create a receipt with an empty retailer
	receipt := models.Receipt{
		Retailer:    "",
		Date:        "2022-03-20",
		Time:        "14:33",
		Total:       "9.00",
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Retailer is required"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}

func TestValidateInvalidTotal(t *testing.T) {
	// Create a receipt with an invalid total
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "2022-03-20",
		Time:        "14:33",
		Total:       "invalid", // Invalid total value
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Total is invalid"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}

func TestValidateInvalidDate(t *testing.T) {
	// Create a receipt with an invalid date
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "invalid", // Invalid date value
		Time:        "14:33",
		Total:       "9.00",
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Date is invalid"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}

func TestValidateInvalidTime(t *testing.T) {
	// Create a receipt with an invalid time
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "2022-03-20",
		Time:        "invalid", // Invalid time value
		Total:       "9.00",
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Time is invalid"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}

func TestValidateInvalidItems(t *testing.T) {
	// Create a receipt with an invalid item
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "2022-03-20",
		Time:        "14:33",
		Total:       "9.00",
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "", // Invalid item description
				Price:       "2.25",
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Item 0 is invalid: Description is required"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}

func TestValidateInvalidItemPrice(t *testing.T) {
	// Create a receipt with an invalid item price
	receipt := models.Receipt{
		Retailer:    "M&M Corner Market",
		Date:        "2022-03-20",
		Time:        "14:33",
		Total:       "9.00",
		TotalPoints: 50,
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "invalid", // Invalid item price
			},
		},
	}

	// Validate the receipt
	err := receipt.Validate()

	// Check if the error message matches the expected value
	expectedErrorMessage := "Item 0 is invalid: Price is invalid"
	if err == nil || err.Error() != expectedErrorMessage {
		t.Errorf("Expected error message '%s', but got: %v", expectedErrorMessage, err)
	}
}
