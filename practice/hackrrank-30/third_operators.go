package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "12.00, 20, 8"
	scaner := bufio.NewScanner(strings.NewReader(input))
	var lines []float64
	for scaner.Scan() {
		readInt, _ := strconv.ParseFloat(scaner.Text(), 64)
		lines = append(lines, readInt)
	}
	mealCost := lines[0]
	tipPercent := lines[1]
	taxPercent := lines[2]

	tip := mealCost * (tipPercent / 100)
	tax := mealCost * (taxPercent / 100)
	totalCost := tip + tax + mealCost

	fmt.Printf("The total meal cost is %.0f dollars.", totalCost)

}
