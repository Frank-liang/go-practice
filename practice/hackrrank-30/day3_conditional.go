package main

import (
	"flag"
	"fmt"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	/* This section has some problems, myInt can not get the value
	scanner := bufio.NewReader(os.Stdin)
	num, _ := scanner.ReadString('\n')
	fmt.Printf("%T %v\n", num, num)
	myInt, _ := strconv.ParseInt(num, 10, 64)
	fmt.Printf("%T %v\n", myInt, myInt)
	*/
	num := flag.Int("number", 3, "Input a number")
	flag.Parse()
	myInt := *num

	//myInt, _ := strconv.Atoi(num)

	if odd(myInt) {
		fmt.Println("Weird")
	} else if even(myInt) && 2 < myInt && myInt < 5 {
		fmt.Println("Not Weird")
	} else if even(myInt) && 6 < myInt && myInt < 20 {
		fmt.Println("Weird")
	} else {
		fmt.Println("Not Weird")
	}
}
func even(num int) bool {
	if num%2 == 0 {
		return true
	}
	return false
}

func odd(num int) bool {
	if num%2 != 0 {
		return true
	}
	return false
}
