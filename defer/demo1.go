package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Taking care of cleaning process!")
	}()

	defer func() {
		fmt.Println("Closing opened file handlers")
	}()

	defer func() {
		fmt.Println("Closing opened database connection")
	}()

	defer func() {
		fmt.Println("Recover error if occurred.")
	}()

	fmt.Println("Doing some complex calculation, remote http call and database call!")
}
