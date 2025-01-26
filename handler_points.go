package main

var points int

func handlerPoints(receipt Receipt) int {

	points = int(receipt.Points)

	return points
}
