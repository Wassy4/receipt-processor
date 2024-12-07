package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestProcessReceipt(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := setupMux(db)

	tests := []struct {
		name             string
		method           string
		payload          string
		expectedStatus   int
		validateResponse bool
	}{
		{
			name:             "Valid request payload",
			method:           "POST",
			payload:          "testdata/provided_example1.json",
			expectedStatus:   http.StatusOK,
			validateResponse: true,
		},
		{
			name:             "Invalid request payload",
			method:           "POST",
			payload:          "testdata/invalid.json",
			expectedStatus:   http.StatusBadRequest,
			validateResponse: false,
		},
		{
			name:             "Method not allowed",
			method:           "GET",
			payload:          "testdata/provided_example1.json",
			expectedStatus:   http.StatusMethodNotAllowed,
			validateResponse: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			file, err := os.ReadFile(test.payload)
			if err != nil {
				t.Fatal(err)
			}
			req := httptest.NewRequest(test.method, "/receipts/process", bytes.NewBuffer(file))

			handler.ServeHTTP(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %v, got %v", test.expectedStatus, w.Code)
			}

			if test.validateResponse {
				var response ProcessReceiptResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				if response.Id == "" {
					t.Error("Expected ID in response")
				}
			}
		})
	}
}

func TestGetPoints(t *testing.T) {
	db, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handler := setupMux(db)

	receipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []ReceiptItem{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
	}

	receiptJSON, _ := json.Marshal(receipt)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/receipts/process",
		bytes.NewBuffer(receiptJSON))
	handler.ServeHTTP(w, req)

	var response ProcessReceiptResponse
	json.NewDecoder(w.Body).Decode(&response)

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectPoints   bool
		expectedPoints int
	}{
		{
			name:           "Valid request",
			method:         "GET",
			path:           "/receipts/" + response.Id + "/points",
			expectedStatus: http.StatusOK,
			expectPoints:   true,
			expectedPoints: 28, // 6 (retailer) + 10 (items) + 6 (multiples of 3) + 6 (purchase date)
		},
		{
			name:           "Invalid request",
			method:         "GET",
			path:           "/receipts/123/points",
			expectedStatus: http.StatusNotFound,
			expectPoints:   false,
		},
		{
			name:           "Wrong HTTP method",
			method:         "POST",
			path:           "/receipts/" + response.Id + "/points",
			expectedStatus: http.StatusMethodNotAllowed,
			expectPoints:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(test.method, test.path, nil)

			handler.ServeHTTP(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status %v, got %v", test.expectedStatus, w.Code)
			}

			if test.expectPoints {
				var response GetPointsResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				if err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				if response.Points != test.expectedPoints {
					t.Errorf("Expected %d points, got %d",
						test.expectedPoints, response.Points)
				}
			}
		})
	}
}
