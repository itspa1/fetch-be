package models

import (
	"fetch-be-assignment/common"
	"strconv"
)

// Item is the model for an item on a receipt (For now it is in memory, but in reality it would be in a database)
type Item struct {
	Description  string  `json:"shortDescription"`
	Price        string  `json:"price"`
	PriceInFloat float64 // price in float format (internal use only)
}

// Validate validates the item, returning an error if it is invalid, or nil if it is valid
func (i *Item) Validate() error {
	if i.Description == "" {
		return &common.ValidationError{Message: "Description is required"}
	}

	if i.Price == "" {
		return &common.ValidationError{Message: "Price is required"}
	}

	// parse price
	var err error
	i.PriceInFloat, err = strconv.ParseFloat(i.Price, 64)
	if err != nil {
		return &common.ValidationError{Message: "Price is invalid"}
	}

	return nil
}
