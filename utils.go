package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	dateLayout = "2006-01-02"
	timeLayout = "15:04"
)

func pointsFromRetailer(points *int, retailer string) {
	// One point for every alphanumeric character in the retailer name.
	alnum := regexp.MustCompile(`[[:alnum:]]`)
	*points += len(alnum.FindAllString(retailer, -1))
}

func pointsFromTotal(points *int, total string) {
	// 50 points if the total is a round dollar amount with no cents.
	// 25 points if the total is a multiple of 0.25.
	convTotal, _ := strconv.ParseFloat(total, 64)

	if math.Mod(convTotal, 1) == 0 {
		*points += 50
	}

	if math.Mod(convTotal, 0.25) == 0 {
		*points += 25
	}
}

func pointsFromItems(points *int, items []ReceiptItem) {
	// 5 points for every two items on the receipt.
	pairs := len(items) / 2
	*points += pairs * 5

	/**
	   If the trimmed length of the item description is a multiple of 3,
	   multiply the price by 0.2 and round up to the nearest integer.
	   The result is the number of points earned.
	**/
	for _, item := range items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			*points += int(math.Ceil(price * 0.2))
		}
	}
}

func pointsFromPurchaseDate(points *int, purchaseDate string) {
	// 6 points if the day in the purchase date is odd.
	parsedDate, _ := time.Parse(dateLayout, purchaseDate)
	if parsedDate.Day()%2 == 1 {
		*points += 6
	}
}

func pointsFromPurchaseTime(points *int, purchaseTime string) {
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	parsedTime, _ := time.Parse(timeLayout, purchaseTime)
	hour := parsedTime.Hour()
	minute := parsedTime.Minute()
	timeInMinutes := hour*60 + minute

	if timeInMinutes > 14*60 && timeInMinutes < 16*60 {
		*points += 10
	}
}

func CalculatePoints(receipt Receipt) int {
	points := 0

	pointsFromRetailer(&points, receipt.Retailer)
	pointsFromTotal(&points, receipt.Total)
	pointsFromItems(&points, receipt.Items)
	pointsFromPurchaseDate(&points, receipt.PurchaseDate)
	pointsFromPurchaseTime(&points, receipt.PurchaseTime)

	return points
}
