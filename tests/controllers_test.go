package tests

import (
	"encoding/json"
	"fetch-be/routes"
	"io"
	"strings"

	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	routes.InitializeRoutes(r)
	return r
}

func sendRequestAndCheckStatusCode(t *testing.T, r *gin.Engine, method, path string, body io.Reader, expectedStatusCode int) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != expectedStatusCode {
		t.Errorf("Expected status code %d, got %d", expectedStatusCode, w.Code)
	}
	return w
}

func unmarshalResponseBody(t *testing.T, w *httptest.ResponseRecorder, responseBody interface{}) {
	if err := json.Unmarshal(w.Body.Bytes(), responseBody); err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}
}

func TestProcessReceipt(t *testing.T) {
	r := setupRouter()

	// Test invalid receipt
	_ = sendRequestAndCheckStatusCode(t, r, "POST", "/receipts/process", nil, 400)

	// Test valid receipt
	obj := `
	{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
		  {
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  }
		],
		"total": "9.00"
	}
	` // valid receipt JSON
	w := sendRequestAndCheckStatusCode(t, r, "POST", "/receipts/process", strings.NewReader(obj), 200)

	// Check if id is returned
	var responseBody map[string]interface{}
	unmarshalResponseBody(t, w, &responseBody)
	if _, idExists := responseBody["id"]; !idExists {
		t.Error("Expected id to be returned")
	}

	// Validation error are tested in models/model_test.go
}

func TestGetReceiptPoints(t *testing.T) {
	r := setupRouter()

	// Create receipt
	obj := `{
		"retailer": "M&M Corner Market",
		"purchaseDate": "2022-03-20",
		"purchaseTime": "14:33",
		"items": [
		  {
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  },{
			"shortDescription": "Gatorade",
			"price": "2.25"
		  }
		],
		"total": "9.00"
	  }` // valid receipt JSON
	w := sendRequestAndCheckStatusCode(t, r, "POST", "/receipts/process", strings.NewReader(obj), 200)
	var responseBody map[string]interface{}
	unmarshalResponseBody(t, w, &responseBody)
	id, idExists := responseBody["id"]
	if !idExists {
		t.Error("Expected id to be returned")
	}

	// Get points for the receipt created
	w = sendRequestAndCheckStatusCode(t, r, "GET", "/receipts/"+id.(string)+"/points", nil, 200)
	unmarshalResponseBody(t, w, &responseBody)

	// Check if the "points" field exists in the response
	if _, pointsExists := responseBody["points"]; !pointsExists {
		t.Error("Expected points to be returned")
	}
	// Check points is 109
	if points := responseBody["points"].(float64); points != 109 {
		t.Error("Expected points to be 109")
	}
}
