package tests

import (
	"fetch-be/models"
	"fetch-be/services"
	"testing"
	"time"
)

func TestCalculatePoints(t *testing.T) {
	TimeParsed, _ := time.Parse("15:04", "14:33")
	DateParsed, _ := time.Parse("2006-01-02", "2022-03-20")

	// Create a valid receipt
	receipt := models.Receipt{
		Retailer: "M&M Corner Market",
		Date:     "2022-03-20",
		Time:     "14:33",
		Total:    "9.00",
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
		TimeParsed: TimeParsed,
		DateParsed: DateParsed,
	}

	// Calculate the points
	points := services.CalculatePoints(&receipt)

	// Check if the points match the expected value
	expectedPoints := 109
	if points != expectedPoints {
		t.Errorf("Expected points %d, but got: %d", expectedPoints, points)
	}
}

func TestCalculatePointsWithOddDate(t *testing.T) {
	TimeParsed, _ := time.Parse("15:04", "14:33")
	DateParsed, _ := time.Parse("2006-01-02", "2022-03-21")

	// Create a valid receipt
	receipt := models.Receipt{
		Retailer: "M&M Corner Market",
		Date:     "2022-03-21",
		Time:     "14:33",
		Total:    "9.00",
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
		TimeParsed: TimeParsed,
		DateParsed: DateParsed,
	}

	// Calculate the points
	points := services.CalculatePoints(&receipt)

	// Check if the points match the expected value
	expectedPoints := 115
	if points != expectedPoints {
		t.Errorf("Expected points %d, but got: %d", expectedPoints, points)
	}
}

func TestCalculatePointsWithOddItems(t *testing.T) {
	TimeParsed, _ := time.Parse("15:04", "14:33")
	DateParsed, _ := time.Parse("2006-01-02", "2022-03-20")

	// Create a valid receipt
	receipt := models.Receipt{
		Retailer: "M&M Corner Market",
		Date:     "2022-03-20",
		Time:     "14:33",
		Total:    "9.00",
		Items: []models.Item{
			{
				Description: "Gatorade",
				Price:       "2.25",
			},
		},
		TimeParsed: TimeParsed,
		DateParsed: DateParsed,
	}

	// Calculate the points
	points := services.CalculatePoints(&receipt)

	// Check if the points match the expected value
	expectedPoints := 99
	if points != expectedPoints {
		t.Errorf("Expected points %d, but got: %d", expectedPoints, points)
	}
}

func TestCalculatePointsWithDescriptionTrimCheck(t *testing.T) {
	TimeParsed, _ := time.Parse("15:04", "14:33")
	DateParsed, _ := time.Parse("2006-01-02", "2022-03-20")

	// Create a valid receipt
	receipt := models.Receipt{
		Retailer: "M&M Corner Market",
		Date:     "2022-03-20",
		Time:     "14:33",
		Total:    "9.00",
		Items: []models.Item{
			{
				Description: "     Gatorade    ", // Description with leading and trailing spaces
				Price:       "2.25",
			},
			{
				Description: "     Gatorade", // Description with leading spaces
				Price:       "2.25",
			},
			{
				Description: "Gatorade", // Description without leading and trailing spaces
				Price:       "2.25",
			},
			{
				Description: "Gatorade      ", // Description with trailing spaces
				Price:       "2.25",
			},
		},
		TimeParsed: TimeParsed,
		DateParsed: DateParsed,
	}

	// Calculate the points
	points := services.CalculatePoints(&receipt)

	// Check if the points match the expected value
	expectedPoints := 109
	if points != expectedPoints {
		t.Errorf("Expected points %d, but got: %d", expectedPoints, points)
	}
}
