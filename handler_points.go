package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

var points int

func handlerPoints(receipt Receipt) int {

	points = 0

	//One point for every alphanumeric character in the retailer name.
	count := 0
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
			count++
		}
	}
	fmt.Printf("+%v points for retailer name\n", count)

	//50 points if the total is a round dollar amount with no cents.
	//25 points if the total is a multiple of '0.25'
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		println("Error parsing receipt total")
	} else {
		cents := int(total * 100)
		if cents%100 == 0 {
			fmt.Println("+50 points for round dollar")
			points += 50
		}
		if cents%25 == 0 {
			fmt.Println("+25 points for multiple of 0.25")
			points += 25
		}
	}

	//5 points for every two items on the receipt
	itemsCount := len(receipt.Items)
	pointsMult := itemsCount / 2
	fmt.Printf("+%v points for every 2 items\n", pointsMult*5)
	points += pointsMult * 5

	//If the trimmed length of the item description is a multiple of 3,
	//multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		trimmedLen := strings.TrimSpace(item.ShortDescription)
		descriptionLen := len(trimmedLen)
		itemPrice, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			println("Error parsing item price")
		} else {
			if descriptionLen%3 == 0 {
				multiplier := math.Ceil(itemPrice * 0.2)
				fmt.Printf("+%v for item description\n", multiplier)
				points += int(multiplier)
			}
		}
	}

	//6 points if the day in the purchase date is odd
	purchaseDay := receipt.PurchaseDate.Day()
	if purchaseDay%2 != 0 {
		fmt.Println("+6 for day")
		points += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm
	purchaseTime := receipt.PurchaseTime
	hour := purchaseTime.Hour()
	if hour >= 14 && hour <= 16 {
		fmt.Println("+10 for time")
		points += 10
	}

	fmt.Println(points)
	return points
}
