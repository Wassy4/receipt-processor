package utils_test

import (
	"encoding/json"
	"testing"

	"github.com/wassy4/receipt-processor/models"
	"github.com/wassy4/receipt-processor/utils"
)

type TestFixture struct {
	Receipt  Receipt
	Expected PointsResponse
}

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		fixture  Receipt
		expected int
	}{
		{
			name:     "No points",
			fixture:  "testdata/no_points.json",
			expected: 0,
		},
		{
			name:     "All of the points",
			fixture:  "testdata/points.json",
			expected: 103, // 6 (retailer) + 75 (total) + 6 (items) + 6 (purchase date) + 10 (purchase time)
		},
		{
			name:     "Provided example case #1",
			fixture:  "testdata/provided_example1.json",
			expected: 28, // 6 (retailer) + 10 (items) + 6 (multiples of 3) + 6 (purchase date)
		},
		{
			name:     "Provided example case #2",
			fixture:  "testdata/provided_example2.json",
			expected: 109, // 14 (retailer) + 75 (total) + 10 (items) + 10 (purchase time)
		},
	}

	for _, test := range tets {
		t.Run(test.name, func(t *testing.T) {
			var receipt models.Receipt
			json.Unmarshal(test.fixture, &receipt)
			points := utils.CalculatePoints(receipt)
			if points != test.expected {
				t.Errorf("Expected %d points, got %d", tt.expected, points)
			}
		})
	}
}
