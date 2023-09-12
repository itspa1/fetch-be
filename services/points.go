package services

import (
	"fetch-be/models"
	"math"
	"strings"
)

// CalculatePoints calculates the points for a receipt based on the rules
func CalculatePoints(receipt *models.Receipt) int {
	points := 0

	// points based on retailer name (alphanumeric characters only)
	for _, c := range receipt.Retailer {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' || c >= '0' && c <= '9' {
			points++
		}

	}

	// points based on total
	if receipt.TotalParsed == float64(int(receipt.TotalParsed)) {
		points += 50 // 50 points if total is a whole number, i.e. no cents
	}

	if receipt.TotalParsed*4 == float64(int(receipt.TotalParsed*4)) {
		points += 25 // 25 points if total is a multiple of 0.25
	}

	// points based on items

	// 5 points for every 2 items
	points += len(receipt.Items) / 2 * 5

	// based on description of item, if trimmed len is multiple of 3, multiply price by 0.2
	// and add rounded up to nearest whole number to points
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.Description))%3 == 0 {
			points += int(math.Ceil(item.PriceInFloat * 0.2))
		}
	}

	// Time based points

	// 6 points if date is odd
	if receipt.DateParsed.Day()%2 == 1 {
		points += 6
	}

	// 10 points if time is between 2pm and 4pm
	if receipt.TimeParsed.Hour() >= 14 && receipt.TimeParsed.Hour() <= 16 {
		points += 10
	}

	return points
}
