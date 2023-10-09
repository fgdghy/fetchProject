package utils

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/fetchProject/app/models"
)

var (
	letterRegex = regexp.MustCompile("[a-zA-Z]")
)

func CalculatePoints(receipt *models.Receipt) (int, error) {
	points := 0

	// Rule 1: One point for every alphanumeric character in the retailer name.
	points += len(letterRegex.FindAllString(receipt.Retailer, -1))

	// Rule 2: 50 points if the total is a round dollar amount with no cents.
	totalFloat, err := strconv.ParseFloat(receipt.Total, 64)
	if err == nil && math.Mod(totalFloat, 1) == 0 {
		points += 50
	}

	// Rule 3: 25 points if the total is a multiple of 0.25.
	if err == nil && math.Mod(totalFloat, 0.25) == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items on the receipt.
	points += len(receipt.Items) / 2 * 5

	// Rule 5: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2
	// and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)

		if len(trimmedDescription)%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, err
			}
			points += int(math.Ceil(priceFloat * 0.2))
		}
	}

	// Rule 6: 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err != nil {
		return 0, err
	} else if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// Rule 7: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	startTime := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 14, 0, 0, 0, purchaseTime.Location())
	endTime := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 16, 0, 0, 0, purchaseTime.Location())

	if err != nil {
		return 0, err
	} else if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		points += 10
	}

	return points, nil
}
