package main

import "fmt"
import "os"
import "strconv"

func main() {
	Args := os.Args
	Sum := 0
	for i := 0; i < len(Args); i++ {
		num, err := strconv.Atoi(Args[i])
		if err != nil {
			fmt.Println(err)
		}
		Sum = Sum + num
	}
	fmt.Println("sum: ", Sum)
}
