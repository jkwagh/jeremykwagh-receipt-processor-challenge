package main

import (
	"strconv"
	"unicode"
)

var points int

func handlerPoints(receipt Receipt) int {

	points = 0

	//One point for every alphanumeric character in the retailer name.
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	//50 points if the total is a round dollar amount with no cents.
	//25 points if the total is a multiple of '0.25'
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		println("Error parsing receipt total")
	} else {
		cents := int(total * 100)
		if cents%100 == 0 {
			points += 50
		}
		if cents%25 == 0 {
			points += 25
		}
	}

	return points
}
