package main

import "unicode"

var points int

func handlerPoints(receipt Receipt) int {

	points = 0

	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			points++
		}
	}

	return points
}
