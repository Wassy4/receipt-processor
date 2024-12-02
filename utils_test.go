package main

import (
	"encoding/json"
	"os"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected int
	}{
		{
			name:     "No points",
			filepath: "testdata/no_points.json",
			expected: 0,
		},
		{
			name:     "All of the points",
			filepath: "testdata/all_points.json",
			expected: 103, // 5 (retailer) + 75 (total) + 6 (items) + 6 (purchase date) + 10 (purchase time)
		},
		{
			name:     "Provided example case #1",
			filepath: "testdata/provided_example1.json",
			expected: 28, // 6 (retailer) + 10 (items) + 6 (multiples of 3) + 6 (purchase date)
		},
		{
			name:     "Provided example case #2",
			filepath: "testdata/provided_example2.json",
			expected: 109, // 14 (retailer) + 75 (total) + 10 (items) + 10 (purchase time)
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var receipt Receipt

			fixture, _ := os.ReadFile(test.filepath)
			json.Unmarshal(fixture, &receipt)
			points := CalculatePoints(receipt)
			if points != test.expected {
				t.Errorf("Expected %d points, got %d", test.expected, points)
			}
		})
	}
}
