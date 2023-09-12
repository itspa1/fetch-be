package models

import (
	"fetch-be-assignment/common"
	"strconv"
	"time"
)

// Receipt is the model for a receipt (For now it is in memory, but in reality it would be in a database)
type Receipt struct {
	ID          string    `json:"id,omitempty"` // id user won't send this initially
	Retailer    string    `json:"retailer"`
	Date        string    `json:"purchaseDate"`
	Time        string    `json:"purchaseTime"`
	Items       []Item    `json:"items"`
	Total       string    `json:"total"`
	TotalPoints int       `json:"totalPoints"` // total points user won't send this initially
	DateParsed  time.Time // date in time.Time format (internal use only)
	TimeParsed  time.Time // time in time.Time format (internal use only)
	TotalParsed float64   // total in float format (internal use only)
}

// Validate validates the receipt, returning an error if it is invalid, or nil if it is valid
func (r *Receipt) Validate() error {

	if r.Retailer == "" {
		return &common.ValidationError{Message: "Retailer is required"}
	}

	if r.Date == "" {
		return &common.ValidationError{Message: "Date is required"}
	}

	if r.Time == "" {
		return &common.ValidationError{Message: "Time is required"}
	}

	if r.Total == "" {
		return &common.ValidationError{Message: "Total is required"}
	}

	if len(r.Items) == 0 {
		return &common.ValidationError{Message: "Items is required"}
	}

	// parse total
	var err error
	r.TotalParsed, err = strconv.ParseFloat(r.Total, 64)
	if err != nil {
		return &common.ValidationError{Message: "Total is invalid"}
	}

	// parse date and time
	r.DateParsed, err = time.Parse("2006-01-02", r.Date)
	if err != nil {
		return &common.ValidationError{Message: "Date is invalid"}
	}

	r.TimeParsed, err = time.Parse("15:04", r.Time)
	if err != nil {
		return &common.ValidationError{Message: "Time is invalid"}
	}

	// parse items and parse price to float
	for i, item := range r.Items {
		itemErr := item.Validate() // validate item from model
		if itemErr != nil {
			return &common.ValidationError{Message: "Item " + strconv.Itoa(i) + " is invalid: " + itemErr.Error()}
		}
		r.Items[i] = item
	}

	return nil
}
