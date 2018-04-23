package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank"

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var myInit uint64
	var myFloat float64
	var myString string

	// Read and save an integer, double, and String to your variables.
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	myInit, _ = strconv.ParseUint(lines[0], 10, 64)
	myFloat, _ = strconv.ParseFloat(lines[1], 64)
	myString = lines[2]

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + myInit)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%.1f\n", d+myFloat)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + myString)
}
